package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(echo.Context) error
}

type pingController struct {
}

func (p pingController) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
