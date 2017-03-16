package api

import (
	"errors"

	"github.com/labstack/echo"
)

var ErrProvideAPIKey error = errors.New("API Key is required in the header. X-Api-Key")
var ErrInternalServer error = errors.New("Internal Server Error")
var ErrInvalidCredentials error = errors.New("Invalid API Key")

func StartAPIRouter(e *echo.Echo) {
	groupedAPI := e.Group("/api/v1/payment", AllowHeaders)
	groupedAPI.POST("", HandlePayments, RequireAPIKey)
	groupedAPI.POST("/transfer", HandlePaymentsTransfer, RequireAPIKey)
	groupedAPI.GET("/callback", HandleCallback)
}
