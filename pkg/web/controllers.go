package web

import (
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/ngenerio/instantly/pkg/models"
	"github.com/ngenerio/instantly/pkg/web/payloads"
	log "github.com/sirupsen/logrus"
)

var ErrEmailExists error = errors.New("Email already exists")
var ErrInvalidCredentials error = errors.New("Invalid email or password")
var ErrInternal error = errors.New("Something happened. Please try again")

func HomeHandler(c echo.Context) error {
	params := new(Params)
	params.Title = "Home - Instant"
	return c.Render(http.StatusOK, "index", params)
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
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrInternal.Error())
		return c.Redirect(http.StatusFound, "/register")
	}

	exists, err := models.DoesUserExist(map[string]interface{}{"email_address": user.Email})
	if err != nil {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrInternal.Error())
		return c.Redirect(http.StatusFound, "/register")
	}

	if exists {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrEmailExists.Error())
		return c.Redirect(http.StatusFound, "/register")
	}

	dbUser, err := models.CreateUser(user)
	if err != nil {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrInternal.Error())
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
	user := new(payloads.User)
	if err := c.Bind(user); err != nil {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrInternal.Error())
		return c.Redirect(http.StatusFound, "/login")
	}

	dbUser := new(models.User)
	if err := dbUser.GetUser(map[string]interface{}{"email_address": user.Email}); err != nil {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrInvalidCredentials.Error())
		return c.Redirect(http.StatusFound, "/login")
	}

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(user.Password))

	if err != nil {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrInvalidCredentials.Error())
		return c.Redirect(http.StatusFound, "/login")
	}

	log.Info(fmt.Sprintf("User with email address %s has logged in", user.Email))
	session := c.Get("session").(*sessions.Session)
	session.Values["id"] = dbUser.ID
	session.Save(c.Request(), c.Response().Writer())
	c.Redirect(http.StatusFound, "/")
	return nil
}

func Logout(c echo.Context) error {
	session := c.Get("session").(*sessions.Session)
	delete(session.Values, "id")
	SetFlash(c, c.Response().Writer(), c.Request(), "success", "You have successfully logged out")
	c.Redirect(http.StatusFound, "/")
	return nil
}
