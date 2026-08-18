package logging

import (
	"context"
	"log/slog"
	"strings"
)

// RedactingHandler 在日志进入底层 Handler 前遮盖指定字段。
type RedactingHandler struct {
	next      slog.Handler
	sensitive map[string]struct{}
}

// NewRedactingHandler 创建敏感字段过滤 Handler，字段名不区分大小写。
func NewRedactingHandler(next slog.Handler, keys ...string) *RedactingHandler {
	sensitive := make(map[string]struct{}, len(keys))
	for _, key := range keys {
		sensitive[strings.ToLower(key)] = struct{}{}
	}
	return &RedactingHandler{next: next, sensitive: sensitive}
}

// Enabled 把级别判断委托给底层 Handler。
func (h *RedactingHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.next.Enabled(ctx, level)
}

// Handle 复制 Record，并递归过滤普通属性和分组属性。
func (h *RedactingHandler) Handle(ctx context.Context, record slog.Record) error {
	filtered := slog.NewRecord(record.Time, record.Level, record.Message, record.PC)
	record.Attrs(func(attr slog.Attr) bool {
		filtered.AddAttrs(h.redact(attr))
		return true
	})
	return h.next.Handle(ctx, filtered)
}

// WithAttrs 返回携带固定属性的新 Handler。
func (h *RedactingHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	filtered := make([]slog.Attr, 0, len(attrs))
	for _, attr := range attrs {
		filtered = append(filtered, h.redact(attr))
	}
	return &RedactingHandler{next: h.next.WithAttrs(filtered), sensitive: h.sensitive}
}

// WithGroup 返回进入指定字段组的新 Handler。
func (h *RedactingHandler) WithGroup(name string) slog.Handler {
	return &RedactingHandler{next: h.next.WithGroup(name), sensitive: h.sensitive}
}

func (h *RedactingHandler) redact(attr slog.Attr) slog.Attr {
	if _, ok := h.sensitive[strings.ToLower(attr.Key)]; ok {
		return slog.String(attr.Key, "[REDACTED]")
	}
	if attr.Value.Kind() != slog.KindGroup {
		return attr
	}
	group := attr.Value.Group()
	filtered := make([]slog.Attr, 0, len(group))
	for _, child := range group {
		filtered = append(filtered, h.redact(child))
	}
	return slog.Group(attr.Key, attrsToAny(filtered)...)
}

func attrsToAny(attrs []slog.Attr) []any {
	values := make([]any, len(attrs))
	for index, attr := range attrs {
		values[index] = attr
	}
	return values
}
