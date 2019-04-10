package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	// "fmt"
	"./controllers"
)

const port = ":8000"

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.GET("/", controllers.Home)
	e.GET("/vehicles/", controllers.VehiclesIndex)
	e.GET("/vehicles/:id", controllers.VehiclesShow)
	e.Logger.Fatal(e.Start(":1323"))
}
