package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	// "fmt"
	"./models"
)

const port = ":8000"

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Sentryo")
}

func vehiclesIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, models.AllVehicles())
}

func vehiclesShow(c echo.Context) error {
	idStr := c.Param("id")

	vehicle, error := models.FindVehicle(idStr)

	if error != nil {
		return c.String(http.StatusNotFound, "Vehicle does not exist")
	} else {
		return c.JSON(http.StatusOK, vehicle)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.GET("/", home)
	e.GET("/vehicles/", vehiclesIndex)
	e.GET("/vehicles/:id", vehiclesShow)
	e.Logger.Fatal(e.Start(":1323"))
}
