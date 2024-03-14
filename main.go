package main

import (
	"fmt"
	"htmx-demo/utils"
)

/*
func homeHandler(c echo.Context) error {
    fmt.Println("in homeHandler")
	// return render(c, views.hello("John"))
	return render(c, views.home("Page Title"))
}

func render(ctx echo.Context, cmp templ.Component) error {
    fmt.Println("in rener")
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}
*/

func main() {
	sum := utils.Add(1, 2);
    fmt.Println("sum =", sum)

	/*
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", homeHandler)

	e.Logger.Fatal(e.Start(":3000"))
    */
}