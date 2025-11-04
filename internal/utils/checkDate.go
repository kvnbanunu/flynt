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

func DaysSince(t time.Time, timezone string) (int, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}
	now := time.Now().In(loc)
	t = t.In(loc)
	return int(now.Sub(t).Hours() / 24), nil
}
