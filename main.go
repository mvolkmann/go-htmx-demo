package main

import (
	"fmt"
	"htmx-demo/model"
	"htmx-demo/views"
	"net/http"
	"sort"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var dogMap = model.DogMap{}
var selectedId = ""

func createDogHandler(c echo.Context) error {
  name := c.FormValue("name")
  breed := c.FormValue("breed")
  dog := model.Add(dogMap, name, breed);
  return render(c, views.DogRow(dog, false));
  // templ.WithStatus(http.StatusCreated)
}

func formHandler(c echo.Context) error {
	selectedDog := dogMap[selectedId]
	return render(c, views.Form(&selectedDog))
}

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}

func rowsHandler(c echo.Context) error {
    dogSlice := []model.Dog{}
    for _, dog := range dogMap {
        dogSlice = append(dogSlice, dog)
    }
	less := func(i, j int) bool {
		return dogSlice[i].Name < dogSlice[j].Name
	}
    sort.Slice(dogSlice, less)
	fmt.Println("dogSlice = %v\n", dogSlice)
	return render(c, views.DogRows(dogSlice))
}

func selectHandler(c echo.Context) error {
	selectedId = c.Param("id")
	c.Response().Header().Set("HX-Trigger", "selection-change")
	return c.NoContent(http.StatusOK);
}

func updateDogHandler(c echo.Context) error {
  id := c.Param("id")
  name := c.FormValue("name")
  breed := c.FormValue("breed")
  updatedDog := model.Dog{
	Id: id,
	Name: name,
	Breed: breed,
  }
  dogMap[id] = updatedDog

  selectedId = "";
  c.Response().Header().Set("HX-Trigger", "selection-change")
  return render(c, views.DogRow(updatedDog, true));
}

func main() {
	// sum := utils.Add(1, 2);
    // fmt.Println("sum =", sum)

	model.AddDog(dogMap, "Comet", "Whippet")
	model.AddDog(dogMap, "Oscar", "German Shorthaired Pointer")
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