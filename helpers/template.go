package helpers

import (
	"time"
)

func DateFormat(t int64, layout string) string {
	tm := time.Unix(t/1000, 0)
	return tm.Format(layout)
}

func Substring(source string, start, end int) string {
	rs := []rune(source)
	length := len(rs)
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	return string(rs[start:end])
}

func Truncate() {

}
