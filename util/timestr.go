package util

import "time"

const layout = "2006-01-02 15:04:05"

func TimeToString(t time.Time) (result string) {
	result = t.Format(layout)
	return result
}
