package date_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetNowAsString(t *testing.T) {
	nowFormat := GetNowAsString()

	assert.NotNil(t, nowFormat)
	assert.EqualValues(t, time.Now().Format("2006-01-02 15:04:05"), nowFormat)
}
