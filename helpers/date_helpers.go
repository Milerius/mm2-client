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

func GetDateFromTimestampStandard(timestamp int64) string {
	//fmt.Println(timestamp)
	tm := time.Unix(0, timestamp)
	return tm.UTC().Format(time.RFC3339)
}

func GetDateFromTime(timestamp time.Time) string {
	return timestamp.UTC().Format(time.RFC3339)
}

func RFCDateToTimestamp(date string) int64 {
	layout := "2006-01-02T15:04:05Z"
	parse, err := time.Parse(layout, date)
	if err != nil {
		return 0
	}
	return parse.Unix()
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

func DateToTimeElapsed(date string) float64 {
	cur, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return (time.Minute * 5).Seconds()
	}
	elapsed := time.Since(cur)
	return elapsed.Seconds()
}

func RFC3339ToTimestamp(date string) int64 {
	parse, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return 0
	}
	return parse.UnixNano()
}

func RFC3339ToTimestampSecond(date string) int64 {
	parse, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return 0
	}
	return parse.Unix()
}
