package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"pehlione.com/app/internal/http/render"
	"pehlione.com/app/internal/modules/products"
	"pehlione.com/app/internal/shared/slug"
	pages "pehlione.com/app/templates/pages/products"
)

// VariantOptions Options JSON yapısını parse etmek için
type VariantOptions struct {
	Color string `json:"color"`
	Size  string `json:"size"`
}

type variantData struct {
	ID             string `json:"id"`
	Color          string `json:"color"`
	Size           string `json:"size"`
	PriceCents     int64  `json:"priceCents"`
	CompareAtCents int64  `json:"compareAtCents"`
	StockQty       int    `json:"stockQty"`
}

func parseVariantOptions(jsonData []byte) VariantOptions {
	var opts VariantOptions
	json.Unmarshal(jsonData, &opts)
	return opts
}

// ProductsHandler handles product listing and detail
type ProductsHandler struct {
	svc *products.Service
}

func NewProductsHandler(svc *products.Service) *ProductsHandler {
	return &ProductsHandler{svc: svc}
}

// List returns the products listing page
func (h *ProductsHandler) List(c *gin.Context) {
	limit := 24
	if v := c.Query("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 && n <= 100 {
			limit = n
		}
	}

	prods, err := h.svc.List(c.Request.Context(), limit, 0)
	if err != nil {
		render.Component(c, http.StatusInternalServerError, pages.ProductsIndexPage(pages.ProductsIndexVM{
			Title:      "Ürünler",
			AlertError: "Ürün listesi yüklenemedi.",
			Products:   []pages.ProductCardVM{},
			CSRFToken:  csrfTokenFrom(c),
		}))
		return
	}

	groups := groupProductsByCategory(prods)

	vm := pages.ProductsIndexVM{
		Title:          "Ürünler",
		Products:       mapProductsForList(prods),
		CategoryGroups: groups,
		CSRFToken:      csrfTokenFrom(c),
	}
	render.Component(c, http.StatusOK, pages.ProductsIndexPage(vm))
}

// Show returns the product detail page
func (h *ProductsHandler) Show(c *gin.Context) {
	slug := c.Param("slug")

	p, err := h.svc.Detail(c.Request.Context(), slug)
	if err != nil {
		c.Status(http.StatusNotFound)
		render.Component(c, http.StatusNotFound, pages.ProductsNotFoundPage(pages.SimpleVM{
			Title: "Ürün bulunamadı",
		}))
		return
	}

	pd := mapProductForDetail(p)

	data := make([]variantData, 0, len(pd.Variants))
	for _, v := range pd.Variants {
		data = append(data, variantData{
			ID:             v.ID,
			Color:          v.Color,
			Size:           v.Size,
			PriceCents:     v.PriceCents,
			CompareAtCents: v.CompareAtCents,
			StockQty:       v.StockQty,
		})
	}

	variantJSON, err := json.Marshal(data)
	if err != nil {
		variantJSON = []byte("[]")
	}
	variantsB64 := base64.StdEncoding.EncodeToString(variantJSON)

	vm := pages.ProductsShowVM{
		Title:       p.Name,
		Product:     pd,
		CSRFToken:   csrfTokenFrom(c),
		VariantsB64: variantsB64,
	}
	render.Component(c, http.StatusOK, pages.ProductsShowPage(vm))
}

func csrfTokenFrom(c *gin.Context) string {
	if token, ok := c.Get("csrf_token"); ok {
		if str, ok := token.(string); ok {
			return str
		}
	}
	return ""
}

func mapProductsForList(items []products.Product) []pages.ProductCardVM {
	out := make([]pages.ProductCardVM, 0, len(items))
	for _, p := range items {
		out = append(out, productCardFromModel(p))
	}
	return out
}

func groupProductsByCategory(items []products.Product) []pages.ProductGroupVM {
	groupsBySlug := map[string]*pages.ProductGroupVM{}
	order := make([]string, 0)
	for _, p := range items {
		name, tagline := parseCategoryFromDescription(p.Description)
		catSlug := slug.FromName(strings.ToLower(name))
		if catSlug == "" {
			catSlug = slug.FromName("tüm ürünler")
		}
		group, ok := groupsBySlug[catSlug]
		if !ok {
			group = &pages.ProductGroupVM{
				Name:        name,
				Slug:        catSlug,
				Description: tagline,
				Products:    []pages.ProductCardVM{},
			}
			groupsBySlug[catSlug] = group
			order = append(order, catSlug)
		}
		if len(group.Products) < 3 {
			group.Products = append(group.Products, productCardFromModel(p))
		}
	}

	out := make([]pages.ProductGroupVM, 0, len(order))
	for _, slugKey := range order {
		out = append(out, *groupsBySlug[slugKey])
	}
	return out
}

func parseCategoryFromDescription(desc string) (string, string) {
	desc = strings.TrimSpace(desc)
	if strings.HasPrefix(strings.ToLower(desc), "category:") {
		body := strings.TrimSpace(desc[len("category:"):])
		parts := strings.SplitN(body, "|", 2)
		name := strings.TrimSpace(parts[0])
		if name == "" {
			name = "Tüm Ürünler"
		}
		tagline := ""
		if len(parts) > 1 {
			tagline = strings.TrimSpace(parts[1])
		}
		return name, tagline
	}
	return "Tüm Ürünler", desc
}

func productCardFromModel(p products.Product) pages.ProductCardVM {
	img := ""
	if len(p.Images) > 0 {
		img = p.Images[0].URL
	}

	price := int64(0)
	defaultVariantID := ""
	if len(p.Variants) > 0 {
		price = int64(p.Variants[0].PriceCents)
		defaultVariantID = p.Variants[0].ID
	}

	return pages.ProductCardVM{
		Title:            p.Name,
		Slug:             p.Slug,
		ImageURL:         img,
		PriceCents:       price,
		Currency:         "EUR",
		DefaultVariantID: defaultVariantID,
		Subtitle:         "",
	}
}

func mapProductForDetail(p products.Product) pages.ProductDetailVM {
	imgs := make([]string, 0, len(p.Images))
	for _, im := range p.Images {
		imgs = append(imgs, im.URL)
	}

	// İlk fiyat ve default variant options
	var price int64
	defaultVariantID := ""
	var defaultColor, defaultSize string

	if len(p.Variants) > 0 {
		price = int64(p.Variants[0].PriceCents)
		defaultVariantID = p.Variants[0].ID

		// Default variant'ın options'ını parse et
		opts := parseVariantOptions(p.Variants[0].Options)
		defaultColor = opts.Color
		defaultSize = opts.Size
	}

	// Renk ve beden setleri
	colorsSet := map[string]struct{}{}
	sizesSet := map[string]struct{}{}
	variants := make([]pages.VariantVM, 0, len(p.Variants))

	for _, vv := range p.Variants {
		opts := parseVariantOptions(vv.Options)

		colorsSet[opts.Color] = struct{}{}
		sizesSet[opts.Size] = struct{}{}

		variants = append(variants, pages.VariantVM{
			ID:             vv.ID,
			Color:          opts.Color,
			Size:           opts.Size,
			PriceCents:     int64(vv.PriceCents),
			CompareAtCents: int64(vv.CompareAtCents),
			StockQty:       vv.Stock,
			IsDefault:      vv.ID == defaultVariantID,
		})
	}

	colors := make([]string, 0, len(colorsSet))
	for k := range colorsSet {
		if k != "" {
			colors = append(colors, k)
		}
	}

	sizes := make([]string, 0, len(sizesSet))
	for k := range sizesSet {
		if k != "" {
			sizes = append(sizes, k)
		}
	}

	return pages.ProductDetailVM{
		ID:               p.ID,
		Slug:             p.Slug,
		Title:            p.Name,
		Description:      strings.TrimSpace(p.Description),
		Images:           imgs,
		PriceCents:       price,
		Currency:         "EUR",
		Colors:           colors,
		Sizes:            sizes,
		Variants:         variants,
		DefaultVariantID: defaultVariantID,
		DefaultColor:     defaultColor,
		DefaultSize:      defaultSize,
	}
}
