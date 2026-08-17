// Package ratelimiter 提供可取消的单机令牌桶限流器。
package ratelimiter

import (
	"context"
	"sync"
	"time"
)

// Limiter 使用 Ticker 补充令牌，并用 channel 容量限制令牌积累。
type Limiter struct {
	tokens chan struct{}
	stop   chan struct{}
	once   sync.Once
}

// New 创建每秒补充 rate 个令牌、最多缓存 burst 个令牌的限流器。
func New(rate, burst int) *Limiter {
	if rate <= 0 || burst <= 0 {
		panic("rate and burst must be positive")
	}
	limiter := &Limiter{tokens: make(chan struct{}, burst), stop: make(chan struct{})}
	for i := 0; i < burst; i++ {
		limiter.tokens <- struct{}{}
	}
	go limiter.refill(rate)
	return limiter
}

// Wait 等待一个令牌；Context 取消时返回 Context 错误。
func (l *Limiter) Wait(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-l.tokens:
		return nil
	}
}

// Close 停止后台补充 Goroutine。
func (l *Limiter) Close() {
	l.once.Do(func() { close(l.stop) })
}

func (l *Limiter) refill(rate int) {
	ticker := time.NewTicker(time.Second / time.Duration(rate))
	defer ticker.Stop()
	for {
		select {
		case <-l.stop:
			return
		case <-ticker.C:
			select {
			case l.tokens <- struct{}{}:
			default:
			}
		}
	}
}
