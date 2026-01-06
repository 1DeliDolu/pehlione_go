package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"pehlione.com/app/internal/http/middleware"
	"pehlione.com/app/internal/http/render"
	"pehlione.com/app/internal/modules/currency"
	"pehlione.com/app/internal/modules/products"
	"pehlione.com/app/internal/modules/wishlist"
	"pehlione.com/app/templates/pages"
)

type WishlistHandler struct {
	wishlist *wishlist.Service
	products products.Repository
	currency *currency.Service
}

func NewWishlistHandler(wsvc *wishlist.Service, prodRepo products.Repository, currSvc *currency.Service) *WishlistHandler {
	return &WishlistHandler{
		wishlist: wsvc,
		products: prodRepo,
		currency: currSvc,
	}
}

func (h *WishlistHandler) List(c *gin.Context) {
	user, ok := middleware.CurrentUser(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	items, err := h.wishlist.Items(c.Request.Context(), user.ID)
	if err != nil {
		render.Component(c, http.StatusInternalServerError, pages.WishlistPage(pages.WishlistVM{
			Items:     nil,
			Message:   "Wishlist yüklenemedi.",
			CSRFToken: middleware.GetCSRFToken(c),
		}))
		return
	}

	productIDs := make([]string, 0, len(items))
	for _, it := range items {
		productIDs = append(productIDs, it.ProductID)
	}

	prods, err := h.products.ListByIDs(c.Request.Context(), productIDs)
	if err != nil {
		render.Component(c, http.StatusInternalServerError, pages.WishlistPage(pages.WishlistVM{
			Items:     nil,
			Message:   "Wishlist yüklenemedi.",
			CSRFToken: middleware.GetCSRFToken(c),
		}))
		return
	}

	displayCurrency := middleware.GetDisplayCurrency(c)
	cardMap := map[string]pages.WishlistItem{}
	for _, p := range prods {
		card := pages.WishlistItem{
			ProductID: p.ID,
			Title:     p.Name,
			Slug:      p.Slug,
			ImageURL:  "",
			Currency:  displayCurrency,
		}
		if len(p.Images) > 0 {
			card.ImageURL = p.Images[0].URL
		}
		if len(p.Variants) > 0 {
			price := int64(p.Variants[0].PriceCents)
			card.PriceCents = convertWishlistPrice(c.Request.Context(), h.currency, price, displayCurrency)
		}
		cardMap[p.ID] = card
	}

	viewItems := make([]pages.WishlistItem, 0, len(items))
	for _, it := range items {
		if card, ok := cardMap[it.ProductID]; ok {
			viewItems = append(viewItems, card)
		}
	}

	vm := pages.WishlistVM{
		Items:     viewItems,
		Message:   "",
		CSRFToken: middleware.GetCSRFToken(c),
	}
	render.Component(c, http.StatusOK, pages.WishlistPage(vm))
}

func (h *WishlistHandler) Add(c *gin.Context) {
	user, ok := middleware.CurrentUser(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	productID := strings.TrimSpace(c.PostForm("product_id"))
	if productID == "" {
		c.Redirect(http.StatusFound, c.Request.Referer())
		return
	}

	if err := h.wishlist.Add(c.Request.Context(), user.ID, productID); err != nil {
		// ignore duplicate errors
	}

	redirectBack(c, c.Request.Referer(), "/wishlist")
}

func (h *WishlistHandler) Remove(c *gin.Context) {
	user, ok := middleware.CurrentUser(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	productID := strings.TrimSpace(c.PostForm("product_id"))
	if productID == "" {
		c.Redirect(http.StatusFound, "/wishlist")
		return
	}

	_ = h.wishlist.Remove(c.Request.Context(), user.ID, productID)
	redirectBack(c, c.Request.Referer(), "/wishlist")
}

func redirectBack(c *gin.Context, referer, fallback string) {
	if referer != "" {
		c.Redirect(http.StatusFound, referer)
		return
	}
	c.Redirect(http.StatusFound, fallback)
}

func convertWishlistPrice(ctx context.Context, currSvc *currency.Service, cents int64, currency string) int64 {
	if currSvc == nil {
		return cents
	}
	converted, _, err := currSvc.ConvertDisplay(ctx, int(cents), currency)
	if err != nil {
		return cents
	}
	return int64(converted)
}
