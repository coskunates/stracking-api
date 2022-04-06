package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"stock/database"
	"stock/services"
	"stock/utils/response_utils"
	"testing"
)

type DividendTestSuite struct {
	suite.Suite
}

// this function executes after all tests executed
func (d *DividendTestSuite) TearDownSuite() {
	database.GetClient().Exec("TRUNCATE dividends")
}

func (d *DividendTestSuite) TestDividendControllerCreate() {
	dividendController := dividendController{service: services.NewDividendService()}

	e := echo.New()
	var jsonStr = []byte(`{"currency_id":1, "stock_id": 1, "quantity": 10, "price": 15.0, "issued_at": "2022-04-06"}`)
	req := httptest.NewRequest(http.MethodPost, "/dividends", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := dividendController.Create(c)
	d.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	d.EqualValues(http.StatusCreated, rec.Code)
	d.False(response.Error)
	d.Equal("success", response.Type)
	d.EqualValues("Dividend created", response.Message)
	d.NotNil(response.Data)
}

func (d *DividendTestSuite) TestDividendControllerCreateBindJsonFail() {
	dividendController := dividendController{service: services.NewDividendService()}

	e := echo.New()
	var jsonStr = []byte(`{"stock_id": 1, "quantity": 10, "price": 15.0}`)
	req := httptest.NewRequest(http.MethodPost, "/dividends", bytes.NewBuffer(jsonStr))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := dividendController.Create(c)
	d.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	d.EqualValues(http.StatusBadRequest, rec.Code)
	d.True(response.Error)
	d.Equal("warning", response.Type)
	d.EqualValues("bind error when trying to bind dividend", response.Message)
	d.Nil(response.Data)
}

func (d *DividendTestSuite) TestDividendControllerCreateValidateFail() {
	dividendController := dividendController{service: services.NewDividendService()}

	e := echo.New()
	var jsonStr = []byte(`{"stock_id": 1, "quantity": 10}`)
	req := httptest.NewRequest(http.MethodPost, "/dividends", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := dividendController.Create(c)
	d.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	d.EqualValues(http.StatusBadRequest, rec.Code)
	d.True(response.Error)
	d.Equal("warning", response.Type)
	d.EqualValues("price is required", response.Message)
	d.Nil(response.Data)
}

func TestDividendTestSuite(t *testing.T) {
	suite.Run(t, new(DividendTestSuite))
}
