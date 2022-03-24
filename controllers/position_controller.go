package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"stock/database"
	"stock/models"
	"stock/services"
)

var (
	PositionController positionControllerInterface = &positionController{service: services.NewStockService(database.GetMockClient())}
)

type positionControllerInterface interface {
	Create(c echo.Context) error
}

type positionController struct {
	service services.StockServiceInterface
}

func (p positionController) Create(c echo.Context) error {
	position := new(models.Position)
	if err := c.Bind(position); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := position.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var serviceErr error
	position, serviceErr = p.service.Create(*position)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}

	return c.JSON(http.StatusCreated, position)
}
