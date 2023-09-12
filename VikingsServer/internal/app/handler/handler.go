package handler

import (
	"VikingsServer/internal/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	baseURL    = "api/v3"
	citiesHTML = "cities"

	cities       = baseURL + "/cities"
	hikes        = baseURL + "/hikes"
	vikingAdd    = baseURL + "/viking/add"
	vikingUpdate = baseURL + "/viking/update"
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
	router.GET(citiesHTML, h.CitiesHTML)
	router.POST(citiesHTML, h.DeleteCityWithStatus)
	router.DELETE(citiesHTML, h.CitiesDelete)
	router.GET(hikes, h.HikesList)
	router.POST(vikingAdd, h.AddViking)
	router.PUT(vikingUpdate, h.UpdateViking)

	registerStatic(router)
}

func registerStatic(router *gin.Engine) {
	router.LoadHTMLGlob("static/html/*")
	router.Static("/static", "./static")
	router.Static("/css", "./static")
	router.Static("/img", "./static")
}
