package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPositionValidateWithoutStockId(t *testing.T) {
	position := Position{Price: 10.0, Quantity: 10, Commission: 2.0}

	err := position.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "code=400, message=stock id is required", err.Error())
}

func TestPositionValidateWithoutQuantity(t *testing.T) {
	position := Position{StockId: 1, Price: 10.0, Commission: 2.0}

	err := position.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "code=400, message=quantity is required", err.Error())
}

func TestPositionValidateWithoutPrice(t *testing.T) {
	position := Position{StockId: 1, Quantity: 10, Commission: 2.0}

	err := position.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "code=400, message=price is required", err.Error())
}

func TestPositionValidateWithoutCommission(t *testing.T) {
	position := Position{StockId: 1, Quantity: 10, Price: 2.0}

	err := position.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "code=400, message=commission is required", err.Error())
}

func TestPositionValidate(t *testing.T) {
	position := Position{StockId: 1, Quantity: 10, Price: 2.0, Commission: 0.5}

	err := position.Validate()

	assert.Nil(t, err)
}
