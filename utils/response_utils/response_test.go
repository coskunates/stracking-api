package response_utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"stock/models"
	"testing"
)

func TestNewErrorResponse(t *testing.T) {
	response := NewErrorResponse(http.StatusBadRequest, "warning", "bad_request")

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.True(t, response.Error)
	assert.EqualValues(t, "bad_request", response.Message)
	assert.EqualValues(t, "warning", response.Type)
	assert.Nil(t, response.Data)
}

func TestNewSuccessResponse(t *testing.T) {
	response := NewSuccessResponse(http.StatusOK, "test", models.Position{ID: 1})

	assert.Equal(t, http.StatusOK, response.Code)
	assert.False(t, response.Error)
	assert.EqualValues(t, "test", response.Message)
	assert.EqualValues(t, "success", response.Type)
	assert.NotNil(t, response.Data)
}

func TestNewSuccessResponseWithEmptyData(t *testing.T) {
	response := NewSuccessResponseWithEmptyData(http.StatusOK, "test")

	assert.Equal(t, http.StatusOK, response.Code)
	assert.False(t, response.Error)
	assert.EqualValues(t, "test", response.Message)
	assert.EqualValues(t, "success", response.Type)
	assert.Nil(t, response.Data)
}
