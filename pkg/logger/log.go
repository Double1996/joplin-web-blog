package logger

import (
	"github.com/gin-gonic/gin"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
)

var (
	once      sync.Once
	singleton *zap.Logger
)

func LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("ip", c.ClientIP()),
			zap.Duration("latency-ms", latency*1000),
			zap.String("error", c.Errors.String()),
		)
	}
}

func newLogger(path string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	// 如果环境为debug
	if os.Getenv("IS_DEBUG") == "1" {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	cfg.OutputPaths = []string{
		path,
	}

	// 关闭采样
	cfg.Sampling = nil
	return cfg.Build()
}

func init() {
	once.Do(func() {
		singleton, _ = newLogger("stdout")
	})
}

func Info(message string, field ...zap.Field) {
	singleton.Info(message, field...)
}

func Debug(message string, field ...zap.Field) {
	singleton.Debug(message, field...)
}

func Error(message string, fields ...zap.Field) {
	singleton.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	singleton.Fatal(message, fields...)
}
