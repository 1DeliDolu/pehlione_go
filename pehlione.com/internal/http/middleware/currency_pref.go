package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"pehlione.com/app/internal/modules/currency"
)

const displayCurrencyKey = "display_currency"

type CurrencyPrefCfg struct {
	Service    *currency.Service
	CookieName string
	Secure     bool
}

func CurrencyPreference(cfg CurrencyPrefCfg) gin.HandlerFunc {
	cookieName := cfg.CookieName
	if cookieName == "" {
		cookieName = "currency_pref"
	}
	return func(c *gin.Context) {
		display := cfg.Service.DefaultDisplayCurrency()
		if cur := strings.TrimSpace(c.Query("currency")); cur != "" {
			if normalized, ok := cfg.Service.NormalizeDisplay(cur); ok {
				display = normalized
				SetCurrencyCookie(c, cookieName, normalized, cfg.Secure)
			}
		} else if cookie, err := c.Cookie(cookieName); err == nil {
			if normalized, ok := cfg.Service.NormalizeDisplay(cookie); ok {
				display = normalized
			}
		}

		c.Set(displayCurrencyKey, display)
		if cfg.Service != nil {
			c.Set("currency_options", cfg.Service.DisplayOptions())
		}
		c.Next()
	}
}

func SetCurrencyCookie(c *gin.Context, name, value string, secure bool) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Expires:  time.Now().Add(365 * 24 * time.Hour),
		HttpOnly: false,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})
}

func GetDisplayCurrency(c *gin.Context) string {
	if v, ok := c.Get(displayCurrencyKey); ok {
		if s, ok := v.(string); ok && s != "" {
			return s
		}
	}
	return ""
}

func GetCurrencyOptions(c *gin.Context) []string {
	if v, ok := c.Get("currency_options"); ok {
		if list, ok := v.([]string); ok {
			return list
		}
	}
	return nil
}
