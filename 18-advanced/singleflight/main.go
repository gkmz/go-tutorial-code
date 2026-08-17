package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var loads atomic.Int32
	service, err := NewArticleService(
		func(ctx context.Context, query ArticleQuery) (Article, error) {
			loads.Add(1)
			select {
			case <-time.After(80 * time.Millisecond):
				return Article{ID: query.ID, Title: "singleflight 实战"}, nil
			case <-ctx.Done():
				return Article{}, ctx.Err()
			}
		},
		time.Minute,
		500*time.Millisecond,
	)
	if err != nil {
		panic(err)
	}

	query := ArticleQuery{Tenant: "tenant-a", Locale: "zh-CN", ID: 1}
	results := make(chan Article, 10)
	var callers sync.WaitGroup
	for range 10 {
		callers.Go(func() {
			article, findErr := service.Find(context.Background(), query)
			if findErr != nil {
				fmt.Println("find article:", findErr)
				return
			}
			results <- article
		})
	}
	callers.Wait()
	close(results)

	for article := range results {
		fmt.Printf("article=%d title=%q\n", article.ID, article.Title)
	}
	fmt.Println("actual loader calls:", loads.Load())
}
