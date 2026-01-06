package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pehlione.com/app/internal/http/middleware"
	"pehlione.com/app/internal/http/render"
	"pehlione.com/app/templates/pages"
)

func Home(c *gin.Context) {
	flash := middleware.GetFlash(c)
	headerCtx := middleware.BuildHeaderCtx(c)
	render.Component(c, http.StatusOK, pages.Home(flash, headerCtx))
}
