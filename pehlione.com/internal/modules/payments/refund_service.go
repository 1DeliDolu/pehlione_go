package payments

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"pehlione.com/app/internal/emails"
	emailmod "pehlione.com/app/internal/modules/email"
	"pehlione.com/app/internal/modules/orders"
)

var (
	ErrNoSucceededPayment = errors.New("no succeeded payment found")
	ErrNotRefundable      = errors.New("order not refundable")
)

type RefundService struct {
	db       *gorm.DB
	provider Provider
	emailSvc *emailmod.OutboxService
	baseURL  string
}

func NewRefundService(db *gorm.DB, p Provider, emailSvc *emailmod.OutboxService, baseURL string) *RefundService {
	return &RefundService{db: db, provider: p, emailSvc: emailSvc, baseURL: baseURL}
}

type RefundOrderInput struct {
	OrderID        string
	ActorUserID    string // admin
	IdempotencyKey string
	AmountCents    int // 0 => full remaining
	Reason         string
}

type RefundOrderResult struct {
	RefundID    string
	Status      string
	AmountCents int
	Idempotent  bool
}

func (s *RefundService) RefundOrder(ctx context.Context, in RefundOrderInput) (RefundOrderResult, error) {
	if in.OrderID == "" || in.ActorUserID == "" || in.IdempotencyKey == "" {
		return RefundOrderResult{}, ErrNotRefundable
	}

	// Phase-1: lock order + find payment + idempotency + create refund(initiated)
	var ord orders.Order
	var pay Payment
	var ref Refund
	var amount int

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// order lock
		if err := tx.WithContext(ctx).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&ord, "id = ?", in.OrderID).Error; err != nil {
			return err
		}

		// refundable gate (MVP): paid/partially_refunded states only
		if ord.Status != "paid" && ord.Status != "partially_refunded" {
			return ErrNotRefundable
		}

		// find succeeded payment (latest)
		if err := tx.WithContext(ctx).
			Order("updated_at DESC").
			First(&pay, "order_id = ? AND status = ?", ord.ID, StatusSucceeded).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrNoSucceededPayment
			}
			return err
		}

		remaining := ord.TotalCents - ord.RefundedCents
		if remaining <= 0 {
			return ErrNotRefundable
		}

		amount = in.AmountCents
		if amount <= 0 {
			amount = remaining // full remaining
		}
		if amount > remaining {
			amount = remaining
		}

		// idempotency: payment_id + key
		var existing Refund
		e := tx.WithContext(ctx).First(&existing, "payment_id = ? AND idempotency_key = ?", pay.ID, in.IdempotencyKey).Error
		if e == nil {
			ref = existing
			return nil
		}
		if e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
			return e
		}

		now := time.Now()
		var reasonPtr *string
		if in.Reason != "" {
			r := in.Reason
			reasonPtr = &r
		}

		ref = Refund{
			ID:             uuid.NewString(),
			OrderID:        ord.ID,
			PaymentID:      pay.ID,
			Provider:       s.provider.Name(),
			ProviderRef:    nil,
			Status:         StatusInitiated,
			AmountCents:    amount,
			Currency:       ord.Currency,
			IdempotencyKey: in.IdempotencyKey,
			Reason:         reasonPtr,
			ErrorMessage:   nil,
			CreatedAt:      now,
			UpdatedAt:      now,
		}
		return tx.WithContext(ctx).Create(&ref).Error
	})
	if err != nil {
		return RefundOrderResult{}, err
	}

	// idempotent hit
	if ref.Status == StatusSucceeded {
		return RefundOrderResult{RefundID: ref.ID, Status: ref.Status, AmountCents: ref.AmountCents, Idempotent: true}, nil
	}

	// Phase-2: provider refund (outside tx)
	paymentRef := ""
	if pay.ProviderRef != nil {
		paymentRef = *pay.ProviderRef
	}
	resp, perr := s.provider.RefundPayment(ctx, RefundRequest{
		OrderID:        ord.ID,
		PaymentID:      pay.ID,
		PaymentRef:     paymentRef,
		AmountCents:    ref.AmountCents,
		Currency:       ref.Currency,
		IdempotencyKey: in.IdempotencyKey,
		Reason:         in.Reason,
	})

	// Phase-3: finalize (tx)
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		now := time.Now()

		// reload+lock order (consistency)
		if err := tx.WithContext(ctx).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&ord, "id = ?", ord.ID).Error; err != nil {
			return err
		}

		// update refund row
		upd := map[string]any{"updated_at": now}
		if resp.ProviderRef != "" {
			upd["provider_ref"] = resp.ProviderRef
		}

		// async: initiated (webhook will finalize)
		if resp.Status == StatusInitiated {
			upd["status"] = StatusInitiated
			if err := tx.WithContext(ctx).Model(&Refund{}).Where("id = ?", ref.ID).Updates(upd).Error; err != nil {
				return err
			}
			// Don't touch order; webhook will finalize it
			return nil
		}

		if perr != nil || resp.Status != StatusSucceeded {
			msg := "refund failed"
			if perr != nil {
				msg = perr.Error()
			}
			upd["status"] = StatusFailed
			upd["error_message"] = msg

			if err := tx.WithContext(ctx).Model(&Refund{}).Where("id = ?", ref.ID).Updates(upd).Error; err != nil {
				return err
			}

			// ledger: refund_failed (optional but good for audit)
			fe := orders.FinancialEntry{
				ID:          uuid.NewString(),
				OrderID:     ord.ID,
				Event:       "refund_failed",
				AmountCents: 0,
				Currency:    ord.Currency,
				RefType:     "refund",
				RefID:       ref.ID,
				CreatedAt:   now,
			}
			_ = tx.WithContext(ctx).Create(&fe).Error

			// order_events (audit)
			ev := orders.OrderEvent{
				ID:          uuid.NewString(),
				OrderID:     ord.ID,
				ActorUserID: in.ActorUserID,
				Action:      "refund",
				FromStatus:  ord.Status,
				ToStatus:    ord.Status,
				Note:        ptr("refund failed: " + msg),
				CreatedAt:   now,
			}
			_ = tx.WithContext(ctx).Create(&ev).Error

			return nil
		}

		// succeeded
		upd["status"] = StatusSucceeded
		upd["error_message"] = nil
		if err := tx.WithContext(ctx).Model(&Refund{}).Where("id = ?", ref.ID).Updates(upd).Error; err != nil {
			return err
		}

		// ledger: refund_succeeded (-out)
		fe := orders.FinancialEntry{
			ID:          uuid.NewString(),
			OrderID:     ord.ID,
			Event:       "refund_succeeded",
			AmountCents: -ref.AmountCents,
			Currency:    ord.Currency,
			RefType:     "refund",
			RefID:       ref.ID,
			CreatedAt:   now,
		}
		if err := tx.WithContext(ctx).Create(&fe).Error; err != nil {
			return err
		}

		// order update: refunded_cents + status
		newRefunded := ord.RefundedCents + ref.AmountCents
		newStatus := ord.Status
		var refundedAt *time.Time

		if newRefunded >= ord.TotalCents {
			newRefunded = ord.TotalCents
			newStatus = "refunded"
			t := now
			refundedAt = &t
		} else {
			newStatus = "partially_refunded"
		}

		if err := tx.WithContext(ctx).Model(&orders.Order{}).
			Where("id = ?", ord.ID).
			Updates(map[string]any{
				"refunded_cents": newRefunded,
				"status":         newStatus,
				"refunded_at":    refundedAt,
				"updated_at":     now,
			}).Error; err != nil {
			return err
		}

		// order_events (audit)
		ev := orders.OrderEvent{
			ID:          uuid.NewString(),
			OrderID:     ord.ID,
			ActorUserID: in.ActorUserID,
			Action:      "refund",
			FromStatus:  ord.Status,
			ToStatus:    newStatus,
			Note:        ptr("refund_id=" + ref.ID),
			CreatedAt:   now,
		}
		if err := tx.WithContext(ctx).Create(&ev).Error; err != nil {
			return err
		}

		var orderItems []orders.OrderItem
		orderItems, _ = s.loadOrderItems(ctx, tx, ord.ID)

		if s.emailSvc != nil {
			if emailAddr, err := s.lookupOrderEmail(ctx, tx, ord); err == nil && emailAddr != "" {
				statusLabel := "Refunded"
				if newStatus == "partially_refunded" {
					statusLabel = "Partially refunded"
				}
				reason := strings.TrimSpace(in.Reason)
				payload := emails.BuildOrderPayload(s.baseURL, ord, orderItems, statusLabel, reason)
				payload["PreviewText"] = "Refund processed - funds will post shortly."
				_ = s.emailSvc.EnqueueTx(ctx, tx, emailmod.Job{
					To:       emailAddr,
					Template: emailmod.TemplateOrderRefunded,
					Payload:  payload,
				})
			}
		}

		return nil
	})
	if err != nil {
		return RefundOrderResult{}, err
	}

	finalStatus := resp.Status
	if perr != nil {
		finalStatus = StatusFailed
	}
	return RefundOrderResult{RefundID: ref.ID, Status: finalStatus, AmountCents: ref.AmountCents, Idempotent: false}, nil
}

func ptr(s string) *string { return &s }

func (s *RefundService) lookupOrderEmail(ctx context.Context, tx *gorm.DB, ord orders.Order) (string, error) {
	if ord.GuestEmail != nil && *ord.GuestEmail != "" {
		return *ord.GuestEmail, nil
	}
	if ord.UserID != nil && *ord.UserID != "" {
		var email string
		if err := tx.WithContext(ctx).Table("users").Select("email").Where("id = ?", *ord.UserID).Take(&email).Error; err == nil {
			return email, nil
		}
	}
	return "", nil
}

func (s *RefundService) loadOrderItems(ctx context.Context, tx *gorm.DB, orderID string) ([]orders.OrderItem, error) {
	var items []orders.OrderItem
	err := tx.WithContext(ctx).Order("created_at ASC").Find(&items, "order_id = ?", orderID).Error
	return items, err
}
