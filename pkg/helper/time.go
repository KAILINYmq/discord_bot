package helper

import (
	"math"
	"time"
)

func ParseTime(timeStr string) (*time.Time, error) {
	layout := []string{
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05",
	}
	var tt time.Time
	var err error

	for _, l := range layout {
		tt, err = time.Parse(l, timeStr)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	return &tt, nil
}

func ParseTimeString(t time.Time) string {
	curTime := time.Now().UTC()
	td := curTime.Sub(t)
	hh := math.Floor(td.Hours())

	if hh < 0 {
		return ""
	} else if hh < 24 {
		return copyWords(hh, "hour")
	} else if hh < 24*30 {
		return copyWords(math.Floor(hh/24), "day")
	} else if hh < 24*365 {
		return copyWords(math.Floor(hh/(24*30)), "month")
	} else {
		return copyWords(math.Floor(hh/(24*365)), "year")
	}
}

func ReportTime() string {
	ct := time.Now().UTC().Add(time.Hour * 8)

	return ct.Format("2006-01-02 15:04:05.000")
}
