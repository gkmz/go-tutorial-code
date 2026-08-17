package simpleflight

import (
	"errors"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func TestDoChanMergesOverlappingCalls(t *testing.T) {
	const callers = 20
	var group Group
	var executions atomic.Int32
	release := make(chan struct{})
	results := make([]<-chan Result, 0, callers)

	for range callers {
		results = append(results, group.DoChan("key", func() (any, error) {
			executions.Add(1)
			<-release
			return "value", nil
		}))
	}
	close(release)

	for _, resultCh := range results {
		result := <-resultCh
		if result.Value != "value" || result.Err != nil || !result.Shared {
			t.Fatalf("DoChan() result = %+v", result)
		}
	}
	if got := executions.Load(); got != 1 {
		t.Fatalf("executions = %d, want 1", got)
	}
}

func TestDoSharesError(t *testing.T) {
	var group Group
	wantErr := errors.New("load failed")
	release := make(chan struct{})
	first := group.DoChan("key", func() (any, error) {
		<-release
		return nil, wantErr
	})
	second := group.DoChan("key", func() (any, error) {
		t.Fatal("duplicate function should not execute")
		return nil, nil
	})
	close(release)

	for _, resultCh := range []<-chan Result{first, second} {
		if result := <-resultCh; !errors.Is(result.Err, wantErr) {
			t.Fatalf("DoChan() error = %v, want %v", result.Err, wantErr)
		}
	}
}

func TestForgetKeepsNewCallRegistered(t *testing.T) {
	var group Group
	firstStarted := make(chan struct{})
	releaseFirst := make(chan struct{})
	first := group.DoChan("key", func() (any, error) {
		close(firstStarted)
		<-releaseFirst
		return "old", nil
	})
	<-firstStarted

	group.Forget("key")
	releaseSecond := make(chan struct{})
	second := group.DoChan("key", func() (any, error) {
		<-releaseSecond
		return "new", nil
	})

	close(releaseFirst)
	if result := <-first; result.Value != "old" {
		t.Fatalf("first result = %+v", result)
	}

	// 如果旧 call 错误删除了新 call，third 会执行自己的函数而不是等待 second。
	third := group.DoChan("key", func() (any, error) {
		t.Fatal("third function should join the second call")
		return nil, nil
	})
	close(releaseSecond)
	if result := <-second; result.Value != "new" {
		t.Fatalf("second result = %+v", result)
	}
	if result := <-third; result.Value != "new" {
		t.Fatalf("third result = %+v", result)
	}
}

func TestPanicBecomesPublishedError(t *testing.T) {
	var group Group
	result := <-group.DoChan("key", func() (any, error) {
		panic("boom")
	})

	var panicErr *PanicError
	if !errors.As(result.Err, &panicErr) {
		t.Fatalf("DoChan() error = %T %v, want *PanicError", result.Err, result.Err)
	}
	if panicErr.Value != "boom" || len(panicErr.Stack) == 0 {
		t.Fatalf("PanicError = %+v", panicErr)
	}
}

func TestGoexitBecomesPublishedError(t *testing.T) {
	var group Group
	resultCh := group.DoChan("key", func() (any, error) {
		runtime.Goexit()
		return nil, nil
	})

	select {
	case result := <-resultCh:
		if !errors.Is(result.Err, ErrGoexit) {
			t.Fatalf("DoChan() error = %v, want ErrGoexit", result.Err)
		}
	case <-time.After(time.Second):
		t.Fatal("DoChan() did not publish the Goexit result")
	}
}

func TestNilFunctionReturnsError(t *testing.T) {
	var group Group
	if _, err, _ := group.Do("key", nil); err == nil {
		t.Fatal("Do() should reject a nil function")
	}
	if result := <-group.DoChan("key", nil); result.Err == nil {
		t.Fatal("DoChan() should reject a nil function")
	}
}
