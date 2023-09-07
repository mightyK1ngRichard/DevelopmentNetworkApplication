package main

import (
	"05-lab-01-DimaPermyakov/config"
	"05-lab-01-DimaPermyakov/internal/app/handler"
	app "05-lab-01-DimaPermyakov/internal/app/pkg"
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
	hand := handler.NewHandler()
	application := app.NewApp(conf, router, logger, hand)
	application.RunApp()
}
