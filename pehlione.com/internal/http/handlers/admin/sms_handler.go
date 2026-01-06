package admin

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"pehlione.com/app/internal/http/flash"
	"pehlione.com/app/internal/http/middleware"
	"pehlione.com/app/internal/http/render"
	"pehlione.com/app/internal/sms"
	"pehlione.com/app/templates/pages"
)

type SmsHandler struct {
	db    *gorm.DB
	flash *flash.Codec
	logger *slog.Logger
}

func NewSmsHandler(db *gorm.DB, flashCodec *flash.Codec, logger *slog.Logger) *SmsHandler {
	return &SmsHandler{
		db:    db,
		flash: flashCodec,
		logger: logger,
	}
}

func (h *SmsHandler) ListFailed(c *gin.Context) {
	var failedSMS []sms.OutboxMessage
	if err := h.db.WithContext(c.Request.Context()).
		Where("status = ?", "failed").
		Order("created_at DESC").
		Find(&failedSMS).Error; err != nil {
		h.logger.Error("failed to list failed sms jobs", "error", err)
		render.ErrorPage(c, http.StatusInternalServerError, "Failed to load failed SMS jobs.", middleware.GetRequestID(c))
		return
	}

	flash := middleware.GetFlash(c)
	pages.AdminFailedSMSList(failedSMS, flash).Render(c.Request.Context(), c.Writer)
}
