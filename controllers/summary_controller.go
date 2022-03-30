package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"stock/services"
	"stock/utils/response_utils"
	"stock/viewmodels"
)

var (
	SummaryController summaryControllerInterface = &summaryController{service: services.NewSummaryService()}
)

type summaryControllerInterface interface {
	GetSummary(echo.Context) error
	GetOpenPositionSummary(echo.Context) error
	GetClosedPositionSummary(echo.Context) error
}

type summaryController struct {
	service services.SummaryServiceInterface
}

func (s summaryController) GetSummary(c echo.Context) error {
	positions, serviceErr := s.service.Summary()
	if serviceErr != nil {
		response := response_utils.NewErrorResponse(serviceErr.StatusCode, serviceErr.NotificationType, serviceErr.Message)
		return c.JSON(response.Code, response)
	}

	var positionListVm viewmodels.SummaryVM
	result := positionListVm.BindPositionsToSummaryVM(positions)

	response := response_utils.NewSuccessResponse(http.StatusOK, "Summary", result)

	return c.JSON(http.StatusOK, response)
}

func (s summaryController) GetOpenPositionSummary(c echo.Context) error {
	positions, serviceErr := s.service.OpenPositions()
	if serviceErr != nil {
		response := response_utils.NewErrorResponse(serviceErr.StatusCode, serviceErr.NotificationType, serviceErr.Message)
		return c.JSON(response.Code, response)
	}

	var positionListVm viewmodels.SummaryVM
	result := positionListVm.BindPositionsToSummaryVM(positions)

	response := response_utils.NewSuccessResponse(http.StatusOK, "Open positions summary", result)

	return c.JSON(http.StatusOK, response)
}

func (s summaryController) GetClosedPositionSummary(c echo.Context) error {
	positions, serviceErr := s.service.ClosedPositions()
	if serviceErr != nil {
		response := response_utils.NewErrorResponse(serviceErr.StatusCode, serviceErr.NotificationType, serviceErr.Message)
		return c.JSON(response.Code, response)
	}

	var positionListVm viewmodels.SummaryVM
	result := positionListVm.BindClosedPositionsToSummaryVM(positions)

	response := response_utils.NewSuccessResponse(http.StatusOK, "Closed positions summary", result)

	return c.JSON(http.StatusOK, response)
}
