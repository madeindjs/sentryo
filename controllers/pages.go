package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Home(c echo.Context) error {
	return c.File("README.md")
	return c.String(http.StatusOK, "Hello Sentryo")
}
