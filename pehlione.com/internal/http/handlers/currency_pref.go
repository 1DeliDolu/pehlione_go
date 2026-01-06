package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pehlione.com/app/internal/http/middleware"
	"pehlione.com/app/internal/modules/currency"
)

type CurrencyPreferenceHandler struct {
	Service      *currency.Service
	CookieName   string
	CookieSecure bool
}

func NewCurrencyPreferenceHandler(svc *currency.Service, cookieName string, secure bool) *CurrencyPreferenceHandler {
	if cookieName == "" {
		cookieName = "currency_pref"
	}
	return &CurrencyPreferenceHandler{Service: svc, CookieName: cookieName, CookieSecure: secure}
}

func (h *CurrencyPreferenceHandler) Post(c *gin.Context) {
	code := c.PostForm("currency")
	if normalized, ok := h.Service.NormalizeDisplay(code); ok {
		middleware.SetCurrencyCookie(c, h.CookieName, normalized, h.CookieSecure)
	}
	referer := c.Request.Referer()
	if referer == "" {
		referer = "/"
	}
	c.Redirect(http.StatusSeeOther, referer)
}
