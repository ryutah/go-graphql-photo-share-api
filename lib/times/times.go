package times

import (
	"time"
)

var loc *time.Location

func init() {
	var err error
	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
}

func Now() time.Time {
	return time.Now().In(loc)
}

func Date(year int, month time.Month, day, hour, min, sec, nsec int) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, loc)
}
