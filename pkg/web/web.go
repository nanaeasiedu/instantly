package web

import (
	"encoding/gob"

	"fmt"

	"github.com/boj/redistore"
	"github.com/labstack/echo"
	"github.com/ngenerio/instantly/pkg/config"
	"github.com/ngenerio/instantly/pkg/models"
	log "github.com/sirupsen/logrus"
)

var Store *redistore.RediStore

func init() {
	var err error
	gob.Register(&models.User{})
	gob.Register(&Flash{})
	Store, err = redistore.NewRediStore(10, "tcp", config.Settings.RedisURL, config.Settings.RedisPassword, []byte("instantly_web_sid"))

	if err != nil {
		log.Info(fmt.Sprintf("Error connecting to redis instance %v", err))
		panic(err)
	}

	Store.SetMaxAge(86400 * 7)
}

func StartWebRouter(e *echo.Echo) {
	web := e.Group("/", SessionMiddleware)
	e.Renderer = NewTemplateRenderer()
	e.Static("/static", "web/static")

	web.GET("", HomeHandler, RequireLogin)
	web.GET("login", LoginHandler)
	web.POST("login", LoginUser)
	web.GET("register", RegisterHandler)
	web.POST("register", RegisterUser)
	web.GET("logout", Logout)
}
