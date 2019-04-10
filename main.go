package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"strconv"
)

type Vehicule struct {
	Id    int    `json:"id"`
	Title string `json:"Title"`
}

type Vehicules []Vehicule

const port = ":8000"

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Sentryo")
}

func vehiculesIndex(c echo.Context) error {
	articles := Vehicules{
		Vehicule{Id: 1, Title: "Hello"},
		Vehicule{Id: 2, Title: "Hello 2"},
	}
	log.Print("GET /vehicules")

	return c.JSON(http.StatusOK, articles)
}

func vehiculesShow(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusNotFound, "Vehicule does not exist")
	}

	article := Vehicule{
		Id:    id,
		Title: "Hello",
	}

	return c.JSON(http.StatusOK, article)
}

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.GET("/", home)
	e.GET("/vehicules/", vehiculesIndex)
	e.GET("/vehicules/:id", vehiculesShow)
	e.Logger.Fatal(e.Start(":1323"))
}
