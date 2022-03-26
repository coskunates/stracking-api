package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClosedPositionValidateWithoutStockId(t *testing.T) {
	closedPosition := ClosedPosition{
		PositionId:     1,
		Price:          10.0,
		Quantity:       10,
		Commission:     2.0,
		SalePrice:      12.0,
		SaleQuantity:   10,
		SaleCommission: 11.0,
	}

	err := closedPosition.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "stock id is required", err.Message)
}

func TestClosedPositionValidateWithoutPositionId(t *testing.T) {
	closedPosition := ClosedPosition{
		StockId:        1,
		Price:          10.0,
		Quantity:       10,
		Commission:     2.0,
		SalePrice:      12.0,
		SaleQuantity:   10,
		SaleCommission: 11.0,
	}

	err := closedPosition.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "position id is required", err.Message)
}

func TestClosedPositionValidateWithoutQuantity(t *testing.T) {
	closedPosition := ClosedPosition{
		StockId:        1,
		PositionId:     1,
		Price:          10.0,
		Commission:     2.0,
		SalePrice:      12.0,
		SaleQuantity:   10,
		SaleCommission: 11.0,
	}

	err := closedPosition.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "quantity is required", err.Message)
}

func TestClosedPositionValidateWithoutPrice(t *testing.T) {
	closedPosition := ClosedPosition{
		StockId:        1,
		PositionId:     1,
		Quantity:       10,
		Commission:     2.0,
		SalePrice:      12.0,
		SaleQuantity:   10,
		SaleCommission: 11.0,
	}

	err := closedPosition.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "price is required", err.Message)
}

func TestClosedPositionValidateWithoutCommission(t *testing.T) {
	closedPosition := ClosedPosition{
		StockId:        1,
		PositionId:     1,
		Price:          10.0,
		Quantity:       10,
		SalePrice:      12.0,
		SaleQuantity:   10,
		SaleCommission: 11.0,
	}

	err := closedPosition.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "commission is required", err.Message)
}

func TestClosedPositionValidateWithoutSaleQuantity(t *testing.T) {
	closedPosition := ClosedPosition{
		StockId:        1,
		PositionId:     1,
		Price:          10.0,
		Quantity:       10,
		Commission:     2.0,
		SalePrice:      12.0,
		SaleCommission: 11.0,
	}

	err := closedPosition.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "sale quantity is required", err.Message)
}

func TestClosedPositionValidateWithoutSalePrice(t *testing.T) {
	closedPosition := ClosedPosition{
		StockId:        1,
		PositionId:     1,
		Price:          10.0,
		Quantity:       10,
		Commission:     2.0,
		SaleQuantity:   10,
		SaleCommission: 11.0,
	}

	err := closedPosition.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "sale price is required", err.Message)
}

func TestClosedPositionValidateWithoutSaleCommission(t *testing.T) {
	closedPosition := ClosedPosition{
		StockId:      1,
		PositionId:   1,
		Price:        10.0,
		Quantity:     10,
		Commission:   2.0,
		SalePrice:    12.0,
		SaleQuantity: 10,
	}

	err := closedPosition.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "sale commission is required", err.Message)
}

func TestClosedPositionValidatePositionQuantityIsLessThanSaleQuantity(t *testing.T) {
	closedPosition := ClosedPosition{
		StockId:        1,
		PositionId:     1,
		Price:          10.0,
		Quantity:       10,
		Commission:     2.0,
		SalePrice:      12.0,
		SaleQuantity:   11,
		SaleCommission: 11.0,
	}

	err := closedPosition.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "position quantity (10) is less than sale quantity (11)", err.Message)
}

func TestClosedPositionValidate(t *testing.T) {
	closedPosition := ClosedPosition{
		StockId:        1,
		PositionId:     1,
		Price:          10.0,
		Quantity:       10,
		Commission:     2.0,
		SalePrice:      12.0,
		SaleQuantity:   10,
		SaleCommission: 11.0,
	}

	err := closedPosition.Validate()

	assert.Nil(t, err)
}
