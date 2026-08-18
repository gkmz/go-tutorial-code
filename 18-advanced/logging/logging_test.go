package logging

import (
	"bytes"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestRedactingHandler(t *testing.T) {
	var output bytes.Buffer
	next := slog.NewJSONHandler(&output, nil)
	logger := slog.New(NewRedactingHandler(next, "password", "token"))

	logger.Info("login", slog.String("password", "secret"), slog.String("user_id", "42"))
	line := output.String()
	if strings.Contains(line, "secret") || !strings.Contains(line, "[REDACTED]") {
		t.Fatalf("redacted log = %s", line)
	}
}

func TestNewMultiLoggerWritesBothTargets(t *testing.T) {
	var console, file bytes.Buffer
	NewMultiLogger(&console, &file).Print("started")
	if !strings.Contains(console.String(), "started") || !strings.Contains(file.String(), "started") {
		t.Fatalf("console/file = %q/%q", console.String(), file.String())
	}
}

func TestHTTPMiddleware(t *testing.T) {
	var output bytes.Buffer
	logger := NewSlog(&output, slog.LevelInfo)
	next := http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		writer.WriteHeader(http.StatusCreated)
	})
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/orders", nil)
	request.Header.Set("X-Request-ID", "req-123")

	HTTPMiddleware(logger, next).ServeHTTP(recorder, request)
	line := output.String()
	for _, expected := range []string{"req-123", "POST", "/orders", `"status":201`} {
		if !strings.Contains(line, expected) {
			t.Fatalf("log %q does not contain %q", line, expected)
		}
	}
}

func TestNormalizeRequestIDRejectsControlCharacters(t *testing.T) {
	if got := normalizeRequestID("valid-id_42"); got != "valid-id_42" {
		t.Fatalf("normalizeRequestID(valid) = %q", got)
	}
	if got := normalizeRequestID("forged\nfield"); got != "" {
		t.Fatalf("normalizeRequestID(invalid) = %q, want empty", got)
	}
}

func TestZapDynamicLevel(t *testing.T) {
	var output bytes.Buffer
	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	logger := NewZap(zapcore.AddSync(&output), level)
	logger.Debug("hidden")
	level.SetLevel(zap.DebugLevel)
	logger.Debug("visible")
	if line := output.String(); strings.Contains(line, "hidden") || !strings.Contains(line, "visible") {
		t.Fatalf("dynamic level output = %q", line)
	}
}

func TestLevelHandlerRequiresBearerToken(t *testing.T) {
	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	handler := NewLevelHandler(level, "admin-token")

	unauthorized := httptest.NewRequest(http.MethodPut, "/level", strings.NewReader(`{"level":"debug"}`))
	unauthorized.Header.Set("Content-Type", "application/json")
	unauthorizedRecorder := httptest.NewRecorder()
	handler.ServeHTTP(unauthorizedRecorder, unauthorized)
	if unauthorizedRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("unauthorized status = %d", unauthorizedRecorder.Code)
	}

	authorized := httptest.NewRequest(http.MethodPut, "/level", strings.NewReader(`{"level":"debug"}`))
	authorized.Header.Set("Content-Type", "application/json")
	authorized.Header.Set("Authorization", "Bearer admin-token")
	authorizedRecorder := httptest.NewRecorder()
	handler.ServeHTTP(authorizedRecorder, authorized)
	if authorizedRecorder.Code != http.StatusOK || level.Level() != zap.DebugLevel {
		t.Fatalf("authorized status/level = %d/%s", authorizedRecorder.Code, level.Level())
	}
}

func BenchmarkStandardLog(b *testing.B) {
	logger := log.New(io.Discard, "", 0)
	for b.Loop() {
		logger.Print("request completed status=200")
	}
}

func BenchmarkSlog(b *testing.B) {
	logger := NewSlog(io.Discard, slog.LevelInfo)
	for b.Loop() {
		logger.Info("request completed", slog.Int("status", 200))
	}
}

func BenchmarkZap(b *testing.B) {
	logger := NewZap(zapcore.AddSync(io.Discard), zap.NewAtomicLevelAt(zap.InfoLevel))
	for b.Loop() {
		logger.Info("request completed", zap.Int("status", 200))
	}
}
