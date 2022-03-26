package date_utils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now()
}

func GetNowAsString() string {
	return GetNow().Format(apiDateLayout)
}
