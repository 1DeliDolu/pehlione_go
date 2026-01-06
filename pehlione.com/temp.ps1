
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
