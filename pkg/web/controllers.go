package web

import (
	"net/http"

	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/ngenerio/instantly/pkg/models"
	"github.com/ngenerio/instantly/pkg/web/payloads"
	log "github.com/sirupsen/logrus"
)

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"Title": "Instant Mobie Money Payments",
	})
}

func LoginHandler(c echo.Context) error {
	params := new(Params)
	params.Title = "Login - Instant"
	session := c.Get("session").(*sessions.Session)
	params.Flashes = session.Flashes()
	session.Save(c.Request(), c.Response())
	return c.Render(http.StatusOK, "login", params)
}

func RegisterHandler(c echo.Context) error {
	params := new(Params)
	params.Title = "Register - Instant"
	session := c.Get("session").(*sessions.Session)
	params.Flashes = session.Flashes()
	session.Save(c.Request(), c.Response().Writer())
	return c.Render(http.StatusOK, "signup", params)
}

func RegisterUser(c echo.Context) error {
	user := new(payloads.User)
	if err := c.Bind(user); err != nil {
		AddFlash(c, c.Response().Writer(), c.Request(), "error", err.Error())
		return c.Redirect(http.StatusFound, "/register")
	}

	dbUser := new(models.User)
	err := dbUser.GetUser(map[string]interface{}{"email_address": user.Email})
	if err := c.Bind(user); err != nil {
		AddFlash(c, c.Response().Writer(), c.Request(), "error", err.Error())
		return c.Redirect(http.StatusFound, "/register")
	}

	dbUser, err = models.CreateUser(user)

	if err != nil {
		AddFlash(c, c.Response().Writer(), c.Request(), "error", err.Error())
		return c.Redirect(http.StatusFound, "/register")
	}

	log.Info(fmt.Sprintf("New user has been created %v", dbUser))
	session := c.Get("session").(*sessions.Session)
	session.Values["id"] = dbUser.ID
	session.Save(c.Request(), c.Response().Writer())
	c.Redirect(http.StatusFound, "/")
	return nil
}

func LoginUser(c echo.Context) error {
	return nil
}
