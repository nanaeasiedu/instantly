package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ngenerio/instantly/pkg/models"
)

func RequireAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.Request().Header.Get("X-Api-Key")

		if key == "" {
			return c.JSON(http.StatusUnauthorized, Response{Status: "error", Message: ErrProvideAPIKey.Error()})
		}

		exists, err := models.DoesUserExist(map[string]interface{}{"token": key})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, Response{Status: "error", Message: ErrInternalServer.Error()})
		}

		if !exists {
			return c.JSON(http.StatusUnauthorized, Response{Status: "error", Message: ErrInvalidCredentials.Error()})
		}

		user := new(models.User)
		user.GetUser(map[string]interface{}{"token": key})
		c.Set("user", user)
		return next(c)
	}
}

func AllowHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Headers", "X-Api-Key, Content-Type")
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}
