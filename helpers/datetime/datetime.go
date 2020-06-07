package datetime

import "time"

var (
	WeekStartDay    = time.Monday
	DateTimeFormats = []string{"1/2/2006", "1/2/2006 15:4:5", "2006", "2006-1-2", "2006-01-02 15:04:05", "20060102150405", "15:4:5 Jan 2, 2006 MST"}
)

const (
	DefaultFormat  = "2006-01-02 15:04:05"
	CompressFormat = "20060102150405"
)

type DateTime struct {
	time.Time
}

func CurrentSecond() time.Time {
	return time.Now().Truncate(time.Second)
}
