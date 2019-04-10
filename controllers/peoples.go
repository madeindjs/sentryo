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

func PeoplesCreate(c echo.Context) error {
	people := &models.People{}

	if err := c.Bind(people); err != nil {
		return err
	}

	if _, err := people.Insert(); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, people)
}

func PeoplesUpdate(c echo.Context) error {
	people := &models.People{}

	if err := c.Bind(people); err != nil {
		return err
	}

	if _, err := people.Save(); err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, people)
}
