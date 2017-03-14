package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ngenerio/instantly/pkg/api"
	"github.com/ngenerio/instantly/pkg/config"
	"github.com/ngenerio/instantly/pkg/models"
	"github.com/ngenerio/instantly/pkg/web"
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

	web.StartWebRouter(echoInstance)
	api.StartAPIRouter(echoInstance)

	echoInstance.Start(":3000")
}
