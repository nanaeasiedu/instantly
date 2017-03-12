package web

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type Template struct {
	Home   *template.Template
	Login  *template.Template
	Signup *template.Template
}

func NewTemplateRenderer() *Template {
	renderer := new(Template)
	renderer.Home, _ = template.New("home").ParseFiles("web/templates/layout.html", "web/templates/index.html")
	renderer.Login, _ = template.New("login").ParseFiles("web/templates/layout.html", "web/templates/login.html")
	renderer.Signup, _ = template.New("signup").ParseFiles("web/templates/layout.html", "web/templates/signup.html")

	return renderer
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var template *template.Template

	switch name {
	case "index":
		template = t.Home
	case "login":
		template = t.Login
	case "signup":
		template = t.Signup
	}

	return template.ExecuteTemplate(w, "base", data)
}
