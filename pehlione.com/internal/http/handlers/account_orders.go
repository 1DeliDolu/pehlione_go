package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"pehlione.com/app/internal/http/flash"
	"pehlione.com/app/internal/http/middleware"
	"pehlione.com/app/internal/http/render"
	"pehlione.com/app/internal/modules/auth"
	"pehlione.com/app/internal/modules/orders"
	"pehlione.com/app/internal/shared/apperr"
	"pehlione.com/app/pkg/view"
	"pehlione.com/app/templates/pages"
)

type AccountOrdersHandler struct {
	ordersRepo *orders.Repo
	authRepo   *auth.Repo
	Flash      *flash.Codec
}

func NewAccountOrdersHandler(ordersRepo *orders.Repo, authRepo *auth.Repo, flashCodec *flash.Codec) *AccountOrdersHandler {
	return &AccountOrdersHandler{ordersRepo: ordersRepo, authRepo: authRepo, Flash: flashCodec}
}

func (h *AccountOrdersHandler) buildPage(c *gin.Context, user auth.User, page, pageSize int, status string, passwordErrors map[string]string) (view.AccountOrdersPage, error) {
	offset := (page - 1) * pageSize

	var items []view.AccountOrderListItem
	var total int64

	query := h.ordersRepo.DB().WithContext(c.Request.Context()).Where("user_id = ?", user.ID)
	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return view.AccountOrdersPage{}, err
	}

	var dbOrders []orders.Order
	if err := query.Order("created_at DESC").Limit(pageSize).Offset(offset).Find(&dbOrders).Error; err != nil {
		return view.AccountOrdersPage{}, err
	}

	for _, o := range dbOrders {
		orderNum := o.ID
		if len(o.ID) > 8 {
			orderNum = o.ID[:8]
		}
		items = append(items, view.AccountOrderListItem{
			ID:         o.ID,
			Number:     orderNum,
			CreatedAt:  o.CreatedAt,
			Status:     o.Status,
			TotalCents: o.TotalCents,
			Currency:   o.Currency,
			ItemCount:  0, // TODO: count items if needed
			PaidAt:     o.PaidAt,
		})
	}

	return view.AccountOrdersPage{
		Account: view.AccountInfo{
			Email:     user.Email,
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
			Verified:  user.EmailVerifiedAt != nil,
		},
		Items:          items,
		Total:          total,
		Page:           page,
		PageSize:       pageSize,
		FilterStatus:   status,
		Statuses:       []string{"pending", "paid", "shipped", "delivered", "cancelled"},
		IsPreviousPage: page > 1,
		IsNextPage:     offset+pageSize < int(total),
		CSRFToken:      middleware.GetCSRFToken(c),
		PasswordErrors: passwordErrors,
	}, nil
}

func (h *AccountOrdersHandler) List(c *gin.Context) {
	user, ok := middleware.CurrentUser(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}

	status := c.Query("status")
	const pageSize = 20

	fullUser, err := h.authRepo.GetByID(c.Request.Context(), user.ID)
	if err != nil {
		c.Error(apperr.Wrap(err))
		return
	}

	pageView, err := h.buildPage(c, *fullUser, page, pageSize, status, nil)
	if err != nil {
		c.Error(apperr.Wrap(err))
		return
	}

	render.Component(c, http.StatusOK, pages.AccountOrders(
		middleware.GetFlash(c),
		pageView,
	))
}

func (h *AccountOrdersHandler) ChangePassword(c *gin.Context) {
	userCtx, ok := middleware.CurrentUser(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	pageNum, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	if pageNum < 1 {
		pageNum = 1
	}
	status := c.PostForm("status")

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
		pageView, buildErr := h.buildPage(c, *fullUser, pageNum, 20, status, errorsMap)
		if buildErr != nil {
			c.Error(apperr.Wrap(buildErr))
			return
		}
		render.Component(c, http.StatusBadRequest, pages.AccountOrders(
			middleware.GetFlash(c),
			pageView,
		))
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
	render.RedirectWithFlash(c, flashCodec, "/account/orders", view.FlashSuccess, "Şifreniz güncellendi.")
}
