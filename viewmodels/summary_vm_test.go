package viewmodels

import (
	"github.com/stretchr/testify/assert"
	"stock/models"
	"testing"
)

func TestSummaryVMBindPositionsToSummaryVM(t *testing.T) {
	var summaryVM SummaryVM
	positions := []models.Position{
		{
			ID:      1,
			StockId: 1,
			Stock: models.Stock{
				ID:        1,
				Name:      "Test Stock",
				ShortName: "TEST1",
				Country:   "Turkey",
				Sector:    "Finance",
				Exchange:  "",
				Currency:  "TRY",
				CreatedAt: "2022-03-28 14:00:00",
				UpdatedAt: "2022-03-28 14:00:00",
			},
			Quantity:   10,
			Price:      10.0,
			Commission: 5.0,
			OpenedAt:   "2022-03-05",
			CreatedAt:  "2022-03-28 14:00:00",
			UpdatedAt:  "2022-03-28 14:00:00",
		},
	}
	response := summaryVM.BindPositionsToSummaryVM(&positions)

	assert.NotNil(t, response.Items)
	assert.EqualValues(t, 1, len(response.Items))
	assert.EqualValues(t, 100.0, response.SubTotalPrice)
	assert.EqualValues(t, 5.0, response.TotalCommissionPrice)
	assert.EqualValues(t, 105.0, response.TotalPrice)
}

func TestSummaryVMBindClosedPositionsToSummaryVM(t *testing.T) {
	var summaryVM SummaryVM
	closedPositions := []models.ClosedPosition{
		{
			ID:         1,
			PositionId: 1,
			StockId:    1,
			Stock: models.Stock{
				ID:        1,
				Name:      "Test Stock",
				ShortName: "TEST1",
				Country:   "Turkey",
				Sector:    "Finance",
				Exchange:  "",
				Currency:  "TRY",
				CreatedAt: "2022-03-28 14:00:00",
				UpdatedAt: "2022-03-28 14:00:00",
			},
			Quantity:       10,
			SaleQuantity:   5,
			Price:          10.0,
			SalePrice:      15.0,
			Commission:     5.0,
			SaleCommission: 5.0,
			OpenedAt:       "2022-03-05",
			CreatedAt:      "2022-03-28 14:00:00",
			UpdatedAt:      "2022-03-28 14:00:00",
		},
	}
	response := summaryVM.BindClosedPositionsToSummaryVM(&closedPositions)

	assert.NotNil(t, response.Items)
	assert.EqualValues(t, 1, len(response.Items))
	assert.EqualValues(t, 100.0, response.SubTotalPrice)
	assert.EqualValues(t, 5.0, response.TotalCommissionPrice)
	assert.EqualValues(t, 105.0, response.TotalPrice)
}
