package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStockValidate(t *testing.T) {
	stock := Stock{
		ID:        1,
		Name:      "Turkcell",
		ShortName: "TCELL",
		Country:   "Turkey",
		Exchange:  "İstanbul",
		Currency:  "TRY",
		CreatedAt: "2022-03-26 14:00:00",
		UpdatedAt: "2022-03-26 14:00:00",
	}

	err := stock.Validate()

	assert.Nil(t, err)
}

func TestStockValidateWithoutName(t *testing.T) {
	stock := Stock{
		ID:        1,
		ShortName: "TCELL",
		Country:   "Turkey",
		Exchange:  "İstanbul",
		Currency:  "TRY",
		CreatedAt: "2022-03-26 14:00:00",
		UpdatedAt: "2022-03-26 14:00:00",
	}

	err := stock.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "stock name is required", err.Message)
}

func TestStockValidateWithoutShortName(t *testing.T) {
	stock := Stock{
		ID:        1,
		Name:      "Turkcell",
		Country:   "Turkey",
		Exchange:  "İstanbul",
		Currency:  "TRY",
		CreatedAt: "2022-03-26 14:00:00",
		UpdatedAt: "2022-03-26 14:00:00",
	}

	err := stock.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "stock short name is required", err.Message)
}

func TestStockValidateWithoutCountry(t *testing.T) {
	stock := Stock{
		ID:        1,
		Name:      "Turkcell",
		ShortName: "TCELL",
		Exchange:  "İstanbul",
		Currency:  "TRY",
		CreatedAt: "2022-03-26 14:00:00",
		UpdatedAt: "2022-03-26 14:00:00",
	}

	err := stock.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "stock country is required", err.Message)
}
