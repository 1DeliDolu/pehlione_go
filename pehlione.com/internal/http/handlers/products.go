package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"pehlione.com/app/internal/http/middleware"
	"pehlione.com/app/internal/http/render"
	"pehlione.com/app/internal/modules/currency"
	"pehlione.com/app/internal/modules/products"
	pages "pehlione.com/app/templates/pages/products"
)

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
	_ = json.Unmarshal(jsonData, &opts)
	return opts
}

type ProductsHandler struct {
	svc      *products.Service
	currency *currency.Service
}

func NewProductsHandler(svc *products.Service, curr *currency.Service) *ProductsHandler {
	return &ProductsHandler{svc: svc, currency: curr}
}

func (h *ProductsHandler) List(c *gin.Context) {
	queryVals := c.Request.URL.Query()
	listFilters, uiState := buildListFilters(queryVals)

	result, err := h.svc.ListWithFilters(c.Request.Context(), listFilters)
	if err != nil {
		render.Component(c, http.StatusInternalServerError, pages.ProductsIndexPage(pages.ProductsIndexVM{
			Title:      "Products",
			AlertError: "Product list could not be loaded.",
			Products:   []pages.ProductCardVM{},
			CSRFToken:  csrfTokenFrom(c),
			Filters: pages.ProductsFilterVM{
				Query: uiState.Query,
			},
			Pagination: pages.PaginationVM{Page: 1, TotalPages: 1},
		}))
		return
	}

	displayCurrency := middleware.GetDisplayCurrency(c)
	productsVM := mapProductsForList(c.Request.Context(), result.Items, displayCurrency, h.currency)
	filterVM := buildFilterVM(result, uiState)
	pagination := buildPaginationVM(result, queryVals, c.Request.URL.Path)

	vm := pages.ProductsIndexVM{
		Title:      "Products",
		Products:   productsVM,
		CSRFToken:  csrfTokenFrom(c),
		Filters:    filterVM,
		Pagination: pagination,
		Total:      result.Total,
	}
	render.Component(c, http.StatusOK, pages.ProductsIndexPage(vm))
}

func (h *ProductsHandler) Show(c *gin.Context) {
	slugParam := c.Param("slug")

	p, err := h.svc.Detail(c.Request.Context(), slugParam)
	if err != nil {
		c.Status(http.StatusNotFound)
		render.Component(c, http.StatusNotFound, pages.ProductsNotFoundPage(pages.SimpleVM{
			Title: "Product not found",
		}))
		return
	}

	displayCurrency := middleware.GetDisplayCurrency(c)
	pd := mapProductForDetail(c.Request.Context(), p, displayCurrency, h.currency)

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

type filterState struct {
	Query       string
	Category    string
	MinPriceStr string
	MaxPriceStr string
	InStock     bool
	Sort        string
	Filters     products.ListFilters
}

func buildListFilters(values url.Values) (products.ListFilters, filterState) {
	state := filterState{}
	filters := products.ListFilters{
		PageSize: 24,
		Page:     1,
		Sort:     "newest",
	}

	if pageStr := strings.TrimSpace(values.Get("page")); pageStr != "" {
		if n, err := strconv.Atoi(pageStr); err == nil && n > 0 {
			filters.Page = n
		}
	}

	if q := strings.TrimSpace(values.Get("q")); q != "" {
		state.Query = q
		filters.Query = q
	}

	if cat := strings.TrimSpace(values.Get("category")); cat != "" {
		state.Category = cat
		filters.Category = cat
	}

	minCents, minStr := parsePriceInput(values.Get("min_price"))
	maxCents, maxStr := parsePriceInput(values.Get("max_price"))
	state.MinPriceStr = minStr
	state.MaxPriceStr = maxStr
	filters.MinPrice = minCents
	filters.MaxPrice = maxCents
	if filters.MaxPrice > 0 && filters.MinPrice > filters.MaxPrice {
		filters.MinPrice, filters.MaxPrice = filters.MaxPrice, filters.MinPrice
		state.MinPriceStr, state.MaxPriceStr = state.MaxPriceStr, state.MinPriceStr
	}

	if values.Get("in_stock") == "1" {
		state.InStock = true
		filters.InStock = true
	}

	switch values.Get("sort") {
	case "price_asc":
		filters.Sort = "price_asc"
		state.Sort = "price_asc"
	case "price_desc":
		filters.Sort = "price_desc"
		state.Sort = "price_desc"
	default:
		filters.Sort = "newest"
		state.Sort = "newest"
	}

	state.Filters = filters
	return filters, state
}

func buildFilterVM(res products.ListResult, state filterState) pages.ProductsFilterVM {
	categoryOptions := make([]pages.CategoryOptionVM, 0, len(res.Categories))
	for _, cat := range res.Categories {
		value := cat.Slug
		if value == "all" {
			value = ""
		}
		categoryOptions = append(categoryOptions, pages.CategoryOptionVM{
			Label:    cat.Name,
			Value:    value,
			Count:    cat.Count,
			Selected: value != "" && value == state.Category,
		})
	}

	sortOptions := []pages.SortOptionVM{
		{Label: "Newest arrivals", Value: "newest", Selected: state.Sort == "newest"},
		{Label: "Price: low to high", Value: "price_asc", Selected: state.Sort == "price_asc"},
		{Label: "Price: high to low", Value: "price_desc", Selected: state.Sort == "price_desc"},
	}

	return pages.ProductsFilterVM{
		Query:       state.Query,
		Category:    state.Category,
		MinPrice:    state.MinPriceStr,
		MaxPrice:    state.MaxPriceStr,
		InStock:     state.InStock,
		Sort:        state.Sort,
		Categories:  categoryOptions,
		SortOptions: sortOptions,
	}
}

func buildPaginationVM(res products.ListResult, vals url.Values, path string) pages.PaginationVM {
	totalPages := 1
	if res.PageSize > 0 {
		totalPages = int(math.Ceil(float64(res.Total) / float64(res.PageSize)))
		if totalPages < 1 {
			totalPages = 1
		}
	}

	pagination := pages.PaginationVM{
		Page:       res.Page,
		TotalPages: totalPages,
	}

	if res.Page > 1 {
		pagination.HasPrev = true
		pagination.PrevURL = buildPageURL(path, vals, res.Page-1)
	}

	if res.Page < totalPages {
		pagination.HasNext = true
		pagination.NextURL = buildPageURL(path, vals, res.Page+1)
	}

	return pagination
}

func buildPageURL(path string, vals url.Values, page int) string {
	cloned := cloneValues(vals)
	if page <= 1 {
		cloned.Del("page")
	} else {
		cloned.Set("page", strconv.Itoa(page))
	}
	if len(cloned) == 0 {
		return path
	}
	return path + "?" + cloned.Encode()
}

func cloneValues(vals url.Values) url.Values {
	out := url.Values{}
	for k, v := range vals {
		for _, s := range v {
			out.Add(k, s)
		}
	}
	return out
}

func parsePriceInput(val string) (int, string) {
	val = strings.TrimSpace(val)
	if val == "" {
		return 0, ""
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil || f < 0 {
		return 0, val
	}
	cents := int(math.Round(f * 100))
	return cents, val
}

func mapProductsForList(ctx context.Context, items []products.Product, displayCurrency string, currSvc *currency.Service) []pages.ProductCardVM {
	vm := make([]pages.ProductCardVM, 0, len(items))
	for _, p := range items {
		img := ""
		if len(p.Images) > 0 {
			img = p.Images[0].URL
		}

		minPrice := int64(0)
		defaultVariantID := ""
		bestVariantID := ""

		for idx, v := range p.Variants {
			price := int64(v.PriceCents)
			if idx == 0 || price < minPrice || minPrice == 0 {
				minPrice = price
				bestVariantID = v.ID
			}
			if v.Stock > 0 && defaultVariantID == "" {
				defaultVariantID = v.ID
			}
		}

		if defaultVariantID == "" {
			defaultVariantID = bestVariantID
		}

		convertedMin := convertPriceValue(ctx, currSvc, minPrice, displayCurrency)

		vm = append(vm, pages.ProductCardVM{
			ProductID:        p.ID,
			Title:            p.Name,
			Slug:             p.Slug,
			ImageURL:         img,
			PriceCents:       convertedMin,
			Currency:         displayCurrency,
			DefaultVariantID: defaultVariantID,
			Subtitle:         p.CategoryName,
		})
	}
	return vm
}

func mapProductForDetail(ctx context.Context, p products.Product, displayCurrency string, currSvc *currency.Service) pages.ProductDetailVM {
	imgs := make([]string, 0, len(p.Images))
	for _, im := range p.Images {
		imgs = append(imgs, im.URL)
	}

	var price int64
	defaultVariantID := ""
	var defaultColor, defaultSize string

	if len(p.Variants) > 0 {
		price = int64(p.Variants[0].PriceCents)
		defaultVariantID = p.Variants[0].ID

		opts := parseVariantOptions(p.Variants[0].Options)
		defaultColor = opts.Color
		defaultSize = opts.Size
	}

	colorsSet := map[string]struct{}{}
	sizesSet := map[string]struct{}{}
	variants := make([]pages.VariantVM, 0, len(p.Variants))

	for _, vv := range p.Variants {
		opts := parseVariantOptions(vv.Options)

		if opts.Color != "" {
			colorsSet[opts.Color] = struct{}{}
		}
		if opts.Size != "" {
			sizesSet[opts.Size] = struct{}{}
		}

		priceCents := convertPriceValue(ctx, currSvc, int64(vv.PriceCents), displayCurrency)
		compareCents := convertPriceValue(ctx, currSvc, int64(vv.CompareAtCents), displayCurrency)

		variants = append(variants, pages.VariantVM{
			ID:             vv.ID,
			Color:          opts.Color,
			Size:           opts.Size,
			PriceCents:     priceCents,
			CompareAtCents: compareCents,
			StockQty:       vv.Stock,
			IsDefault:      vv.ID == defaultVariantID,
		})
	}

	colors := make([]string, 0, len(colorsSet))
	for k := range colorsSet {
		colors = append(colors, k)
	}

	sizes := make([]string, 0, len(sizesSet))
	for k := range sizesSet {
		sizes = append(sizes, k)
	}

	convertedDefault := convertPriceValue(ctx, currSvc, price, displayCurrency)

	return pages.ProductDetailVM{
		ID:               p.ID,
		Slug:             p.Slug,
		Title:            p.Name,
		Description:      strings.TrimSpace(p.Description),
		Images:           imgs,
		PriceCents:       convertedDefault,
		Currency:         displayCurrency,
		Colors:           colors,
		Sizes:            sizes,
		Variants:         variants,
		DefaultVariantID: defaultVariantID,
		DefaultColor:     defaultColor,
		DefaultSize:      defaultSize,
	}

}

func convertPriceValue(ctx context.Context, currSvc *currency.Service, cents int64, currency string) int64 {
	if currSvc == nil {
		return cents
	}
	converted, _, err := currSvc.ConvertDisplay(ctx, int(cents), currency)
	if err != nil {
		return cents
	}
	return int64(converted)
}
