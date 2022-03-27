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
	DividendController dividendControllerInterface = &dividendController{service: services.NewDividendService()}
)

type dividendControllerInterface interface {
	Create(echo.Context) error
}

type dividendController struct {
	service services.DividendServiceInterface
}

func (d dividendController) Create(c echo.Context) error {
	dividend := models.Dividend{}
	if err := c.Bind(&dividend); err != nil {
		restErr := error_utils.NewBadRequestError("bind error when trying to bind dividend", error_utils.JsonBindError)
		response := response_utils.NewErrorResponse(restErr.StatusCode, restErr.NotificationType, restErr.Message)
		return c.JSON(response.Code, response)
	}

	if err := dividend.Validate(); err != nil {
		response := response_utils.NewErrorResponse(err.StatusCode, err.NotificationType, err.Message)
		return c.JSON(response.Code, response)
	}

	result, serviceErr := d.service.Create(dividend)
	if serviceErr != nil {
		response := response_utils.NewErrorResponse(serviceErr.StatusCode, serviceErr.NotificationType, serviceErr.Message)
		return c.JSON(response.Code, response)
	}

	response := response_utils.NewSuccessResponse(http.StatusCreated, "Dividend created", result)

	return c.JSON(http.StatusCreated, response)
}
