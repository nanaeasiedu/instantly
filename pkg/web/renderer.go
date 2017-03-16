package web

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type Template struct{}

func NewTemplateRenderer() *Template {
	renderer := new(Template)
	return renderer
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var templateName string
	switch name {
	case "index":
		templateName = "index"
	case "login":
		templateName = "login"
	case "signup":
		templateName = "signup"
	}

	template, _ := template.New("home").ParseFiles("web/templates/layout.html", "web/templates/flash.html", "web/templates/"+templateName+".html")
	return template.ExecuteTemplate(w, "base", data)
}
