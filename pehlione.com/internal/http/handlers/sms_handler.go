package handlers

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"pehlione.com/app/internal/http/flash"
	"pehlione.com/app/internal/http/middleware"
	"pehlione.com/app/internal/http/render"
	"pehlione.com/app/internal/modules/auth"
	"pehlione.com/app/internal/sms"
	"pehlione.com/app/pkg/view"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

type SmsHandler struct {
	db      *gorm.DB
	smsRepo *sms.OutboxRepository
	logger  *slog.Logger
	flash   *flash.Codec
}

func NewSmsHandler(db *gorm.DB, smsRepo *sms.OutboxRepository, flashCodec *flash.Codec, logger *slog.Logger) *SmsHandler {
	return &SmsHandler{
		db:      db,
		smsRepo: smsRepo,
		flash:   flashCodec,
		logger:  logger,
	}
}

func (h *SmsHandler) PostAccountSMS(c *gin.Context) {
	// 1. Get user from context
	ctxUser, ok := middleware.CurrentUser(c)
	if !ok {
		render.Redirect(c, "/login")
		return
	}
	var currentUser auth.User
	if err := h.db.First(&currentUser, "id = ?", ctxUser.ID).Error; err != nil {
		render.Redirect(c, "/login")
		return
	}

	// 2. Parse form
	phone := c.PostForm("phone")
	optIn := c.PostForm("sms_opt_in") == "on"

	// 3. Basic validation
	if phone == "" {
		render.RedirectWithFlash(c, h.flash, "/account", view.FlashError, "Phone number is required.")
		return
	}

	// 4. Update user
	updates := map[string]interface{}{
		"phone_e164": &phone,
		"sms_opt_in": optIn,
	}
	if !optIn {
		updates["sms_opt_out_at"] = time.Now()
	}
	err := h.db.WithContext(c.Request.Context()).Model(&currentUser).Updates(updates).Error

	if err != nil {
		h.logger.Error("failed to update user phone", "error", err, "user_id", currentUser.ID)
		render.RedirectWithFlash(c, h.flash, "/account", view.FlashError, "Failed to update your settings.")
		return
	}

	// 5. Create consent record
	action := "opt_out"
	if optIn {
		action = "opt_in"
	}
	consent := sms.Consent{
		UserID:    currentUser.ID,
		PhoneE164: phone,
		Action:    action,
		Source:    "profile",
		IP:        c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
		CreatedAt: time.Now(),
	}
	if err := h.db.WithContext(c.Request.Context()).Create(&consent).Error; err != nil {
		h.logger.Error("failed to create consent record", "error", err, "user_id", currentUser.ID)
		// Do not block the user for this, just log it.
	}

	render.RedirectWithFlash(c, h.flash, "/account", view.FlashSuccess, "Your SMS settings have been updated.")
}

func (h *SmsHandler) PostAccountSMSVerify(c *gin.Context) {
	// 1. Get user from context
	ctxUser, ok := middleware.CurrentUser(c)
	if !ok {
		render.Redirect(c, "/login")
		return
	}
	var currentUser auth.User
	if err := h.db.First(&currentUser, "id = ?", ctxUser.ID).Error; err != nil {
		render.Redirect(c, "/login")
		return
	}

	// 2. Parse form
	code := c.PostForm("code")

	// 3. Basic validation
	if code == "" {
		render.RedirectWithFlash(c, h.flash, "/account", view.FlashError, "Verification code is required.")
		return
	}

	// 4. Find the verification record
	var verification sms.PhoneVerification
	err := h.db.WithContext(c.Request.Context()).
		Where("user_id = ? AND used_at IS NULL", currentUser.ID).
		Order("created_at DESC").
		First(&verification).Error

	if err != nil {
		render.RedirectWithFlash(c, h.flash, "/account", view.FlashError, "Invalid verification code.")
		return
	}

	// 5. Check expiry
	if time.Now().After(verification.ExpiresAt) {
		render.RedirectWithFlash(c, h.flash, "/account", view.FlashError, "Verification code has expired.")
		return
	}

	// 6. Check code hash
	hasher := sha256.New()
	hasher.Write([]byte(code))
	codeHash := hex.EncodeToString(hasher.Sum(nil))

	if codeHash != verification.CodeHash {
		render.RedirectWithFlash(c, h.flash, "/account", view.FlashError, "Invalid verification code.")
		return
	}

	// 7. Mark as verified
	now := time.Now()
	verification.UsedAt = &now
	if err := h.db.WithContext(c.Request.Context()).Save(&verification).Error; err != nil {
		h.logger.Error("failed to mark verification as used", "error", err, "verification_id", verification.ID)
		render.RedirectWithFlash(c, h.flash, "/account", view.FlashError, "An error occurred during verification.")
		return
	}

	if err := h.db.WithContext(c.Request.Context()).Model(&currentUser).Update("phone_verified_at", &now).Error; err != nil {
		h.logger.Error("failed to update user phone_verified_at", "error", err, "user_id", currentUser.ID)
		render.RedirectWithFlash(c, h.flash, "/account", view.FlashError, "An error occurred during verification.")
		return
	}

	render.RedirectWithFlash(c, h.flash, "/account", view.FlashSuccess, "Your phone number has been verified.")
}

func (h *SmsHandler) PostSendCode(c *gin.Context) {
	// 1. Get user from context
	ctxUser, ok := middleware.CurrentUser(c)
	if !ok {
		render.Redirect(c, "/login")
		return
	}
	var currentUser auth.User
	if err := h.db.First(&currentUser, "id = ?", ctxUser.ID).Error; err != nil {
		render.Redirect(c, "/login")
		return
	}

	// 2. Check for phone number
	if currentUser.PhoneE164 == nil || *currentUser.PhoneE164 == "" {
		render.RedirectWithFlash(c, h.flash, "/account", view.FlashError, "Please add a phone number first.")
		return
	}

	// 3. Generate OTP
	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	hasher := sha256.New()
	hasher.Write([]byte(otp))
	otpHash := hex.EncodeToString(hasher.Sum(nil))

	// 4. Store verification record
	verification := sms.PhoneVerification{
		UserID:    currentUser.ID,
		PhoneE164: *currentUser.PhoneE164,
		CodeHash:  otpHash,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	if err := h.db.WithContext(c.Request.Context()).Create(&verification).Error; err != nil {
		h.logger.Error("failed to create phone verification record", "error", err, "user_id", currentUser.ID)
		render.RedirectWithFlash(c, h.flash, "/account", view.FlashError, "Failed to send verification code.")
		return
	}

	// 5. Enqueue SMS
	_, err := h.smsRepo.Enqueue(c.Request.Context(), sms.EnqueueOptions{
		ToPhoneE164: *currentUser.PhoneE164,
		Template:    "otp",
		Payload: map[string]interface{}{
			"code": otp,
		},
	})
	if err != nil {
		h.logger.Error("failed to enqueue OTP sms", "error", err, "user_id", currentUser.ID)
		render.RedirectWithFlash(c, h.flash, "/account", view.FlashError, "Failed to send verification code.")
		return
	}

	render.RedirectWithFlash(c, h.flash, "/account", view.FlashSuccess, "A verification code has been sent to your phone.")
}
