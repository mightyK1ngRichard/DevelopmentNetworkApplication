//go:build !appengine && !appenginevm

// Package main is the main package.
package main

import (
	"VikingsServer/internal/app/config"
	"VikingsServer/internal/app/dsn"
	"VikingsServer/internal/app/handler"
	"VikingsServer/internal/app/kingMinio"
	app "VikingsServer/internal/app/pkg"
	"VikingsServer/internal/app/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @title VIKINGS
// @version 1.0
// @description Viking's hikes
// @contact.name API Support
// @contact.url https://github.com/mightyK1ngRichard
// @contact.email dimapermyakov55@gmai.com

// @host localhost:7070
// @schemes http
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// ShowAccount godoc
// @Summary      Cities
// @Description  Get cities list
// @Tags         cities
// @Produce      json
// @Router       /api/v3/cities [get]
func main() {
	logger := logrus.New()
	minioClient := kingMinio.NewMinioClient(logger)
	router := gin.Default()
	conf, err := config.NewConfig(logger)
	if err != nil {
		logger.Fatalf("Error with configuration reading: %s", err)
	}

	postgresString, errPost := dsn.FromEnv()
	if errPost != nil {
		logger.Fatalf("Error of reading postgres line: %s", errPost)
	}
	fmt.Println(postgresString)
	rep, errRep := repository.NewRepository(postgresString, logger)
	if errRep != nil {
		logger.Fatalf("Error from repository: %s", err)
	}
	hand := handler.NewHandler(logger, rep, minioClient, conf)
	application := app.NewApp(conf, router, logger, hand)
	application.RunApp()
}
