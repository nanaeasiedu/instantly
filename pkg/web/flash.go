package web

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

type Flash struct {
	Type    string
	Message string
}

func AddFlash(c echo.Context, w http.ResponseWriter, r *http.Request, flashType, message string) {
	session := c.Get("session").(*sessions.Session)
	session.AddFlash(Flash{flashType, message})
	session.Save(r, w)
}
