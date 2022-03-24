package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"stock/database"
	"stock/services"
	"stock/utils/response_utils"
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

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	assert.EqualValues(t, http.StatusCreated, rec.Code)
	assert.False(t, response.Error)
	assert.Equal(t, "success", response.Type)
	assert.EqualValues(t, "Position created", response.Message)
	assert.NotNil(t, response.Data)
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

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.True(t, response.Error)
	assert.Equal(t, "warning", response.Type)
	assert.EqualValues(t, "bind error when trying to bind position", response.Message)
	assert.Nil(t, response.Data)
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

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.True(t, response.Error)
	assert.Equal(t, "warning", response.Type)
	assert.EqualValues(t, "stock id is required", response.Message)
	assert.Nil(t, response.Data)
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
	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.True(t, response.Error)
	assert.Equal(t, "warning", response.Type)
	assert.EqualValues(t, "stock id is required", response.Message)
	assert.Nil(t, response.Data)
}
