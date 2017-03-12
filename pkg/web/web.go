package web

import (
	"net/http"

	"github.com/labstack/echo"
)

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"Title": "Instant Mobie Money Payments",
	})
}

func LoginHandler(e echo.Context) error {
	return nil
}
