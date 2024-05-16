package logs

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = customTimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	var err error
	logger, err = config.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
}

func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	logger.Warn(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		logger.Error(v.Error(), fields...)
	case string:
		logger.Error(v, fields...)
	}
}
func Fatal(message string, fields ...zap.Field) {
	logger.Fatal(message, fields...)
}

func Infof(format string, args ...interface{}) {
	logger.Info((fmt.Sprintf(format, args...)))
}
func Warnf(format string, args ...interface{}) {
	logger.Warn((fmt.Sprintf(format, args...)))
}
func Debugf(format string, args ...interface{}) {
	logger.Debug((fmt.Sprintf(format, args...)))
}

func Errorf(format string, args ...interface{}) {
	logger.Error((fmt.Sprintf(format, args...)))
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatal(fmt.Sprintf(format, args...))
}

func LogHttp(c *fiber.Ctx) error {
	Infof("HTTP request - status: %d, method: %s, path: %s, ip: %s",
		c.Response().StatusCode(),
		c.Method(),
		c.Path(),
		c.IP(),
	)
	start := time.Now()

	err := c.Next()

	duration := time.Since(start)

	Infof("HTTP response - status: %d, method: %s, path: %s, ip: %s, duration: %s",
		c.Response().StatusCode(),
		c.Method(),
		c.Path(),
		c.IP(),
		duration,
	)
	return err
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
