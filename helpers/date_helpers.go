package helpers

import (
	"time"
)

func GetDateFromTimestamp(timestamp int64, withoutSecond bool) string {
	tm := time.Unix(timestamp, 0)
	if withoutSecond {
		return tm.UTC().Format("2 Jan 2006 15:04")
	}
	return tm.UTC().Format("2 Jan 2006 15:04:05")
}

func DateToTimestamp(date string, withoutSecond bool) int64 {
	layout := "2 Jan 2006 15:04:05"
	if withoutSecond {
		layout = "2 Jan 2006 15:04"
	}
	parse, err := time.Parse(layout, date)
	if err != nil {
		return 0
	}
	return parse.Unix()
}

func SimpleDateToTimestamp(date string) int64 {
	layout := "02-01-2006"
	parse, err := time.Parse(layout, date)
	if err != nil {
		return 0
	}
	return parse.Unix()
}
