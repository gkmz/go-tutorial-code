// Package logging 提供结构化日志、敏感字段过滤和 HTTP 请求日志示例。
package logging

import (
	"io"
	"log"
	"log/slog"
)

// NewMultiLogger 创建同时写入控制台和文件目标的文本 Logger。
func NewMultiLogger(console, file io.Writer) *log.Logger {
	return log.New(io.MultiWriter(console, file), "service ", log.LstdFlags|log.LUTC)
}

// NewSlog 创建输出 JSON 的 slog.Logger。
func NewSlog(output io.Writer, level slog.Leveler) *slog.Logger {
	handler := slog.NewJSONHandler(output, &slog.HandlerOptions{Level: level})
	return slog.New(handler)
}

// WithRequestID 返回包含请求标识的子 Logger。
func WithRequestID(logger *slog.Logger, requestID string) *slog.Logger {
	return logger.With(slog.String("request_id", requestID))
}
