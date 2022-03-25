package error_utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("bad request error", 1)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, 1, err.ErrorCode)
	assert.Equal(t, "warning", err.NotificationType)
	assert.Equal(t, "bad request error", err.Message)
	assert.Nil(t, err.Causes)
	assert.Equal(t, "bad_request", err.Error)
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("internal server error", 2)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, 2, err.ErrorCode)
	assert.Equal(t, "danger", err.NotificationType)
	assert.Equal(t, "internal server error", err.Message)
	assert.Nil(t, err.Causes)
	assert.Equal(t, "internal_server_error", err.Error)
}

func TestNewForbiddenError(t *testing.T) {
	err := NewForbiddenError("forbidden error", 3)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusForbidden, err.StatusCode)
	assert.Equal(t, 3, err.ErrorCode)
	assert.Equal(t, "danger", err.NotificationType)
	assert.Equal(t, "forbidden error", err.Message)
	assert.Nil(t, err.Causes)
	assert.Equal(t, "forbidden", err.Error)
}

func TestNewTooManyRequestsError(t *testing.T) {
	err := NewTooManyRequestsError("too many requests error", 4)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusTooManyRequests, err.StatusCode)
	assert.Equal(t, 4, err.ErrorCode)
	assert.Equal(t, "info", err.NotificationType)
	assert.Equal(t, "too many requests error", err.Message)
	assert.Nil(t, err.Causes)
	assert.Equal(t, "too_many_requests", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("not found error", 5)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.StatusCode)
	assert.Equal(t, 5, err.ErrorCode)
	assert.Equal(t, "warning", err.NotificationType)
	assert.Equal(t, "not found error", err.Message)
	assert.Nil(t, err.Causes)
	assert.Equal(t, "not_found", err.Error)
}

func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError("unauthorized error", 6)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusUnauthorized, err.StatusCode)
	assert.Equal(t, 6, err.ErrorCode)
	assert.Equal(t, "danger", err.NotificationType)
	assert.Equal(t, "unauthorized error", err.Message)
	assert.Nil(t, err.Causes)
	assert.Equal(t, "unauthorized", err.Error)
}
