package handler

import (
	"github.com/gin-gonic/gin"
)

const (
	baseURL    = "api/v3"
	citiesHTML = "cities"
	cities     = baseURL + "/cities"
	hikes      = baseURL + "/hikes"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.GET(citiesHTML, h.CitiesHTML)

	registerStatic(router)
}

func registerStatic(router *gin.Engine) {
	router.LoadHTMLGlob("static/html/*")
	router.Static("/static", "./static")
	router.Static("/css", "./static")
	router.Static("/img", "./static")
}
