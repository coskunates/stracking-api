package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPositionValidateWithoutStockId(t *testing.T) {
	position := Position{Price: 10.0, Quantity: 10, Commission: 2.0}

	err := position.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "stock id is required", err.Message)
}

func TestPositionValidateWithoutQuantity(t *testing.T) {
	position := Position{StockId: 1, Price: 10.0, Commission: 2.0}

	err := position.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "quantity is required", err.Message)
}

func TestPositionValidateWithoutPrice(t *testing.T) {
	position := Position{StockId: 1, Quantity: 10, Commission: 2.0}

	err := position.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "price is required", err.Message)
}

func TestPositionValidateWithoutCommission(t *testing.T) {
	position := Position{StockId: 1, Quantity: 10, Price: 2.0}

	err := position.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "commission is required", err.Message)
}

func TestPositionValidate(t *testing.T) {
	position := Position{StockId: 1, Quantity: 10, Price: 2.0, Commission: 0.5}

	err := position.Validate()

	assert.Nil(t, err)
}
