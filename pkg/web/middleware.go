package web

import (
	"net/http"

	"fmt"

	"github.com/labstack/echo"
	"github.com/ngenerio/instantly/pkg/models"
	log "github.com/sirupsen/logrus"
)

func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := Store.Get(c.Request(), "_sid")
		log.Info(fmt.Sprintf("%v", session))
		// session.Save(c.Request(), c.Response().Writer())
		c.Set("session", session)

		if id, ok := session.Values["id"]; ok {
			user := &models.User{}
			err := user.GetUser(map[string]interface{}{"id": id.(int)})

			if err != nil {
				return err
			}

			c.Set("user", user)
		}

		return next(c)
	}
}

func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if u := c.Get("user"); u != nil {
			return next(c)
		}

		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}
