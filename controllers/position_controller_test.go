package controllers

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"stock/database"
	"stock/services"
	"testing"
)

func TestPositionControllerCreate(t *testing.T) {
	positionController := positionController{service: services.NewStockService(database.GetMockClient())}

	e := echo.New()
	var jsonStr = []byte(`{"stock_id":123123123123, "quantity": 1, "price": 10.0, "commission":2.5}`)
	req := httptest.NewRequest(http.MethodPost, "/positions", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := positionController.Create(c)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusCreated, rec.Code)
}

func TestPositionControllerCreateBindJsonFail(t *testing.T) {
	positionController := positionController{service: services.NewStockService(database.GetMockClient())}

	e := echo.New()
	var jsonStr = []byte(`{"stock_id":123123123123, "quantity": 1, "price": 10.0, "commission":2.5}`)
	req := httptest.NewRequest(http.MethodPost, "/positions", bytes.NewBuffer(jsonStr))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := positionController.Create(c)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
}

func TestPositionControllerCreateValidationFail(t *testing.T) {
	positionController := positionController{service: services.NewStockService(database.GetMockClient())}

	e := echo.New()
	var jsonStr = []byte(`{}`)
	req := httptest.NewRequest(http.MethodPost, "/positions", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := positionController.Create(c)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
}

func TestPositionControllerCreateServiceFail(t *testing.T) {
	positionController := positionController{service: services.NewStockService(database.GetMockClient())}

	e := echo.New()
	var jsonStr = []byte(`{"real": "madrid"}`)
	req := httptest.NewRequest(http.MethodPost, "/positions", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := positionController.Create(c)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
}
