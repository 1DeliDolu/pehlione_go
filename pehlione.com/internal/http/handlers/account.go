package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"pehlione.com/app/internal/http/flash"
	"pehlione.com/app/internal/http/middleware"
	"pehlione.com/app/internal/http/render"
	"pehlione.com/app/internal/modules/auth"
	"pehlione.com/app/pkg/view"
	"pehlione.com/app/templates/pages"
	"pehlione.com/app/internal/shared/apperr"
	"strings"
)

type AccountHandler struct {
	authRepo *auth.Repo
	Flash    *flash.Codec
}

func NewAccountHandler(authRepo *auth.Repo, flashCodec *flash.Codec) *AccountHandler {
	return &AccountHandler{authRepo: authRepo, Flash: flashCodec}
}

func (h *AccountHandler) Get(c *gin.Context) {
	user, ok := middleware.CurrentUser(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	render.Component(c, http.StatusOK, pages.Account(
		middleware.GetFlash(c),
		middleware.GetCSRFToken(c),
		user.Email,
	))
}

func (h *AccountHandler) ChangePassword(c *gin.Context) {
	userCtx, ok := middleware.CurrentUser(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	fullUser, err := h.authRepo.GetByID(c.Request.Context(), userCtx.ID)
	if err != nil {
		c.Error(apperr.Wrap(err))
		return
	}

	errorsMap := map[string]string{}
	current := strings.TrimSpace(c.PostForm("current_password"))
	newPass := strings.TrimSpace(c.PostForm("new_password"))
	confirm := strings.TrimSpace(c.PostForm("confirm_password"))

	if current == "" {
		errorsMap["current_password"] = "Mevcut şifreyi girin."
	}
	if len(newPass) < 6 {
		errorsMap["new_password"] = "Yeni şifre en az 6 karakter olmalı."
	}
	if newPass != confirm {
		errorsMap["confirm_password"] = "Şifreler eşleşmiyor."
	}

	if len(errorsMap) == 0 {
		if err := bcrypt.CompareHashAndPassword([]byte(fullUser.PasswordHash), []byte(current)); err != nil {
			errorsMap["current_password"] = "Mevcut şifre hatalı."
		}
	}

	if len(errorsMap) > 0 {
		// How to render errors? The form is on a different page now.
		// For now, let's just redirect back with a generic error flash.
		render.RedirectWithFlash(c, h.Flash, "/account", view.FlashError, "Şifre güncellenemedi, lütfen hataları kontrol edin.")
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		c.Error(apperr.Wrap(err))
		return
	}

	if err := h.authRepo.UpdatePassword(c.Request.Context(), fullUser.ID, string(hashed)); err != nil {
		c.Error(apperr.Wrap(err))
		return
	}

	flashCodec := h.Flash
	render.RedirectWithFlash(c, flashCodec, "/account", view.FlashSuccess, "Şifreniz güncellendi.")
}