package controllers

import (
	"github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/madeindjs/sentryo/models"
	"net/http"
)

func VehiclesIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, models.AllVehicles())
}

func VehiclesShow(c echo.Context) error {
	idStr := c.Param("id")

	vehicle, error := models.FindVehicle(idStr)

	if error != nil {
		glog.Error(error)
		return c.String(http.StatusNotFound, "Vehicle does not exist")
	} else {
		return c.JSON(http.StatusOK, vehicle)
	}
}
