package helper

import (
	"fmt"
	"runtime"
)

func Panicf(format string, values ...interface{}) {
	if _, file, line, ok := runtime.Caller(1); ok {
		format += "\n\t in %d line of the %q file"
		values = append(values, line, file)
	}

	panic(fmt.Sprintf(format, values...))
}
