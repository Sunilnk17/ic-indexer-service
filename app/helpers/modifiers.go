package helpers

import (
	"ic-indexer-service/app/model/request"
	"time"
)

func GetArrayString(cusArrayString request.CusArrayString, existing *[]string) *[]string {
	if cusArrayString.Set {
		return cusArrayString.Value
	}
	return existing
}

func GetString(cusString request.CusString, existing *string) *string {
	if cusString.Set {
		return cusString.Value
	}
	return existing
}

func GetInt64(passed request.CusInt64, existing *int64) *int64 {
	if passed.Set {
		return passed.Value
	}
	return existing
}

func GetTime(cusTime request.CusTime, existing *time.Time) *time.Time {
	if cusTime.Set && !cusTime.Value.Equal(time.Time{}) {
		return cusTime.Value
	}
	return existing
}
