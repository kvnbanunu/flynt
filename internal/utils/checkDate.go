package utils

import (
	"time"
	_ "time/tzdata"
)

func CheckDayPassed(t time.Time, timezone string) (bool, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return false, err
	}
	if time.Now().In(loc).Day() > t.In(loc).Day() {
		return true, nil
	}
	return false, nil
}
