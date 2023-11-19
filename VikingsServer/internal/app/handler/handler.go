package handler

import (
	_ "VikingsServer/docs"
	"VikingsServer/internal/app/config"
	"VikingsServer/internal/app/ds"
	"VikingsServer/internal/app/redis"
	"VikingsServer/internal/app/repository"
	"VikingsServer/internal/app/role"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

const (
	baseURL                      = "api/v3"
	citiesHTML                   = "cities"
	cities                       = baseURL + "/cities"
	addCityIntoHike              = baseURL + "/cities/add-city-into-hike"
	addCityImage                 = baseURL + "/cities/upload-image"
	hikes                        = baseURL + "/hikes"
	hikesUpdateStatus            = baseURL + "/hikes-update-status"
	hikeUpdateStatusForModerator = baseURL + "/hikes/update/status-for-moderator"
	hikeUpdateStatusForUser      = baseURL + "/hikes/update/status-for-user"
	users                        = baseURL + "/users"
	login                        = users + "/login"
	signup                       = users + "/sign_up"
	logout                       = users + "/logout"
	destinationHikes             = baseURL + "/destination-hikes"
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
	h.CityCRUD(router)
	h.HikeCRUD(router)
	h.DestinationHikesCRUD(router)
	h.UserCRUD(router)
	registerStatic(router)
}

func (h *Handler) CityCRUD(router *gin.Engine) {
	router.GET(cities, h.CitiesList)
	router.POST(cities, h.AddCity)
	router.POST(addCityImage, h.AddImage)
	router.PUT(cities, h.UpdateCity)
	router.DELETE(cities, h.DeleteCity)
	router.POST(addCityIntoHike, h.AddCityIntoHike)
	router.Use(cors.Default()).DELETE("/api/v3/cities/delete/:id", h.DeleteCityWithParam)
}

func (h *Handler) UserCRUD(router *gin.Engine) {
	//router.GET(users, h.UsersList)
	router.POST(login, h.Login)
	router.POST(signup, h.Register)
	router.GET(logout, h.Logout)

	// TODO: Delete this endpoint from lab05
	router.Use(h.WithAuthCheck(role.Manager, role.Admin)).GET("/ping", h.Ping)
}

func (h *Handler) DestinationHikesCRUD(router *gin.Engine) {
	router.GET(destinationHikes, h.DestinationHikesList)
	router.POST(destinationHikes, h.AddDestinationToHike)
	router.PUT(destinationHikes, h.UpdateDestinationHikeNumber)
	router.DELETE(destinationHikes, h.DeleteDestinationToHike)
}

func (h *Handler) HikeCRUD(router *gin.Engine) {
	//router.POST(hikes, h.AddHike)
	router.GET(hikes, h.HikesList)
	router.DELETE(hikes, h.DeleteHike)
	//router.PUT(hikesUpdateStatus, h.UpdateHikeStatus)
	router.PUT(hikeUpdateStatusForModerator, h.UpdateStatusForModerator)
	router.PUT(hikeUpdateStatusForUser, h.UpdateStatusForUser)
	router.PUT(hikes, h.UpdateHike)
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

// Ping godoc
// @Summary      Show hello text
// @Description  very friendly response
// @Tags         Tests
// @Produce      json
// @Success      200  {object}  pingResp
// @Router       /ping [get]
func (h *Handler) Ping(c *gin.Context) {
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		h.errorHandler(c, http.StatusUnauthorized, errors.New("cookie is empty"))
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.Config.JWT.Token), nil
	})

	if err != nil {
		h.errorHandler(c, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	if claims, ok := token.Claims.(*ds.JWTClaims); ok && token.Valid {
		userID := claims.UserID
		c.JSON(http.StatusOK, gin.H{"user_id": userID})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}
