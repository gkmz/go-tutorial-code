package main

import (
	"context"
	"fmt"
	"time"
)

// fetch 模拟一个可取消的下游调用。
func fetch(ctx context.Context) error {
	select {
	case <-time.After(50 * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	fmt.Println(fetch(ctx))
}
