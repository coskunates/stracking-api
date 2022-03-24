package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"stock/utils/response_utils"
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
	response := response_utils.NewSuccessResponseWithEmptyData(http.StatusOK, "pong")

	return c.JSON(http.StatusOK, response)
}
