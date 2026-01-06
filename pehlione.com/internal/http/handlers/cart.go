package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"pehlione.com/app/internal/http/cartcookie"
	"pehlione.com/app/internal/http/flash"
	"pehlione.com/app/internal/http/middleware"
	"pehlione.com/app/internal/http/render"
	"pehlione.com/app/internal/modules/cart"
	"pehlione.com/app/pkg/view"
	"pehlione.com/app/templates/pages"
)

// CartHandler handles cart operations (GET /cart, POST /cart/add)
type CartHandler struct {
	DB    *gorm.DB
	Flash *flash.Codec
	CK    *cartcookie.Codec
}

func NewCartHandler(db *gorm.DB, flashCodec *flash.Codec, ck *cartcookie.Codec) *CartHandler {
	return &CartHandler{DB: db, Flash: flashCodec, CK: ck}
}

// Add handles POST /cart/add - adds item to cart and redirects to /cart
func (h *CartHandler) Add(c *gin.Context) {
	variantID := strings.TrimSpace(c.PostForm("variant_id"))
	qtyStr := strings.TrimSpace(c.PostForm("qty"))

	qty := 1
	if qtyStr != "" {
		if n, err := strconv.Atoi(qtyStr); err == nil && n > 0 && n <= 99 {
			qty = n
		}
	}

	if variantID == "" {
		render.RedirectWithFlash(c, h.Flash, "/products", view.FlashError, "Variant seçilemedi.")
		return
	}

	// Check if user is logged in
	if u, ok := middleware.CurrentUser(c); ok && u.ID != "" {
		// Logged-in user: add to DB cart
		cartRepo := cart.NewRepo(h.DB)

		// Get or create user's cart
		userCart, err := cartRepo.GetOrCreateUserCart(c, u.ID)
		if err != nil {
			log.Printf("CartAdd: error getting cart: %v", err)
			render.RedirectWithFlash(c, h.Flash, "/products", view.FlashError, "Sepete ekleme başarısız.")
			return
		}

		// Add item to cart
		if err := cartRepo.AddItem(c, userCart.ID, variantID, qty); err != nil {
			log.Printf("CartAdd: error adding item: %v", err)
			render.RedirectWithFlash(c, h.Flash, "/products", view.FlashError, "Sepete ekleme başarısız.")
			return
		}

		// Clear cache
		middleware.ClearSessionCartCache(c)

		render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashSuccess, "✓ Sepete eklendi.")
		return
	}

	// Guest user: add to cookie cart
	cc, _ := h.CK.Get(c)
	if cc == nil {
		cc = &cartcookie.Cart{}
	}
	cc.AddItem(variantID, qty)
	h.CK.Set(c, cc)

	render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashSuccess, "✓ Sepete eklendi.")
}

// Update handles POST /cart/items/update - updates item quantity
func (h *CartHandler) Update(c *gin.Context) {
	variantID := strings.TrimSpace(c.PostForm("variant_id"))
	qtyStr := strings.TrimSpace(c.PostForm("qty"))
	qty := 1
	if qtyStr != "" {
		if n, err := strconv.Atoi(qtyStr); err == nil {
			qty = n
		}
	}

	if variantID == "" {
		render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashError, "Ürün bulunamadı.")
		return
	}

	qty = clamp(qty, 0, 99)

	if u, ok := middleware.CurrentUser(c); ok && u.ID != "" {
		repo := cart.NewRepo(h.DB)
		userCart, err := repo.GetOrCreateUserCart(c.Request.Context(), u.ID)
		if err != nil {
			log.Printf("CartUpdate: error getting cart: %v", err)
			render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashError, "Sepet güncellenemedi.")
			return
		}

		if err := repo.UpdateItemQty(c.Request.Context(), userCart.ID, variantID, qty); err != nil {
			log.Printf("CartUpdate: update item error: %v", err)
			render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashError, "Miktar güncellenemedi.")
			return
		}

		middleware.ClearSessionCartCache(c)
		render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashSuccess, "Miktar güncellendi.")
		return
	}

	cc, _ := h.CK.Get(c)
	if cc == nil {
		cc = &cartcookie.Cart{}
	}
	cc.UpdateQuantity(variantID, qty)
	h.CK.Set(c, cc)
	render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashSuccess, "Miktar güncellendi.")
}

// Remove handles POST /cart/items/remove - removes item from cart
func (h *CartHandler) Remove(c *gin.Context) {
	variantID := strings.TrimSpace(c.PostForm("variant_id"))
	if variantID == "" {
		render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashWarning, "Ürün bulunamadı.")
		return
	}

	if u, ok := middleware.CurrentUser(c); ok && u.ID != "" {
		repo := cart.NewRepo(h.DB)
		userCart, err := repo.GetOrCreateUserCart(c.Request.Context(), u.ID)
		if err != nil {
			log.Printf("CartRemove: error getting cart: %v", err)
			render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashError, "Sepet güncellenemedi.")
			return
		}

		if err := repo.RemoveItem(c.Request.Context(), userCart.ID, variantID); err != nil {
			log.Printf("CartRemove: remove item error: %v", err)
			render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashError, "Ürün silinemedi.")
			return
		}

		middleware.ClearSessionCartCache(c)
		render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashSuccess, "Ürün sepetten çıkarıldı.")
		return
	}

	cc, _ := h.CK.Get(c)
	if cc == nil {
		cc = &cartcookie.Cart{}
	}
	cc.RemoveItem(variantID)
	h.CK.Set(c, cc)
	render.RedirectWithFlash(c, h.Flash, "/cart", view.FlashSuccess, "Ürün sepetten çıkarıldı.")
}

// Get handles GET /cart - displays cart page
func (h *CartHandler) Get(c *gin.Context) {
	flash := middleware.GetFlash(c)
	svc := cart.NewService(h.DB)

	// Check if user is logged in
	if u, ok := middleware.CurrentUser(c); ok {
		// Logged-in user: fetch from DB
		cartPage, err := svc.BuildCartPageForUser(c, u.ID)
		if err != nil {
			log.Printf("CartGet: error building page: %v", err)
			render.Component(c, http.StatusOK, pages.Cart(flash, view.CartPage{Items: []view.CartItem{}}))
			return
		}
		cartPage.CSRFToken = csrfTokenFrom(c)
		render.Component(c, http.StatusOK, pages.Cart(flash, cartPage))
		return
	}

	// Guest user: fetch from cookie
	cc, _ := h.CK.Get(c)
	cartPage, err := svc.BuildCartPageFromCookie(c, cc)
	if err != nil {
		log.Printf("CartGet: error building guest cart: %v", err)
		render.Component(c, http.StatusOK, pages.Cart(flash, view.CartPage{Items: []view.CartItem{}}))
		return
	}
	cartPage.CSRFToken = csrfTokenFrom(c)

	render.Component(c, http.StatusOK, pages.Cart(flash, cartPage))
}

func clamp(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}
