package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
	component := hello("John")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
*/

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// component := hello("John")
	// http.Handle("/", templ.Handler(component))

	/*
	e.GET("/", func(c echo.Context) error {
		// return c.String(http.StatusOK, templ.Handler(component))
		return c.String(http.StatusOK, "Hello from Echo!")
	})
	*/

	e.Renderer = newTemplates()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.Logger.Fatal(e.Start(":3000"))
}