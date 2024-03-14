package main

import (
	"fmt"
	"htmx-demo/utils"
	"htmx-demo/views"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func homeHandler(c echo.Context) error {
	return render(c, views.Hello("John"))
	// return render(c, views.Home("Page Title"))
}

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}

func main() {
	sum := utils.Add(1, 2);
    fmt.Println("sum =", sum)

	app := echo.New()
	app.Use(middleware.Logger())

	app.GET("/", homeHandler)

	app.Logger.Fatal(app.Start(":3000"))
}