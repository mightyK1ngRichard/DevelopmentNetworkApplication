package main

import (
	"VikingsServer/internal/app/config"
	"VikingsServer/internal/app/dsn"
	"VikingsServer/internal/app/handler"
	app "VikingsServer/internal/app/pkg"
	"VikingsServer/internal/app/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
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
	hand := handler.NewHandler(logger, rep)
	application := app.NewApp(conf, router, logger, hand)
	application.RunApp()
}
