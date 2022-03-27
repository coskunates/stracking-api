package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"stock/services"
	"stock/utils/response_utils"
)

var (
	StockController stockControllerInterface = &stockController{service: services.NewStockService()}
)

type stockControllerInterface interface {
	List(c echo.Context) error
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
