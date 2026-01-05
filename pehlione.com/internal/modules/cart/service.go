package cart

import (
	"context"
	"errors"
	"sort"

	"gorm.io/gorm"

	"pehlione.com/app/internal/http/cartcookie"
	"pehlione.com/app/pkg/view"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

type cartRow struct {
	VariantID  string `gorm:"column:variant_id"`
	Qty        int    `gorm:"column:qty"`
	PriceCents int    `gorm:"column:price_cents"`
	Currency   string `gorm:"column:currency"`

	ProductName string `gorm:"column:product_name"`
	ProductSlug string `gorm:"column:product_slug"`
	ImageURL    string `gorm:"column:image_url"`
}

var ErrMixedCurrency = errors.New("cart contains multiple currencies")

func (s *Service) BuildCartPageForUser(ctx context.Context, userID string) (view.CartPage, error) {
	if userID == "" {
		return view.CartPage{}, errors.New("missing userID")
	}

	// Not: carts'ta status alanı yoksa WHERE c.status = 'open' kısmını kaldırın
	const q = `
SELECT
  ci.variant_id AS variant_id,
  ci.quantity   AS qty,
  v.price_cents AS price_cents,
  v.currency    AS currency,
  p.name        AS product_name,
  p.slug        AS product_slug,
  '' AS image_url
FROM carts c
JOIN cart_items ci ON ci.cart_id = c.id
JOIN product_variants v ON v.id = ci.variant_id
JOIN products p ON p.id = v.product_id
WHERE c.user_id = ?
ORDER BY ci.created_at ASC;
`

	var rows []cartRow
	if err := s.db.WithContext(ctx).Raw(q, userID).Scan(&rows).Error; err != nil {
		return view.CartPage{}, err
	}

	return buildCartVMFromRows(rows)
}

func (s *Service) BuildCartPageFromCookie(ctx context.Context, c *cartcookie.Cart) (view.CartPage, error) {
	if c == nil || len(c.Items) == 0 {
		return view.CartPage{Items: []view.CartItem{}}, nil
	}

	// variantID -> qty
	qtyByID := make(map[string]int, len(c.Items))
	ids := make([]string, 0, len(c.Items))
	for _, it := range c.Items {
		if it.VariantID == "" || it.Qty <= 0 {
			continue
		}
		if _, ok := qtyByID[it.VariantID]; !ok {
			ids = append(ids, it.VariantID)
		}
		qtyByID[it.VariantID] += it.Qty
	}
	if len(ids) == 0 {
		return view.CartPage{Items: []view.CartItem{}}, nil
	}

	// IN sorgusu deterministik olsun diye sırala
	sort.Strings(ids)

	// Cookie'de qty var; DB'den sadece ürün/variant bilgisi çekeceğiz.
	const q = `
SELECT
  v.id          AS variant_id,
  0             AS qty,
  v.price_cents AS price_cents,
  v.currency    AS currency,
  p.name        AS product_name,
  p.slug        AS product_slug,
  '' AS image_url
FROM product_variants v
JOIN products p ON p.id = v.product_id
WHERE v.id IN ?;
`

	var rows []cartRow
	if err := s.db.WithContext(ctx).Raw(q, ids).Scan(&rows).Error; err != nil {
		return view.CartPage{}, err
	}

	// DB sonuçlarını map'le
	infoByID := make(map[string]cartRow, len(rows))
	for _, r := range rows {
		infoByID[r.VariantID] = r
	}

	// Cookie sırasını koruyarak final rows üret
	final := make([]cartRow, 0, len(ids))
	for _, it := range c.Items {
		if it.VariantID == "" || it.Qty <= 0 {
			continue
		}
		r, ok := infoByID[it.VariantID]
		if !ok {
			// Variant silinmiş / yok: bu item'ı atlayalım
			continue
		}
		r.Qty = it.Qty
		final = append(final, r)
	}

	return buildCartVMFromRows(final)
}

func buildCartVMFromRows(rows []cartRow) (view.CartPage, error) {
	vm := view.CartPage{Items: make([]view.CartItem, 0, len(rows))}

	currency := ""
	subtotalCents := 0
	count := 0

	for _, r := range rows {
		if r.Qty <= 0 {
			continue
		}
		if currency == "" {
			currency = r.Currency
		} else if r.Currency != "" && r.Currency != currency {
			return view.CartPage{}, ErrMixedCurrency
		}

		line := r.PriceCents * r.Qty
		subtotalCents += line
		count += r.Qty

		vm.Items = append(vm.Items, view.CartItem{
			ProductName: r.ProductName,
			ProductSlug: r.ProductSlug,
			ImageURL:    r.ImageURL,
			VariantID:   r.VariantID,
			Qty:         r.Qty,

			UnitPriceCents: r.PriceCents,
			LineTotalCents: line,

			UnitPrice: view.MoneyFromCents(r.PriceCents, currency),
			LineTotal: view.MoneyFromCents(line, currency),
		})
	}

	vm.Currency = currency
	vm.Count = count
	vm.SubtotalCents = subtotalCents
	vm.Subtotal = view.MoneyFromCents(subtotalCents, currency)

	// MVP: total = subtotal (shipping/tax/discount sonra)
	vm.TotalCents = subtotalCents
	vm.Total = view.MoneyFromCents(subtotalCents, currency)

	return vm, nil
}
