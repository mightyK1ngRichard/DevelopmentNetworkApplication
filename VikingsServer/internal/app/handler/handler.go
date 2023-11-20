//go:build !appengine && !appenginevm

package handler

import (
	_ "VikingsServer/docs"
	"VikingsServer/internal/app/config"
	"VikingsServer/internal/app/redis"
	"VikingsServer/internal/app/repository"
	"VikingsServer/internal/app/role"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

const (
	baseURL         = "api/v3"
	citiesHTML      = "cities"
	cities          = baseURL + "/cities"
	addCityIntoHike = baseURL + "/cities/add-city-into-hike"
	addCityImage    = baseURL + "/cities/upload-image"

	hikesUpdateStatus            = baseURL + "/hikes-update-status"
	hikes                        = baseURL + "/hikes"
	hikeUpdateStatusForModerator = baseURL + "/hikes/update/status-for-moderator"
	hikeUpdateStatusForUser      = baseURL + "/hikes/update/status-for-user"

	users            = baseURL + "/users"
	login            = users + "/login"
	signup           = users + "/sign_up"
	logout           = users + "/logout"
	destinationHikes = baseURL + "/destination-hikes"
)

type Handler struct {
	Logger     *logrus.Logger
	Repository *repository.Repository
	Minio      *minio.Client
	Config     *config.Config
	Redis      *redis.Client
}

func NewHandler(
	l *logrus.Logger,
	r *repository.Repository,
	m *minio.Client,
	conf *config.Config,
	red *redis.Client,
) *Handler {
	return &Handler{
		Logger:     l,
		Repository: r,
		Minio:      m,
		Config:     conf,
		Redis:      red,
	}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h.UserCRUD(router)
	h.CityCRUD(router)
	h.HikeCRUD(router)
	registerStatic(router)
}

func (h *Handler) CityCRUD(router *gin.Engine) {
	router.GET(cities, h.CitiesList)
	router.POST(cities, h.WithAuthCheck(role.Manager, role.Admin), h.AddCity)
	router.PUT(addCityImage, h.AddImage)
	router.PUT(cities, h.WithAuthCheck(role.Manager, role.Admin), h.UpdateCity)
	router.DELETE(cities, h.WithAuthCheck(role.Manager, role.Admin), h.DeleteCity)
	//router.POST(addCityIntoHike, h.AddCityIntoHike)
	router.POST(addCityIntoHike, h.WithAuthCheck(role.Manager, role.Admin), h.AddCityIntoHike)
	router.GET("/ping", h.WithAuthCheck(role.Manager, role.Admin), h.Ping)
}

func (h *Handler) HikeCRUD(router *gin.Engine) {
	//router.POST(hikes, h.AddHike)
	//router.PUT(hikesUpdateStatus, h.UpdateHikeStatus)
	router.GET(hikes, h.HikesList)
	router.DELETE(hikes, h.WithAuthCheck(role.Manager, role.Admin), h.DeleteHike)
	router.PUT(hikeUpdateStatusForModerator, h.WithAuthCheck(role.Manager, role.Admin), h.UpdateStatusForModerator)
	router.PUT(hikeUpdateStatusForUser, h.WithAuthCheck(role.Manager, role.Admin), h.UpdateStatusForUser)
	router.PUT(hikes, h.WithAuthCheck(role.Manager, role.Admin), h.UpdateHike)
}

func (h *Handler) UserCRUD(router *gin.Engine) {
	//router.GET(users, h.UsersList)
	router.Use(cors.Default()).DELETE("/api/v3/cities/delete/:id", h.DeleteCityWithParam)
	router.POST(login, h.Login)
	router.POST(signup, h.Register)
	router.GET(logout, h.Logout)
}

func (h *Handler) DestinationHikesCRUD(router *gin.Engine) {
	router.GET(destinationHikes, h.DestinationHikesList)
	router.POST(destinationHikes, h.AddDestinationToHike)
	router.PUT(destinationHikes, h.UpdateDestinationHikeNumber)
	router.DELETE(destinationHikes, h.DeleteDestinationToHike)
}

func registerStatic(router *gin.Engine) {
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
type errorResp struct {
	Status      string `json:"status" example:"error"`
	Description string `json:"description" example:"Описание ошибки"`
}

func (h *Handler) errorHandler(ctx *gin.Context, errorStatusCode int, err error) {
	h.Logger.Error(err.Error())
	ctx.JSON(errorStatusCode, gin.H{
		"status":      "error",
		"description": err.Error(),
	})
}

// MARK: - Success handler

type successResp struct {
	Status string      `json:"status" example:"success"`
	Data   interface{} `json:"data"`
}

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

// Ping godoc
// @Summary      Show hello text
// @Description  very friendly response
// @Tags         Tests
// @Security ApiKeyAuth
// @Produce      json
// @Router       /ping [get]
func (h *Handler) Ping(c *gin.Context) {
	if userID, exists := c.Get("user_id"); exists {
		c.JSON(http.StatusOK, gin.H{"user_id": userID})
		return
	}
	h.errorHandler(c, http.StatusInternalServerError, errors.New("user_id not found"))
}
