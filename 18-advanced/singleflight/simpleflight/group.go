// Package simpleflight 提供用于讲解请求合并机制的教学实现。
//
// 它会把 panic 和 runtime.Goexit 转换为错误发布给等待者，因此不是
// golang.org/x/sync/singleflight 的生产替代品。
package simpleflight

import (
	"errors"
	"fmt"
	"runtime/debug"
	"sync"
)

// ErrGoexit 表示共享函数调用了 runtime.Goexit。
var ErrGoexit = errors.New("shared function called runtime.Goexit")

// PanicError 保存教学实现从共享函数捕获的 panic 值和栈。
type PanicError struct {
	Value any
	Stack []byte
}

// Error 返回 panic 的文本描述。
func (e *PanicError) Error() string {
	return fmt.Sprintf("shared function panicked: %v", e.Value)
}

// Result 保存 DoChan 向单个等待者发布的结果。
type Result struct {
	Value  any
	Err    error
	Shared bool
}

type call struct {
	done sync.WaitGroup

	// value 和 err 在 done.Done 之前写入，等待者在 done.Wait 之后读取。
	value any
	err   error

	// duplicates 和 channels 只在 Group.mu 保护下修改。
	duplicates int
	channels   []chan<- Result
}

// Group 管理按 key 索引的进行中调用。零值可以直接使用，使用后不得复制。
type Group struct {
	mu    sync.Mutex
	calls map[string]*call
}

// Do 同步执行共享函数。同 key 已有调用时，当前调用等待并取得相同结果。
func (g *Group) Do(key string, fn func() (any, error)) (value any, err error, shared bool) {
	if fn == nil {
		return nil, errors.New("shared function must not be nil"), false
	}

	g.mu.Lock()
	if g.calls == nil {
		g.calls = make(map[string]*call)
	}
	if current, ok := g.calls[key]; ok {
		current.duplicates++
		g.mu.Unlock()
		current.done.Wait()
		return current.value, current.err, true
	}

	current := &call{}
	current.done.Add(1)
	g.calls[key] = current
	g.mu.Unlock()

	g.execute(current, key, fn)
	return current.value, current.err, current.duplicates > 0
}

// DoChan 异步执行共享函数，并返回只会接收一个 Result 的 Channel。
// 返回的 Channel 不会关闭。
func (g *Group) DoChan(key string, fn func() (any, error)) <-chan Result {
	resultCh := make(chan Result, 1)
	if fn == nil {
		resultCh <- Result{Err: errors.New("shared function must not be nil")}
		return resultCh
	}

	g.mu.Lock()
	if g.calls == nil {
		g.calls = make(map[string]*call)
	}
	if current, ok := g.calls[key]; ok {
		current.duplicates++
		current.channels = append(current.channels, resultCh)
		g.mu.Unlock()
		return resultCh
	}

	current := &call{channels: []chan<- Result{resultCh}}
	current.done.Add(1)
	g.calls[key] = current
	g.mu.Unlock()

	go g.execute(current, key, fn)
	return resultCh
}

// Forget 删除 key 的进行中索引，使后续调用可以创建新的 call。
// 它不会取消旧调用，也不会唤醒旧调用的等待者。
func (g *Group) Forget(key string) {
	g.mu.Lock()
	delete(g.calls, key)
	g.mu.Unlock()
}

func (g *Group) execute(current *call, key string, fn func() (any, error)) {
	normalReturn := false
	panicRecovered := false

	defer func() {
		if !normalReturn && !panicRecovered {
			current.err = ErrGoexit
		}

		g.mu.Lock()
		current.done.Done()
		// Forget 后可能已经存在新 call，旧调用只能删除自己的索引。
		if g.calls[key] == current {
			delete(g.calls, key)
		}
		channels := append([]chan<- Result(nil), current.channels...)
		result := Result{
			Value:  current.value,
			Err:    current.err,
			Shared: current.duplicates > 0,
		}
		g.mu.Unlock()

		// 每个 Channel 容量都是 1，即使等待者已经离开也不会阻塞结果发布。
		for _, resultCh := range channels {
			resultCh <- result
		}
	}()

	func() {
		defer func() {
			if !normalReturn {
				if recovered := recover(); recovered != nil {
					current.err = &PanicError{Value: recovered, Stack: debug.Stack()}
				}
			}
		}()

		current.value, current.err = fn()
		normalReturn = true
	}()

	if !normalReturn {
		panicRecovered = true
	}
}
