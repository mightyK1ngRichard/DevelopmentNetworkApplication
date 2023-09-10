package handler

import (
	"VikingsServer/internal/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	baseURL    = "api/v3"
	citiesHTML = "cities"
	cities     = baseURL + "/cities"
	hikes      = baseURL + "/hikes"
)

type Handler struct {
	Logger     *logrus.Logger
	Repository *repository.Repository
}

func NewHandler(l *logrus.Logger, r *repository.Repository) *Handler {
	return &Handler{
		Logger:     l,
		Repository: r,
	}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.GET(cities, h.CitiesList)
	router.GET(hikes, h.HikesList)
	router.GET(citiesHTML, h.CitiesHTML)
	router.POST(citiesHTML, h.DeleteCityWithStatus)
	router.DELETE(citiesHTML, h.CitiesDelete)

	registerStatic(router)
}

func registerStatic(router *gin.Engine) {
	router.LoadHTMLGlob("static/html/*")
	router.Static("/static", "./static")
	router.Static("/css", "./static")
	router.Static("/img", "./static")
}
