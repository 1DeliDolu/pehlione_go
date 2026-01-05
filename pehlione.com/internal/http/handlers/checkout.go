package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"pehlione.com/app/internal/http/cartcookie"
	"pehlione.com/app/internal/http/flash"
	"pehlione.com/app/internal/http/middleware"
	"pehlione.com/app/internal/http/render"
	"pehlione.com/app/internal/http/validation"
	cartmod "pehlione.com/app/internal/modules/cart"
	"pehlione.com/app/internal/modules/checkout"
	"pehlione.com/app/internal/modules/orders"
	"pehlione.com/app/internal/shared/apperr"
	"pehlione.com/app/pkg/view"
	"pehlione.com/app/templates/pages"
)

type CheckoutHandler struct {
	DB      *gorm.DB
	Flash   *flash.Codec
	CartCK  *cartcookie.Codec
	OrderSv *orders.Service
}

func NewCheckoutHandler(db *gorm.DB, fl *flash.Codec, ck *cartcookie.Codec, osvc *orders.Service) *CheckoutHandler {
	return &CheckoutHandler{DB: db, Flash: fl, CartCK: ck, OrderSv: osvc}
}

type checkoutInput struct {
	Email string `form:"email" binding:"omitempty,email,max=255"`

	FirstName  string `form:"first_name" binding:"required,min=2,max=100"`
	LastName   string `form:"last_name" binding:"required,min=2,max=100"`
	Address1   string `form:"address1" binding:"required,min=5,max=255"`
	Address2   string `form:"address2" binding:"omitempty,max=255"`
	City       string `form:"city" binding:"required,min=2,max=100"`
	PostalCode string `form:"postal_code" binding:"required,min=2,max=32"`
	Country    string `form:"country" binding:"required,len=2"`
	Phone      string `form:"phone" binding:"required,min=5,max=32"`

	ShippingMethod string `form:"shipping_method" binding:"required,oneof=standard express"`
	IdemKey        string `form:"idempotency_key" binding:"omitempty,max=64"`
}

type addressJSON struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2,omitempty"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	Phone      string `json:"phone"`
}

func (h *CheckoutHandler) Get(c *gin.Context) {
	u, authed := middleware.CurrentUser(c)

	summary, itemsCount, currency, err := h.buildCartSummary(c)
	if err != nil {
		log.Printf("Checkout GET: buildCartSummary failed for user %v: %v", u.ID, err)
		middleware.Fail(c, apperr.Wrap(err))
		return
	}
	if itemsCount == 0 {
		render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashError, "Sepet boş.")
		return
	}

	idem := randHex(16)
	form := view.CheckoutForm{
		ShippingMethod: "standard",
		IdemKey:        idem,
	}
	if authed {
		form.Email = u.Email
	}

	opts := shippingOptions(currency)
	render.Component(c, http.StatusOK, pages.Checkout(
		middleware.GetFlash(c),
		middleware.GetCSRFToken(c),
		form,
		nil,
		"",
		opts,
		summary,
		authed,
	))
}

func (h *CheckoutHandler) Post(c *gin.Context) {
	u, authed := middleware.CurrentUser(c)

	summary, itemsCount, currency, err := h.buildCartSummary(c)
	if err != nil {
		middleware.Fail(c, apperr.Wrap(err))
		return
	}
	if itemsCount == 0 {
		render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashError, "Sepet boş.")
		return
	}

	var in checkoutInput
	if err := c.ShouldBind(&in); err != nil {
		errs := validation.FromBindError(err, &in)
		h.renderCheckoutWithErrors(c, authed, summary, currency, errs, "", in)
		return
	}

	if !authed && strings.TrimSpace(in.Email) == "" {
		errs := validation.FieldErrors{"email": "Email zorunludur."}
		h.renderCheckoutWithErrors(c, authed, summary, currency, errs, "", in)
		return
	}

	shipCents := shippingCents(in.ShippingMethod)

	addr := addressJSON{
		FirstName:  strings.TrimSpace(in.FirstName),
		LastName:   strings.TrimSpace(in.LastName),
		Address1:   strings.TrimSpace(in.Address1),
		Address2:   strings.TrimSpace(in.Address2),
		City:       strings.TrimSpace(in.City),
		PostalCode: strings.TrimSpace(in.PostalCode),
		Country:    strings.ToUpper(strings.TrimSpace(in.Country)),
		Phone:      strings.TrimSpace(in.Phone),
	}
	addrBytes, err := json.Marshal(addr)
	if err != nil {
		middleware.Fail(c, apperr.Wrap(err))
		return
	}

	var userID *string
	var guestEmail *string
	var cartID string

	if authed {
		userID = &u.ID
		// Get user cart ID
		crt, err := cartmod.NewRepo(h.DB).GetOrCreateUserCart(c.Request.Context(), u.ID)
		if err != nil {
			middleware.Fail(c, apperr.Wrap(err))
			return
		}
		cartID = crt.ID
	} else {
		em := strings.ToLower(strings.TrimSpace(in.Email))
		guestEmail = &em

		// Guest: create temporary cart from cookie
		cc, _ := h.CartCK.Get(c)
		if cc == nil || len(cc.Items) == 0 {
			render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashError, "Sepet boş.")
			return
		}

		// Create temp cart in DB for order creation
		tempCart, err := h.createTempCartFromCookie(c, cc)
		if err != nil {
			middleware.Fail(c, apperr.Wrap(err))
			return
		}
		cartID = tempCart.ID
	}

	idem := strings.TrimSpace(in.IdemKey)
	if idem == "" {
		idem = randHex(16)
	}
	idemKey := &idem
	if !authed {
		idemKey = nil
	}

	log.Printf("Creating order: cartID=%s, userID=%v, guestEmail=%v, shipCents=%d", cartID, userID, guestEmail, shipCents)

	res, err := h.OrderSv.CreateFromCart(c.Request.Context(), orders.CreateFromCartInput{
		CartID:              cartID,
		UserID:              userID,
		GuestEmail:          guestEmail,
		IdempotencyKey:      idemKey,
		TaxCents:            0,
		ShippingCents:       shipCents,
		DiscountCents:       0,
		ShippingAddressJSON: addrBytes,
		BillingAddressJSON:  nil,
	})
	if err != nil {
		var oos *checkout.OutOfStockError
		if errors.As(err, &oos) {
			log.Printf("Checkout failed: out of stock - %v", err)
			render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashError, "Bazı ürünler stokta yok. Lütfen sepeti güncelleyin.")
			return
		}
		if errors.Is(err, orders.ErrCartEmpty) {
			log.Printf("Checkout failed: cart empty")
			render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashError, "Sepet boş.")
			return
		}
		if errors.Is(err, orders.ErrProductUnavailable) {
			log.Printf("Checkout failed: product unavailable")
			render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashError, "Bazı ürünler mevcut değil.")
			return
		}
		if errors.Is(err, orders.ErrCurrencyMismatch) {
			log.Printf("Checkout failed: currency mismatch")
			render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashError, "Para birimi uyuşmazlığı.")
			return
		}
		log.Printf("Checkout error (unhandled): %T - %v", err, err)
		h.renderCheckoutWithErrors(c, authed, summary, currency, nil, "Checkout başarısız. Lütfen tekrar deneyin.", in)
		return
	}

	if !authed {
		h.CartCK.Clear(c)
	}

	// Clear session cart cache (forces refresh on next request)
	middleware.ClearSessionCartCache(c)

	render.RedirectWithFlash(c, h.Flash, "/orders/"+res.OrderID, view.FlashSuccess, "Sipariş oluşturuldu.")
}

// --- helpers ---

func (h *CheckoutHandler) buildCartSummary(c *gin.Context) (view.CheckoutSummary, int, string, error) {
	svc := cartmod.NewService(h.DB)

	var cartPage view.CartPage
	var err error

	if u, ok := middleware.CurrentUser(c); ok {
		// Logged-in user
		cartPage, err = svc.BuildCartPageForUser(c.Request.Context(), u.ID)
	} else {
		// Guest user
		cc, _ := h.CartCK.Get(c)
		cartPage, err = svc.BuildCartPageFromCookie(c.Request.Context(), cc)
	}

	if err != nil {
		return view.CheckoutSummary{}, 0, "EUR", err
	}

	ship := shippingCents("standard")
	total := cartPage.SubtotalCents + ship

	return view.CheckoutSummary{
		Currency: cartPage.Currency,
		Subtotal: cartPage.Subtotal,
		Shipping: view.MoneyFromCents(ship, cartPage.Currency),
		Total:    view.MoneyFromCents(total, cartPage.Currency),
		Items:    cartPage.Count,
	}, cartPage.Count, cartPage.Currency, nil
}

func (h *CheckoutHandler) createTempCartFromCookie(c *gin.Context, cc *cartcookie.Cart) (*cartmod.Cart, error) {
	repo := cartmod.NewRepo(h.DB)

	// Create empty cart with UUID
	tempCart := cartmod.Cart{
		ID:     uuid.NewString(),
		UserID: nil,
	}
	if err := h.DB.Create(&tempCart).Error; err != nil {
		log.Printf("createTempCartFromCookie: failed to create cart: %v", err)
		return nil, err
	}

	// Add items from cookie
	for _, it := range cc.Items {
		if it.VariantID == "" || it.Qty <= 0 {
			continue
		}
		if err := repo.AddItem(c.Request.Context(), tempCart.ID, it.VariantID, it.Qty); err != nil {
			log.Printf("createTempCartFromCookie: failed to add item %s: %v", it.VariantID, err)
			return nil, err
		}
	}

	return &tempCart, nil
}

func (h *CheckoutHandler) renderCheckoutWithErrors(c *gin.Context, authed bool, summary view.CheckoutSummary, currency string, errs validation.FieldErrors, pageErr string, in checkoutInput) {
	form := view.CheckoutForm{
		Email:          in.Email,
		FirstName:      in.FirstName,
		LastName:       in.LastName,
		Address1:       in.Address1,
		Address2:       in.Address2,
		City:           in.City,
		PostalCode:     in.PostalCode,
		Country:        in.Country,
		Phone:          in.Phone,
		ShippingMethod: in.ShippingMethod,
		IdemKey:        in.IdemKey,
	}
	if form.ShippingMethod == "" {
		form.ShippingMethod = "standard"
	}
	if form.IdemKey == "" {
		form.IdemKey = randHex(16)
	}

	ship := shippingCents(form.ShippingMethod)
	summary.Shipping = view.MoneyFromCents(ship, currency)

	opts := shippingOptions(currency)

	render.Component(c, http.StatusBadRequest, pages.Checkout(
		middleware.GetFlash(c),
		middleware.GetCSRFToken(c),
		form,
		errs,
		pageErr,
		opts,
		summary,
		authed,
	))
}

func shippingOptions(currency string) []view.ShippingOption {
	return []view.ShippingOption{
		{Code: "standard", Label: "Standard (2-4 gün)", Price: view.MoneyFromCents(shippingCents("standard"), currency)},
		{Code: "express", Label: "Express (1-2 gün)", Price: view.MoneyFromCents(shippingCents("express"), currency)},
	}
}

func shippingCents(method string) int {
	switch method {
	case "express":
		return 1500
	default:
		return 500
	}
}
