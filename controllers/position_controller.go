package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"stock/database"
	"stock/models"
	"stock/services"
	"stock/utils/error_utils"
	"stock/utils/response_utils"
)

var (
	PositionController positionControllerInterface = &positionController{service: services.NewStockService(database.GetClient())}
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
		restErr := error_utils.NewBadRequestError("bind error when trying to bind position", 1)
		response := response_utils.NewErrorResponse(restErr.StatusCode, restErr.NotificationType, restErr.Message)
		return c.JSON(response.Code, response)
	}

	if err := position.Validate(); err != nil {
		response := response_utils.NewErrorResponse(err.StatusCode, err.NotificationType, err.Message)
		return c.JSON(response.Code, response)
	}

	position, serviceErr := p.service.Create(*position)
	if serviceErr != nil {
		response := response_utils.NewErrorResponse(serviceErr.StatusCode, serviceErr.NotificationType, serviceErr.Message)
		return c.JSON(response.Code, response)
	}

	response := response_utils.NewSuccessResponse(http.StatusCreated, "Position created", position)

	return c.JSON(http.StatusCreated, response)
}
