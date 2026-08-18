package logging

import (
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/felixge/httpsnoop"
)

// HTTPMiddleware 记录 HTTP 请求方法、路径、状态码、耗时和 request ID。
func HTTPMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requestID := normalizeRequestID(request.Header.Get("X-Request-ID"))
		if requestID == "" {
			requestID = newRequestID()
		}

		startedAt := time.Now()
		// httpsnoop 会保留 Flusher、Hijacker 等可选 ResponseWriter 接口。
		metrics := httpsnoop.CaptureMetrics(next, writer, request)

		WithRequestID(logger, requestID).InfoContext(request.Context(), "http request completed",
			slog.String("method", request.Method),
			slog.String("path", request.URL.Path),
			slog.Int("status", metrics.Code),
			slog.Int64("duration_ms", time.Since(startedAt).Milliseconds()),
		)
	})
}

func normalizeRequestID(value string) string {
	if value == "" || len(value) > 128 {
		return ""
	}
	if strings.IndexFunc(value, func(character rune) bool {
		return !((character >= 'a' && character <= 'z') ||
			(character >= 'A' && character <= 'Z') ||
			(character >= '0' && character <= '9') ||
			character == '-' || character == '_' || character == '.')
	}) >= 0 {
		return ""
	}
	return value
}

func newRequestID() string {
	var data [8]byte
	if _, err := rand.Read(data[:]); err != nil {
		return "request-id-unavailable"
	}
	return hex.EncodeToString(data[:])
}
