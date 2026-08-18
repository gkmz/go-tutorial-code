package race

import (
	"sync/atomic"
	"time"
)

// Watchdog 使用原子时间戳记录最近一次心跳。
type Watchdog struct {
	lastSeen atomic.Int64
}

// KeepAlive 记录当前心跳时间。
func (w *Watchdog) KeepAlive(now time.Time) {
	w.lastSeen.Store(now.UnixNano())
}

// Expired 判断最近一次心跳是否早于截止时间。
func (w *Watchdog) Expired(deadline time.Time) bool {
	return w.lastSeen.Load() < deadline.UnixNano()
}
