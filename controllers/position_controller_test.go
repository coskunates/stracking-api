package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"stock/database"
	"stock/models"
	"stock/services"
	"stock/utils/response_utils"
	"testing"
)

type PositionTestSuite struct {
	suite.Suite
	positionId int
}

// this function executes before the test suite begins execution
func (s *PositionTestSuite) SetupSuite() {
	positions := []models.Position{
		{
			ID:         1,
			StockId:    1,
			Quantity:   10,
			Price:      5.0,
			Commission: 0.5,
			OpenedAt:   "2022-03-25",
			CreatedAt:  "2022-03-25 14:00:00",
			UpdatedAt:  "2022-03-25 14:00:00",
		},
	}

	database.GetClient().Create(&positions)
}

// this function executes after all tests executed
func (s *PositionTestSuite) TearDownSuite() {
	database.GetClient().Exec("TRUNCATE positions")
	database.GetClient().Exec("TRUNCATE closed_positions")
}

func (s PositionTestSuite) TestPositionControllerCreate() {
	positionController := positionController{service: services.NewPositionService()}

	e := echo.New()
	var jsonStr = []byte(`{"stock_id":1, "quantity": 20, "price": 10.0, "commission":2.5, "opened_at": "2002-03-25"}`)
	req := httptest.NewRequest(http.MethodPost, "/positions", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := positionController.Create(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusCreated, rec.Code)
	s.False(response.Error)
	s.Equal("success", response.Type)
	s.EqualValues("Position created", response.Message)
	s.NotNil(response.Data)

}

func (s PositionTestSuite) TestPositionControllerCreateBindJsonFail() {
	positionController := positionController{service: services.NewPositionService()}

	e := echo.New()
	var jsonStr = []byte(`{"stock_id":123123123123, "quantity": 1, "price": 10.0, "commission":2.5}`)
	req := httptest.NewRequest(http.MethodPost, "/positions", bytes.NewBuffer(jsonStr))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := positionController.Create(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusBadRequest, rec.Code)
	s.True(response.Error)
	s.Equal("warning", response.Type)
	s.EqualValues("bind error when trying to bind position", response.Message)
	s.Nil(response.Data)
}

func (s PositionTestSuite) TestPositionControllerCreateServiceFail() {
	positionController := positionController{service: services.NewPositionService()}

	e := echo.New()
	var jsonStr = []byte(`{"real": "madrid"}`)
	req := httptest.NewRequest(http.MethodPost, "/positions", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := positionController.Create(c)
	s.Nil(err)
	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusBadRequest, rec.Code)
	s.True(response.Error)
	s.Equal("warning", response.Type)
	s.EqualValues("stock id is required", response.Message)
	s.Nil(response.Data)
}

func (s PositionTestSuite) TestPositionControllerClose() {
	positionController := positionController{service: services.NewPositionService()}

	e := echo.New()
	var jsonStr = []byte(`{"sale_quantity":10, "sale_price": 15, "sale_commission": 1.0}`)
	req := httptest.NewRequest(http.MethodPost, "/positions/:position_id/close", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:position_id")
	c.SetParamNames("position_id")
	c.SetParamValues("1")

	err := positionController.Close(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusOK, rec.Code)
	s.False(response.Error)
	s.Equal("success", response.Type)
	s.EqualValues("Position closed successfully", response.Message)
	s.NotNil(response.Data)
}

func (s PositionTestSuite) TestPositionControllerCloseBindJsonFail() {
	positionController := positionController{service: services.NewPositionService()}

	e := echo.New()
	var jsonStr = []byte(`{"sale_quantity":10, "sale_price": 15, "sale_commission": 1.0}`)
	req := httptest.NewRequest(http.MethodPost, "/positions/:position_id/close", bytes.NewBuffer(jsonStr))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:position_id")
	c.SetParamNames("position_id")
	c.SetParamValues("1")

	err := positionController.Close(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusBadRequest, rec.Code)
	s.True(response.Error)
	s.Equal("warning", response.Type)
	s.EqualValues("bind error when trying to bind closed position", response.Message)
	s.Nil(response.Data)
}

func (s PositionTestSuite) TestPositionControllerCloseValidateFail() {
	positionController := positionController{service: services.NewPositionService()}

	e := echo.New()
	var jsonStr = []byte(`{"sale_quantity":10, "sale_price": 15}`)
	req := httptest.NewRequest(http.MethodPost, "/positions/:position_id/close", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:position_id")
	c.SetParamNames("position_id")
	c.SetParamValues("1")

	err := positionController.Close(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusBadRequest, rec.Code)
	s.True(response.Error)
	s.Equal("warning", response.Type)
	s.EqualValues("sale commission is required", response.Message)
	s.Nil(response.Data)
}

func (s PositionTestSuite) TestPositionControllerCloseWithoutPathParameters() {
	positionController := positionController{service: services.NewPositionService()}

	e := echo.New()
	var jsonStr = []byte(`{"sale_quantity":10, "sale_price": 15}`)
	req := httptest.NewRequest(http.MethodPost, "/positions/:position_id/close", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := positionController.Close(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusBadRequest, rec.Code)
	s.True(response.Error)
	s.Equal("warning", response.Type)
	s.EqualValues("invalid position id", response.Message)
	s.Nil(response.Data)
}

func (s PositionTestSuite) TestPositionControllerClosePositionGetFail() {
	positionController := positionController{service: services.NewPositionService()}

	e := echo.New()
	var jsonStr = []byte(`{"sale_quantity":10, "sale_price": 15}`)
	req := httptest.NewRequest(http.MethodPost, "/positions/:position_id/close", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:position_id")
	c.SetParamNames("position_id")
	c.SetParamValues("100000000")

	err := positionController.Close(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusBadRequest, rec.Code)
	s.True(response.Error)
	s.Equal("warning", response.Type)
	s.EqualValues("sale commission is required", response.Message)
	s.Nil(response.Data)
}

func TestPositionTestSuite(t *testing.T) {
	suite.Run(t, new(PositionTestSuite))
}
