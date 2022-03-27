package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDividendValidate(t *testing.T) {
	dividend := Dividend{
		ID:        1,
		StockId:   1,
		Quantity:  10,
		Price:     1.35,
		CreatedAt: "2022-03-26 14:00:00",
		UpdatedAt: "2022-03-26 14:00:00",
	}

	err := dividend.Validate()

	assert.Nil(t, err)
}

func TestDividendValidateWithoutStockId(t *testing.T) {
	dividend := Dividend{
		ID:        1,
		Quantity:  10,
		Price:     1.35,
		CreatedAt: "2022-03-26 14:00:00",
		UpdatedAt: "2022-03-26 14:00:00",
	}

	err := dividend.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "stock id is required", err.Message)
}

func TestDividendValidateWithoutQuantity(t *testing.T) {
	dividend := Dividend{
		ID:        1,
		StockId:   1,
		Price:     1.35,
		CreatedAt: "2022-03-26 14:00:00",
		UpdatedAt: "2022-03-26 14:00:00",
	}

	err := dividend.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "quantity is required", err.Message)
}

func TestDividendValidateWithoutPrice(t *testing.T) {
	dividend := Dividend{
		ID:        1,
		StockId:   1,
		Quantity:  10,
		CreatedAt: "2022-03-26 14:00:00",
		UpdatedAt: "2022-03-26 14:00:00",
	}

	err := dividend.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "price is required", err.Message)
}
