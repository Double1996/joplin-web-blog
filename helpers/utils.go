package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/snluu/uuid"
)

func Panicf(format string, values ...interface{}) {
	if _, file, line, ok := runtime.Caller(1); ok {
		format += "\n\t in %d line of the %q file"
		values = append(values, line, file)
	}

	panic(fmt.Sprintf(format, values...))
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func Truncate(s string, n int) string {
	runes := []rune(s)
	if len(runes) > n {
		return string(runes[:n])
	}
	return s
}

func UUID() string {
	return uuid.Rand().Hex()
}
