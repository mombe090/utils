package dates_utils

import "time"

const format = "01-02-2006 15:04:05Z"

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(format)
}
