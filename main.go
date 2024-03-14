package main

import (
	"fmt"
	"htmx-demo/dogs"
	"htmx-demo/views"
	"net/http"
	"sort"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var dogMap = dogs.DogMap{}
var selectedId = ""

func createDogHandler(c echo.Context) error {
  name := c.FormValue("name")
  breed := c.FormValue("breed")
  const dog = dogs.Add(dogMap, name, breed);
  return c.HTML(http.StatusCreated, views.DogRow(dog));
}

func formHandler(c echo.Context) error {
	const selectedDog = dogMap[selectedId]

	return render(c, views.Form(selectedDog))
}

// This is just for testing.
func homeHandler(c echo.Context) error {
	return render(c, views.Hello("John"))
	// return render(c, views.Home("Page Title"))
}

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}

func rowsHandler(c echo.Context) error {
    dogSlice := []dogs.Dog
    for _, dog := range dogMap {
        dogSlice = append(dogSlice, dog)
    }
	less := func(i, j int) bool {
		return dogSlice[i].name < dogSlice[j].name
	}
    sort.Slice(dogSlice, less)
	fmt.Println("dogSlice = %v\n", dogSlice)
	return render(c, views.DogRows(dogSlice))
}

func selectHandler(c echo.Context) error {
	selectedId := c.Param("id")
	c.Response().Header().Set("HX-Trigger", "selection-change")
	return c.NoContent(http.StatusOK);
}

func updateDogHandler(c echo.Context) error {
  id := c.Param("id")
  name := c.FormValue("name")
  breed := c.FormValue("breed")
  const updatedDog = dogs.Dog {
	id: id,
	name: name,
	breed: breed,
  }
  dogMap[id] = updatedDog

  selectedId = "";
  c.Response().Header().Set("HX-Trigger", "selection-change")
  return c.HTML(http.StatusCreated, views.DogRow(updatedDog));
}

func main() {
	// sum := utils.Add(1, 2);
    // fmt.Println("sum =", sum)

	dogs.Add(dogMap, "Comet", "Whippet")
	dogs.Add(dogMap, "Oscar", "German Shorthaired Pointer")
    fmt.Println("dogMap =", dogMap)

	e := echo.New()
	e.Use(middleware.Logger())
    e.Static("/", "public")
    e.File("/", "public/index.html")

	// e.GET("/", homeHandler)
	e.GET("/form", formHandler)
	e.GET("/rows", rowsHandler)
	e.GET("/select/:id", selectHandler)
	e.POST("/dog", createDogHandler)
	e.PUT("/dog", updateDogHandler)

	e.Logger.Fatal(e.Start(":3000"))
}