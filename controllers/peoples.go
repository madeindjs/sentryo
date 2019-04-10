package controllers

import (
	"../models"
	"github.com/labstack/echo"
	"net/http"
)

func PeoplesIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, models.AllPeoples())
}

func PeoplesShow(c echo.Context) error {
	idStr := c.Param("id")

	vehicle, error := models.FindPeople(idStr)

	if error != nil {
		return c.String(http.StatusNotFound, "People does not exist")
	} else {
		return c.JSON(http.StatusOK, vehicle)
	}
}
