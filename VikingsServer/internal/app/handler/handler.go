package handler

import (
	"VikingsServer/internal/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	baseURL    = "api/v3"
	citiesHTML = "cities"

	cities = baseURL + "/cities"
	hikes  = baseURL + "/hikes"
	viking = baseURL + "/vikings"
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
	router.POST(cities, h.AddCity)
	router.PUT(cities, h.UpdateCity)
	router.DELETE(cities, h.DeleteCity)

	router.GET(citiesHTML, h.CitiesHTML)
	router.POST(citiesHTML, h.DeleteCityHTML)

	router.GET(hikes, h.HikesList)
	router.POST(hikes, h.AddHike)
	router.DELETE(hikes, h.DeleteHike)
	router.PUT(hikes, h.UpdateHike)

	router.GET(viking, h.VikingsList)
	router.POST(viking, h.AddViking)
	router.PUT(viking, h.UpdateViking)

	registerStatic(router)
}

func registerStatic(router *gin.Engine) {
	router.LoadHTMLGlob("static/html/*")
	router.Static("/static", "./static")
	router.Static("/css", "./static")
	router.Static("/img", "./static")
}

// MARK: - Error handler

func (h *Handler) errorHandler(ctx *gin.Context, errorStatusCode int, err error) {
	h.Logger.Error(err.Error())
	ctx.JSON(errorStatusCode, gin.H{
		"status":      "error",
		"description": err.Error(),
	})
}

// MARK: - Success handler

func (h *Handler) successHandler(ctx *gin.Context, key string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		key:      data,
	})
}
