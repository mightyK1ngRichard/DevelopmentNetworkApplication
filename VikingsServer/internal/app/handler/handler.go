package handler

import (
	"VikingsServer/internal/app/repository"
	"github.com/gin-gonic/gin"
)

const (
	baseURL    = "api/v3"
	citiesHTML = "cities"
	cities     = baseURL + "/cities"
	hikes      = baseURL + "/hikes"
)

type Handler struct {
	Repository *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{
		Repository: r,
	}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.GET(cities, h.CitiesList)
	router.GET(hikes, h.HikesList)
	router.GET(citiesHTML, h.CitiesHTML)

	registerStatic(router)
}

func registerStatic(router *gin.Engine) {
	router.LoadHTMLGlob("static/html/*")
	router.Static("/static", "./static")
	router.Static("/css", "./static")
	router.Static("/img", "./static")
}
