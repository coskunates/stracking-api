package controllers

import (
	"encoding/json"
	"fmt"
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

type StockTestSuite struct {
	suite.Suite
}

// this function executes before the test suite begins execution
func (s *StockTestSuite) seedStock() {
	stocks := []models.Stock{
		{
			Name:      "Test 1",
			ShortName: "TEST1",
			Country:   "Turkey",
			Exchange:  "İstanbul",
			Currency:  "TRY",
			CreatedAt: "2022-03-26 11:00:00",
			UpdatedAt: "2022-03-26 11:00:00",
		},
	}

	database.GetClient().Create(&stocks)
}

func (s *StockTestSuite) clearStocks() {
	fmt.Println(database.GetClient().Exec("DELETE FROM stocks").Error)
	database.GetClient().Exec("DELETE FROM stocks")
}

// this function executes after all tests executed
func (s *StockTestSuite) TearDownSuite() {
	database.GetClient().Exec("TRUNCATE stocks")
}

func (s StockTestSuite) TestStockControllerList() {
	s.seedStock()
	stockController := stockController{service: services.NewStockService()}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/stocks", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := stockController.List(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusOK, rec.Code)
	s.False(response.Error)
	s.Equal("success", response.Type)
	s.EqualValues("Stocks listed", response.Message)
	s.NotNil(response.Data)
}

func (s StockTestSuite) TestStockControllerListNotFound() {
	s.clearStocks()
	stockController := stockController{service: services.NewStockService()}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/stocks", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := stockController.List(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusNotFound, rec.Code)
	s.True(response.Error)
	s.Equal("warning", response.Type)
	s.EqualValues("there is no stock", response.Message)
	s.Nil(response.Data)
}

func TestStockTestSuite(t *testing.T) {
	suite.Run(t, new(StockTestSuite))
}
