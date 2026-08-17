// Package webexamples 提供 HTTP 服务工程化面试的标准库示例。
package webexamples

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"
	"time"
)

// ServerConfig 描述 HTTP Server 的基础超时配置。
type ServerConfig struct {
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

// NewServer 根据配置创建带超时的 HTTP Server。
func NewServer(addr string, handler http.Handler, config ServerConfig) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
		ReadTimeout:       config.ReadTimeout,
		WriteTimeout:      config.WriteTimeout,
		IdleTimeout:       config.IdleTimeout,
	}
}

// RequestIDMiddleware 为请求补充请求 ID，并通过响应头返回。
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = newRequestID()
		}
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), requestIDKey{}, requestID)))
	})
}

type requestIDKey struct{}

// RequestID 从 Context 中读取请求 ID。
func RequestID(ctx context.Context) string {
	value, _ := ctx.Value(requestIDKey{}).(string)
	return value
}

func newRequestID() string {
	var value [8]byte
	if _, err := rand.Read(value[:]); err != nil {
		return "generated"
	}
	return hex.EncodeToString(value[:])
}

// Broadcaster 是一个带慢消费者隔离策略的进程内广播器。
type Broadcaster[T any] struct {
	mu          sync.RWMutex
	nextID      int
	subscribers map[int]chan T
	buffer      int
}

// NewBroadcaster 创建广播器；buffer 是每个订阅者的缓冲容量。
func NewBroadcaster[T any](buffer int) *Broadcaster[T] {
	if buffer < 1 {
		buffer = 1
	}
	return &Broadcaster[T]{subscribers: make(map[int]chan T), buffer: buffer}
}

// Subscribe 创建一个订阅；调用方必须调用返回的 cancel 函数释放资源。
func (b *Broadcaster[T]) Subscribe() (<-chan T, func()) {
	b.mu.Lock()
	id := b.nextID
	b.nextID++
	channel := make(chan T, b.buffer)
	b.subscribers[id] = channel
	b.mu.Unlock()

	var once sync.Once
	cancel := func() {
		once.Do(func() {
			b.mu.Lock()
			if existing, ok := b.subscribers[id]; ok {
				delete(b.subscribers, id)
				close(existing)
			}
			b.mu.Unlock()
		})
	}
	return channel, cancel
}

// Publish 向订阅者广播消息；慢消费者的满队列会丢弃当前消息。
func (b *Broadcaster[T]) Publish(value T) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, channel := range b.subscribers {
		select {
		case channel <- value:
		default:
		}
	}
}
