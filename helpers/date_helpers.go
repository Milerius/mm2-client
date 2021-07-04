package helpers

import (
	"time"
)

func GetDateFromTimestamp(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.UTC().String()
}
