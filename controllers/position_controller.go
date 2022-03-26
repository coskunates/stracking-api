package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"stock/models"
	"stock/services"
	"stock/utils/error_utils"
	"stock/utils/logger_utils"
	"stock/utils/response_utils"
	"strconv"
)

var (
	PositionController positionControllerInterface = &positionController{service: services.NewPositionService()}
)

type positionControllerInterface interface {
	Create(c echo.Context) error
	Close(c echo.Context) error
}

type positionController struct {
	service services.PositionServiceInterface
}

func (p positionController) Create(c echo.Context) error {
	position := models.Position{}
	if err := c.Bind(&position); err != nil {
		restErr := error_utils.NewBadRequestError("bind error when trying to bind position", error_utils.JsonBindError)
		response := response_utils.NewErrorResponse(restErr.StatusCode, restErr.NotificationType, restErr.Message)
		return c.JSON(response.Code, response)
	}

	if err := position.Validate(); err != nil {
		response := response_utils.NewErrorResponse(err.StatusCode, err.NotificationType, err.Message)
		return c.JSON(response.Code, response)
	}

	result, serviceErr := p.service.Create(position)
	if serviceErr != nil {
		response := response_utils.NewErrorResponse(serviceErr.StatusCode, serviceErr.NotificationType, serviceErr.Message)
		return c.JSON(response.Code, response)
	}

	response := response_utils.NewSuccessResponse(http.StatusCreated, "Position created", result)

	return c.JSON(http.StatusCreated, response)
}

func (p positionController) Close(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("position_id"), 10, 64)
	if err != nil {
		logger_utils.Error(err.Error(), err)
		restErr := error_utils.NewBadRequestError("invalid position id", error_utils.InvalidRequestParams)
		response := response_utils.NewErrorResponse(restErr.StatusCode, restErr.NotificationType, restErr.Message)
		return c.JSON(response.Code, response)
	}

	position, restErr := p.service.Get(models.Position{ID: id})
	if restErr != nil {
		response := response_utils.NewErrorResponse(restErr.StatusCode, restErr.NotificationType, restErr.Message)
		return c.JSON(response.Code, response)
	}

	closedPosition := models.ClosedPosition{
		PositionId: position.ID,
		StockId:    position.StockId,
		Quantity:   position.Quantity,
		Price:      position.Price,
		Commission: position.Commission,
		OpenedAt:   position.OpenedAt,
	}

	if err := c.Bind(&closedPosition); err != nil {
		restErr := error_utils.NewBadRequestError("bind error when trying to bind closed position", error_utils.JsonBindError)
		response := response_utils.NewErrorResponse(restErr.StatusCode, restErr.NotificationType, restErr.Message)
		return c.JSON(response.Code, response)
	}

	if err := closedPosition.Validate(); err != nil {
		response := response_utils.NewErrorResponse(err.StatusCode, err.NotificationType, err.Message)
		return c.JSON(response.Code, response)
	}

	result, serviceErr := p.service.ClosePosition(closedPosition)
	if serviceErr != nil {
		response := response_utils.NewErrorResponse(serviceErr.StatusCode, serviceErr.NotificationType, serviceErr.Message)
		return c.JSON(response.Code, response)
	}

	response := response_utils.NewSuccessResponse(http.StatusOK, "Position closed successfully", result)

	return c.JSON(http.StatusOK, response)
}
