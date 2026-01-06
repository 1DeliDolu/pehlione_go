package handlers

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"pehlione.com/app/internal/http/cartcookie"
	"pehlione.com/app/internal/http/flash"
	"pehlione.com/app/internal/http/middleware"
	"pehlione.com/app/internal/http/render"
	"pehlione.com/app/internal/http/validation"
	"pehlione.com/app/internal/modules/auth"
	cartmod "pehlione.com/app/internal/modules/cart"
	"pehlione.com/app/pkg/view"
	"pehlione.com/app/templates/pages"
)

// normalizeReturnTo validates and sanitizes the return_to parameter.
// Open redirect protection: only relative paths are accepted.
func normalizeReturnTo(s string) string {
	if s == "" {
		return ""
	}
	if len(s) < 1 || s[0] != '/' {
		return ""
	}
	// "//evil.com" gibi protocol-relative engeli
	if len(s) >= 2 && s[0:2] == "//" {
		return ""
	}
	// "http://", "https://" gibi şema engeli
	if containsScheme(s) {
		return ""
	}
	return s
}

func containsScheme(s string) bool {
	for i := 0; i+2 < len(s); i++ {
		if s[i] == ':' && s[i+1] == '/' && s[i+2] == '/' {
			return true
		}
	}
	return false
}

// AuthHandlers contains handlers for authentication routes.
type AuthHandlers struct {
	db        *gorm.DB
	flash     *flash.Codec
	sessCfg   middleware.SessionCfg
	repo      *auth.Repo
	cartCK    *cartcookie.Codec
	verifySvc interface {
		StartEmailVerification(ctx context.Context, userID, userEmail string) error
	}
}

// NewAuthHandlers creates a new AuthHandlers instance.
func NewAuthHandlers(db *gorm.DB, flashCodec *flash.Codec, sessCfg middleware.SessionCfg, cartCK *cartcookie.Codec) *AuthHandlers {
	return &AuthHandlers{
		db:        db,
		flash:     flashCodec,
		sessCfg:   sessCfg,
		repo:      auth.NewRepo(db),
		cartCK:    cartCK,
		verifySvc: nil,
	}
}

// SetVerifyService sets the verification service
func (h *AuthHandlers) SetVerifyService(svc interface {
	StartEmailVerification(ctx context.Context, userID, userEmail string) error
}) {
	h.verifySvc = svc
}

// SignupGet renders the signup page.
func (h *AuthHandlers) SignupGet(c *gin.Context) {
	returnTo := normalizeReturnTo(c.Query("return_to"))
	render.Component(c, http.StatusOK, pages.Signup(
		middleware.GetFlash(c),
		middleware.GetCSRFToken(c),
		returnTo,
		view.SignupForm{},
		nil,
	))
}

type signupInput struct {
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required,min=6"`
	PasswordConfirm string `form:"password_confirm" binding:"required,eqfield=Password"`
}

// SignupPost handles user registration.
func (h *AuthHandlers) SignupPost(c *gin.Context) {
	returnTo := normalizeReturnTo(c.PostForm("return_to"))

	var in signupInput
	if err := c.ShouldBind(&in); err != nil {
		errs := validation.FromBindError(err, &in)
		render.Component(c, http.StatusBadRequest, pages.Signup(
			middleware.GetFlash(c),
			middleware.GetCSRFToken(c),
			returnTo,
			view.SignupForm{Email: in.Email},
			errs,
		))
		return
	}

	// Check if email already exists
	if _, err := h.repo.GetByEmail(in.Email); err == nil {
		render.Component(c, http.StatusConflict, pages.Signup(
			middleware.GetFlash(c),
			middleware.GetCSRFToken(c),
			returnTo,
			view.SignupForm{Email: in.Email},
			map[string]string{"email": "Bu e-posta adresi zaten kullanılıyor."},
		))
		return
	}

	// Hash password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Error(err)
		return
	}

	// Create user
	user := &auth.User{
		Email:        strings.ToLower(in.Email),
		PasswordHash: string(hashedPwd),
		Status:       "pending",
	}
	if err := h.repo.Create(user); err != nil {
		c.Error(err)
		return
	}

	// Create session immediately (user can verify email while logged in)
	sess, err := middleware.CreateSession(h.sessCfg, user.ID)
	if err != nil {
		c.Error(err)
		return
	}
	c.SetCookie(h.sessCfg.CookieName, sess.ID, int(h.sessCfg.TTL.Seconds()), "/", "", h.sessCfg.Secure, true)

	// Start email verification if service is available
	if h.verifySvc != nil {
		if err := h.verifySvc.StartEmailVerification(c.Request.Context(), user.ID, user.Email); err != nil {
			log.Printf("Failed to start email verification: %v", err)
			c.Error(err)
			return
		}
	}

	// Redirect to verify page
	dest := "/verify"
	if returnTo != "" {
		dest = "/verify?return_to=" + returnTo
	}
	render.RedirectWithFlash(c, h.flash, dest, view.FlashSuccess, "Doğrulama kodu e-postanıza gönderildi.")
}

// LoginGet renders the login page.
func (h *AuthHandlers) LoginGet(c *gin.Context) {
	returnTo := normalizeReturnTo(c.Query("return_to"))
	render.Component(c, http.StatusOK, pages.Login(
		middleware.GetFlash(c),
		middleware.GetCSRFToken(c),
		returnTo,
		view.LoginForm{},
		nil,
		"",
	))
}

// LoginPost handles user login.
func (h *AuthHandlers) LoginPost(c *gin.Context) {
	returnTo := normalizeReturnTo(c.PostForm("return_to"))

	var in loginInput
	if err := c.ShouldBind(&in); err != nil {
		errs := validation.FromBindError(err, &in)
		render.Component(c, http.StatusBadRequest, pages.Login(
			middleware.GetFlash(c),
			middleware.GetCSRFToken(c),
			returnTo,
			view.LoginForm{Email: in.Email},
			errs,
			"",
		))
		return
	}

	// Find user by email
	user, err := h.repo.GetByEmail(in.Email)
	if err != nil {
		render.Component(c, http.StatusUnauthorized, pages.Login(
			middleware.GetFlash(c),
			middleware.GetCSRFToken(c),
			returnTo,
			view.LoginForm{Email: in.Email},
			nil,
			"E-posta veya şifre hatalı.",
		))
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.Password)); err != nil {
		render.Component(c, http.StatusUnauthorized, pages.Login(
			middleware.GetFlash(c),
			middleware.GetCSRFToken(c),
			returnTo,
			view.LoginForm{Email: in.Email},
			nil,
			"E-posta veya şifre hatalı.",
		))
		return
	}

	// Create session
	sess, err := middleware.CreateSession(h.sessCfg, user.ID)
	if err != nil {
		c.Error(err)
		return
	}

	// Set session cookie
	c.SetCookie(h.sessCfg.CookieName, sess.ID, int(h.sessCfg.TTL.Seconds()), "/", "", h.sessCfg.Secure, true)

	// Merge guest cart to user cart
	if cc, _ := h.cartCK.Get(c); cc != nil && len(cc.Items) > 0 {
		h.mergeGuestCart(c, user.ID, cc)
	}

	// Redirect to return_to or home
	dest := "/"
	if returnTo != "" {
		dest = returnTo
	}
	render.RedirectWithFlash(c, h.flash, dest, view.FlashSuccess, "Giriş başarılı.")
}

// LogoutPost handles user logout.
func (h *AuthHandlers) LogoutPost(c *gin.Context) {
	sessionID, err := c.Cookie(h.sessCfg.CookieName)
	if err == nil && sessionID != "" {
		_ = middleware.DeleteSession(h.sessCfg, sessionID)
	}

	// Clear session cookie
	c.SetCookie(h.sessCfg.CookieName, "", -1, "/", "", h.sessCfg.Secure, true)

	render.RedirectWithFlash(c, h.flash, "/", view.FlashInfo, "Çıkış yapıldı.")
}

// mergeGuestCart merges guest cart items into user's DB cart
func (h *AuthHandlers) mergeGuestCart(c *gin.Context, userID string, cc *cartcookie.Cart) {
	repo := cartmod.NewRepo(h.db)

	// Get or create user cart
	userCart, err := repo.GetOrCreateUserCart(c.Request.Context(), userID)
	if err != nil {
		log.Printf("mergeGuestCart: failed to get user cart: %v", err)
		return
	}

	// Add each cookie item to DB cart
	for _, item := range cc.Items {
		if item.VariantID == "" || item.Qty <= 0 {
			continue
		}
		if err := repo.AddItem(c.Request.Context(), userCart.ID, item.VariantID, item.Qty); err != nil {
			log.Printf("mergeGuestCart: failed to add item %s: %v", item.VariantID, err)
			// Continue with other items even if one fails
		}
	}

	// Clear guest cart cookie
	h.cartCK.Clear(c)

	// Clear session cart cache
	middleware.ClearSessionCartCache(c)
}

// VerifyGet renders the email verification page
func (h *AuthHandlers) VerifyGet(c *gin.Context) {
	returnTo := normalizeReturnTo(c.Query("return_to"))
	render.Component(c, http.StatusOK, pages.VerifyEmail(
		middleware.GetFlash(c),
		middleware.GetCSRFToken(c),
		returnTo,
	))
}

// VerifyPost verifies the email code
