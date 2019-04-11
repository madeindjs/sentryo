package controllers

import (
	"../models"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func PeoplesIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, models.AllPeoples())
}

func PeoplesShow(c echo.Context) error {
	idStr := c.Param("id")
	people, error := models.FindPeople(idStr)

	if error != nil {
		return c.String(http.StatusNotFound, "People does not exist")
	} else {
		return c.JSON(http.StatusOK, people)
	}
}

func PeoplesCreate(c echo.Context) error {
	people := &models.People{}

	if err := c.Bind(people); err != nil {
		log.Fatal(err)
		return err
	}

	if _, err := people.Insert(); err != nil {
		log.Fatal(err)
		return err
	}

	return c.JSON(http.StatusCreated, people)
}

func PeoplesUpdate(c echo.Context) error {
	people := &models.People{}
	people.Id = c.Param("id")

	if err := c.Bind(people); err != nil {
		return err
	}

	if _, err := people.Save(); err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, people)
}

func PeoplesDelete(c echo.Context) error {
	idStr := c.Param("id")
	people, error := models.FindPeople(idStr)

	if error != nil {
		log.Fatal(error)
		return c.String(http.StatusNotFound, "People does not exist")
	}

	if _, err := people.Delete(); err != nil {
		log.Fatal(err)
		return err
	}

	return c.String(http.StatusNoContent, "People removed")
}
