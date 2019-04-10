package controllers

import (
	"../models"
	"github.com/labstack/echo"
	"net/http"
)

func StarshipsIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, models.AllStarships())
}

func StarshipsShow(c echo.Context) error {
	idStr := c.Param("id")

	vehicle, error := models.FindStarship(idStr)

	if error != nil {
		return c.String(http.StatusNotFound, "Starships does not exist")
	} else {
		return c.JSON(http.StatusOK, vehicle)
	}
}
