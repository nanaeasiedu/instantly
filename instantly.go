package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ngenerio/instantly/pkg/api"
	"github.com/ngenerio/instantly/pkg/config"
	"github.com/ngenerio/instantly/pkg/models"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := models.Setup()

	if err != nil {
		log.Error(err)
	}

	echoInstance := echo.New()

	echoInstance.Use(middleware.Recover())
	echoInstance.Use(middleware.Logger())
	echoInstance.Use(middleware.Gzip())
	echoInstance.Use(middleware.CORS())

	if config.Settings.Env == "development" {
		echoInstance.Debug = true
	}

	api.StartAPIRouter(echoInstance)

	echoInstance.Start(":5000")
}
