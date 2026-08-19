// Package exercises 提供 Wire 章节代码练习的参考答案。
package exercises

import (
	"context"
	"sync"
	"testing"

	wireexample "github.com/hankmor/go-tutorial-code/18-advanced/wire"
)

type fakeRepository struct {
	mu       sync.Mutex
	messages []string
}

func (r *fakeRepository) SaveGreeting(ctx context.Context, message string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.messages = append(r.messages, message)
	return nil
}

func (r *fakeRepository) Count() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.messages)
}

func TestManualInjectionWithFakeRepository(t *testing.T) {
	// 测试直接装配业务组件，不需要启动真实仓库或运行 Wire 生成器。
	repository := &fakeRepository{}
	service := wireexample.NewService(repository, wireexample.Config{GreetingPrefix: "Test"})
	app := wireexample.NewApp(service)

	message, err := app.Greet(context.Background(), "Ada")
	if err != nil {
		t.Fatal(err)
	}
	if want := "Test, Ada!"; message != want {
		t.Fatalf("Greet() = %q, want %q", message, want)
	}
	if got := repository.Count(); got != 1 {
		t.Fatalf("Count() = %d, want 1", got)
	}
}
