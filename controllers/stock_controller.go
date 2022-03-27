package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"stock/models"
	"stock/services"
	"stock/utils/error_utils"
	"stock/utils/response_utils"
)

var (
	StockController stockControllerInterface = &stockController{service: services.NewStockService()}
)

type stockControllerInterface interface {
	List(echo.Context) error
	Create(echo.Context) error
}

type stockController struct {
	service services.StockServiceInterface
}

func (s stockController) List(c echo.Context) error {
	result, serviceErr := s.service.List()
	if serviceErr != nil {
		response := response_utils.NewErrorResponse(serviceErr.StatusCode, serviceErr.NotificationType, serviceErr.Message)
		return c.JSON(response.Code, response)
	}

	response := response_utils.NewSuccessResponse(http.StatusCreated, "Stocks listed", result)

	return c.JSON(http.StatusOK, response)
}

func (s stockController) Create(c echo.Context) error {
	stock := models.Stock{}
	if err := c.Bind(&stock); err != nil {
		restErr := error_utils.NewBadRequestError("bind error when trying to bind stock", error_utils.JsonBindError)
		response := response_utils.NewErrorResponse(restErr.StatusCode, restErr.NotificationType, restErr.Message)
		return c.JSON(response.Code, response)
	}

	if err := stock.Validate(); err != nil {
		response := response_utils.NewErrorResponse(err.StatusCode, err.NotificationType, err.Message)
		return c.JSON(response.Code, response)
	}

	result, serviceErr := s.service.Create(stock)
	if serviceErr != nil {
		response := response_utils.NewErrorResponse(serviceErr.StatusCode, serviceErr.NotificationType, serviceErr.Message)
		return c.JSON(response.Code, response)
	}

	response := response_utils.NewSuccessResponse(http.StatusCreated, "Stock created", result)

	return c.JSON(http.StatusCreated, response)
}
