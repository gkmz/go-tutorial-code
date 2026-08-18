package logging

import (
	"crypto/subtle"
	"net/http"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// NewLevelHandler 创建使用 Bearer Token 保护的 zap 动态级别接口。
func NewLevelHandler(level zap.AtomicLevel, token string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		provided := strings.TrimPrefix(request.Header.Get("Authorization"), "Bearer ")
		if token == "" || subtle.ConstantTimeCompare([]byte(provided), []byte(token)) != 1 {
			http.Error(writer, "unauthorized", http.StatusUnauthorized)
			return
		}
		level.ServeHTTP(writer, request)
	})
}

// NewZap 创建使用 JSON 编码和动态日志级别的 zap.Logger。
func NewZap(output zapcore.WriteSyncer, level zap.AtomicLevel) *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), output, level)
	return zap.New(core)
}

// NewRotatingZap 创建写入轮转文件的 zap.Logger。
func NewRotatingZap(filename string, maxSizeMB, maxBackups, maxAgeDays int, level zap.AtomicLevel) *zap.Logger {
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSizeMB,
		MaxBackups: maxBackups,
		MaxAge:     maxAgeDays,
		Compress:   true,
	})
	return NewZap(writer, level)
}
