package controllers

import (
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

type SummaryTestSuite struct {
	suite.Suite
}

// this function executes before the test suite begins execution
func (s *SummaryTestSuite) seed() {
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

	positions := []models.Position{
		{
			ID:         1,
			StockId:    1,
			Quantity:   10,
			Price:      5.0,
			Commission: 5.0,
			OpenedAt:   "2022-03-25",
			CreatedAt:  "2022-03-25 14:00:00",
			UpdatedAt:  "2022-03-25 14:00:00",
		},
		{
			ID:         2,
			StockId:    1,
			Quantity:   20,
			Price:      10.0,
			Commission: 10.0,
			OpenedAt:   "2022-03-25",
			CreatedAt:  "2022-03-25 14:00:00",
			UpdatedAt:  "2022-03-25 14:00:00",
		},
		{
			ID:         3,
			StockId:    1,
			Quantity:   30,
			Price:      15.0,
			Commission: 15.0,
			OpenedAt:   "2022-03-25",
			CreatedAt:  "2022-03-25 14:00:00",
			UpdatedAt:  "2022-03-25 14:00:00",
		},
	}

	database.GetClient().Create(&positions)

	closedPositions := []models.ClosedPosition{
		{
			ID:             1,
			PositionId:     1,
			StockId:        1,
			Quantity:       10,
			Price:          5.0,
			Commission:     5.0,
			SaleQuantity:   10,
			SalePrice:      5.0,
			SaleCommission: 5,
			OpenedAt:       "2022-03-25",
			ClosedAt:       "2022-03-25 14:00:00",
			CreatedAt:      "2022-03-25 14:00:00",
			UpdatedAt:      "2022-03-25 14:00:00",
		},
	}

	database.GetClient().Create(&closedPositions)
}

// this function executes after all tests executed
func (s *SummaryTestSuite) clear() {
	database.GetClient().Exec("TRUNCATE stocks")
	database.GetClient().Exec("TRUNCATE positions")
	database.GetClient().Exec("TRUNCATE closed_positions")
}

func (s *SummaryTestSuite) TestSummaryControllerGetSummary() {
	s.seed()
	summaryController := summaryController{service: services.NewSummaryService()}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/summaries", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := summaryController.GetSummary(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusOK, rec.Code)
	s.False(response.Error)
	s.Equal("success", response.Type)
	s.EqualValues("Summary", response.Message)
	s.NotNil(response.Data)
	s.clear()
}

func (s *SummaryTestSuite) TestSummaryControllerGetSummaryServiceError() {
	summaryController := summaryController{service: services.NewSummaryService()}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/summaries", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := summaryController.GetSummary(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusNotFound, rec.Code)
	s.True(response.Error)
	s.Equal("warning", response.Type)
	s.EqualValues("no summary information", response.Message)
	s.Nil(response.Data)
}

func (s *SummaryTestSuite) TestSummaryControllerGetOpenPositionSummary() {
	s.seed()
	summaryController := summaryController{service: services.NewSummaryService()}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/summaries/opened", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := summaryController.GetOpenPositionSummary(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusOK, rec.Code)
	s.False(response.Error)
	s.Equal("success", response.Type)
	s.EqualValues("Open positions summary", response.Message)
	s.NotNil(response.Data)
	s.clear()
}

func (s *SummaryTestSuite) TestSummaryControllerGetOpenPositionSummaryServiceError() {
	summaryController := summaryController{service: services.NewSummaryService()}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/summaries/closed", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := summaryController.GetClosedPositionSummary(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusNotFound, rec.Code)
	s.True(response.Error)
	s.Equal("warning", response.Type)
	s.EqualValues("there is no closed position", response.Message)
	s.Nil(response.Data)
}

func (s *SummaryTestSuite) TestSummaryControllerGetClosedPositionSummary() {
	s.seed()
	summaryController := summaryController{service: services.NewSummaryService()}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/summaries/closed", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := summaryController.GetClosedPositionSummary(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusOK, rec.Code)
	s.False(response.Error)
	s.Equal("success", response.Type)
	s.EqualValues("Closed positions summary", response.Message)
	s.NotNil(response.Data)
	s.clear()
}

func (s *SummaryTestSuite) TestSummaryControllerGetClosedPositionServiceError() {
	summaryController := summaryController{service: services.NewSummaryService()}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/summaries/closed", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := summaryController.GetClosedPositionSummary(c)
	s.Nil(err)

	var response response_utils.Response
	_ = json.Unmarshal([]byte(rec.Body.String()), &response)

	s.EqualValues(http.StatusNotFound, rec.Code)
	s.True(response.Error)
	s.Equal("warning", response.Type)
	s.EqualValues("there is no closed position", response.Message)
	s.Nil(response.Data)
}

func TestSummaryTestSuite(t *testing.T) {
	suite.Run(t, new(SummaryTestSuite))
}
