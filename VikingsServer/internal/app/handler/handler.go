package handler

import (
	_ "VikingsServer/docs"
	"VikingsServer/internal/app/config"
	"VikingsServer/internal/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

const (
	baseURL    = "api/v3"
	citiesHTML = "cities"

	cities            = baseURL + "/cities"
	addCityImage      = baseURL + "/cities/upload-image"
	hikes             = baseURL + "/hikes"
	hikesUpdateStatus = baseURL + "/hikes-update-status"
	users             = baseURL + "/users"
	login             = baseURL + "/login"
	DestinationHikes  = baseURL + "/destination-hikes"
)

type Handler struct {
	Logger     *logrus.Logger
	Repository *repository.Repository
	Minio      *minio.Client
	Config     *config.Config
}

func NewHandler(l *logrus.Logger, r *repository.Repository, m *minio.Client, conf *config.Config) *Handler {
	return &Handler{
		Logger:     l,
		Repository: r,
		Minio:      m,
		Config:     conf,
	}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	//router.GET(citiesHTML, h.CitiesHTML)
	//router.POST(citiesHTML, h.DeleteCityHTML)

	router.GET(cities, h.CitiesList)
	router.POST(cities, h.AddCity)
	router.POST(addCityImage, h.AddImage)
	router.PUT(cities, h.UpdateCity)
	router.DELETE(cities, h.DeleteCity)

	router.GET(hikes, h.HikesList)
	//router.POST(hikes, h.AddHike)
	router.DELETE(hikes, h.DeleteHike)
	router.PUT(hikesUpdateStatus, h.UpdateHikeStatus)
	router.PUT(hikes, h.UpdateHike)

	router.GET(users, h.UsersList)
	router.POST(login, h.Login)

	router.GET(DestinationHikes, h.DestinationHikesList)
	router.POST(DestinationHikes, h.AddDestinationToHike)
	router.PUT(DestinationHikes, h.UpdateDestinationHikeNumber)
	router.DELETE(DestinationHikes, h.DeleteDestinationToHike)

	registerStatic(router)
}

func registerStatic(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.LoadHTMLGlob("static/html/*")
	router.Static("/static", "./static")
	router.Static("/css", "./static")
	router.Static("/img", "./static")
}

func registerFrontHeaders(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "http://localhost:5173")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
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
	registerFrontHeaders(ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		key:      data,
	})
}

func (h *Handler) successAddHandler(ctx *gin.Context, key string, data interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		key:      data,
	})
}
