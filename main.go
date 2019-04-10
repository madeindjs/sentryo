package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"strconv"

	"fmt"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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
	database, error := sql.Open("sqlite3", "swapi.dat")

	if error != nil {
		panic("Could not open database")
	}

	statement, error := database.Prepare("CREATE TABLE IF NOT EXISTS vehicles (name TEXT,model TEXT,manufacturer TEXT,cost_in_credits TEXT,length TEXT,max_atmosphering_speed TEXT,crew TEXT,passengers TEXT,cargo_capacity TEXT,consumables TEXT,vehicle_class TEXT,pilots TEXT,films TEXT,created TEXT,edited TEXT,url TEXT,id TEXT);")

	if error != nil {
		panic(error)
	}

	statement.Exec()

	rows, _ := database.Query("SELECT name, model FROM vehicles")
	var name string
	var model string
	for rows.Next() {
		rows.Scan(&name, &model)
		fmt.Println(": " + name + " " + model)
	}

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.GET("/", home)
	e.GET("/vehicules/", vehiculesIndex)
	e.GET("/vehicules/:id", vehiculesShow)
	e.Logger.Fatal(e.Start(":1323"))
}
