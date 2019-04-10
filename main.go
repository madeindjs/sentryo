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
	// vehicles
	e.GET("/vehicles/", controllers.VehiclesIndex)
	e.GET("/vehicles/:id", controllers.VehiclesShow)
	// starships
	e.GET("/starships/", controllers.StarshipsIndex)
	e.GET("/starships/:id", controllers.StarshipsShow)
	// peoples
	e.GET("/peoples/", controllers.PeoplesIndex)
	e.GET("/peoples/:id", controllers.PeoplesShow)
	e.POST("/peoples/", controllers.PeoplesCreate)
	e.PUT("/peoples/:id", controllers.PeoplesUpdate)

	e.Logger.Fatal(e.Start(":1323"))
}
