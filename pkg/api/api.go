package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ngenerio/instantly/pkg/config"
)

func StartAPIRouter(e *echo.Echo) {
	authMiddleware := middleware.BasicAuth(func(username, password string) bool {
		if username == config.Settings.APIClientUsername && password == config.Settings.APIClientPassword {
			return true
		}

		return false
	})

	groupedAPI := e.Group("/api/v1/payment")
	groupedAPI.POST("", HandlePayments, authMiddleware)
	groupedAPI.POST("/transfer", HandlePaymentsTransfer, authMiddleware)
	groupedAPI.GET("/callback", HandleCallback)
}
