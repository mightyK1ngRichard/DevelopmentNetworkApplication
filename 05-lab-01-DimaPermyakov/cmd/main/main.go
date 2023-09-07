package main

import (
	app "05-lab-01-DimaPermyakov/internal/app/pkg"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	router := gin.Default()

	application := app.NewApp(router, logger)
	application.RunApp()
}
