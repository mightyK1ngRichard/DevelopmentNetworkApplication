package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Application struct {
	//Config     *config.Config
	Logger *logrus.Logger
	Router *gin.Engine
	//Repository *repository.Repository
}

func NewApp(r *gin.Engine, l *logrus.Logger) *Application {
	return &Application{
		Logger: l,
		Router: r,
	}
}

func (a *Application) RunApp() {
	a.Logger.Info("Start Application")

	a.Logger.Info("Terminate Application")
}
