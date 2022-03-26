package error_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorCodesConstants(t *testing.T) {
	assert.Equal(t, 1, JsonBindError)
	assert.Equal(t, 2, InvalidRequestParams)
	assert.Equal(t, 3, DatabaseCreateError)
}
