package web

import (
	"html/template"
	"io"

	"fmt"

	"github.com/labstack/echo"
)

type Template struct{}

func NewTemplateRenderer() *Template {
	renderer := new(Template)
	return renderer
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	template, err := template.New("home").ParseFiles("web/templates/layout.html", "web/templates/flash.html", "web/templates/sidebar.html", "web/templates/nav.html", "web/templates/"+name+".html")
	fmt.Println(err)
	return template.ExecuteTemplate(w, "base", data)
}
