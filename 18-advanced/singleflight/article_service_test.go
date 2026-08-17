package main

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

func TestGroupMergesOverlappingCalls(t *testing.T) {
	const callers = 20
	var group singleflight.Group
	var loads atomic.Int32
	started := make(chan struct{})
	release := make(chan struct{})
	resultChannels := make([]<-chan singleflight.Result, 0, callers)

	for range callers {
		resultChannels = append(resultChannels, group.DoChan("article:1", func() (any, error) {
			if loads.Add(1) == 1 {
				close(started)
			}
			<-release
			return "article-1", nil
		}))
	}
	<-started
	close(release)

	for _, resultCh := range resultChannels {
		result := <-resultCh
		if result.Err != nil {
			t.Fatalf("DoChan() error = %v", result.Err)
		}
		if result.Val != "article-1" {
			t.Fatalf("DoChan() value = %v, want article-1", result.Val)
		}
		if !result.Shared {
			t.Fatal("DoChan() Shared = false, want true")
		}
	}
	if got := loads.Load(); got != 1 {
		t.Fatalf("loader calls = %d, want 1", got)
	}
}

func TestArticleServiceCachesSuccessfulLoad(t *testing.T) {
	var loads atomic.Int32
	service := newTestArticleService(t, func(context.Context, ArticleQuery) (Article, error) {
		loads.Add(1)
		return Article{ID: 1, Title: "cached"}, nil
	})
	query := testQuery(1)

	for range 2 {
		article, err := service.Find(context.Background(), query)
		if err != nil {
			t.Fatalf("Find() error = %v", err)
		}
		if article.Title != "cached" {
			t.Fatalf("Find() title = %q, want cached", article.Title)
		}
	}
	if got := loads.Load(); got != 1 {
		t.Fatalf("loader calls = %d, want 1", got)
	}
}

func TestArticleServiceDoesNotMergeDifferentKeys(t *testing.T) {
	var loads atomic.Int32
	service := newTestArticleService(t, func(_ context.Context, query ArticleQuery) (Article, error) {
		loads.Add(1)
		return Article{ID: query.ID, Title: query.Tenant + ":" + query.Locale}, nil
	})

	queries := []ArticleQuery{
		{Tenant: "tenant-a", Locale: "zh-CN", ID: 1},
		{Tenant: "tenant-b", Locale: "zh-CN", ID: 1},
		{Tenant: "tenant-a", Locale: "en-US", ID: 1},
	}
	for _, query := range queries {
		if _, err := service.Find(context.Background(), query); err != nil {
			t.Fatalf("Find(%+v) error = %v", query, err)
		}
	}
	if got := loads.Load(); got != int32(len(queries)) {
		t.Fatalf("loader calls = %d, want %d", got, len(queries))
	}
}

func TestWaiterTimeoutDoesNotCancelSharedLoad(t *testing.T) {
	loaderStarted := make(chan struct{})
	releaseLoader := make(chan struct{})
	service := newTestArticleService(t, func(ctx context.Context, query ArticleQuery) (Article, error) {
		close(loaderStarted)
		select {
		case <-releaseLoader:
			return Article{ID: query.ID, Title: "loaded"}, nil
		case <-ctx.Done():
			return Article{}, ctx.Err()
		}
	})
	query := testQuery(1)

	shortCtx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	shortResult := make(chan error, 1)
	go func() {
		_, err := service.Find(shortCtx, query)
		shortResult <- err
	}()
	<-loaderStarted

	longResult := make(chan error, 1)
	go func() {
		article, err := service.Find(context.Background(), query)
		if err == nil && article.Title != "loaded" {
			err = errors.New("unexpected article result")
		}
		longResult <- err
	}()

	if err := <-shortResult; !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("short waiter error = %v, want context deadline", err)
	}
	close(releaseLoader)
	if err := <-longResult; err != nil {
		t.Fatalf("long waiter error = %v", err)
	}

	// 共享加载完成后，后续请求应直接命中缓存。
	if _, err := service.Find(context.Background(), query); err != nil {
		t.Fatalf("cached Find() error = %v", err)
	}
}

func TestArticleServiceSharesErrorsButDoesNotCacheThem(t *testing.T) {
	downstreamErr := errors.New("database unavailable")
	var loads atomic.Int32
	service := newTestArticleService(t, func(context.Context, ArticleQuery) (Article, error) {
		loads.Add(1)
		return Article{}, downstreamErr
	})
	query := testQuery(1)

	for range 2 {
		_, err := service.Find(context.Background(), query)
		if !errors.Is(err, downstreamErr) {
			t.Fatalf("Find() error = %v, want downstream error", err)
		}
	}
	if got := loads.Load(); got != 2 {
		t.Fatalf("loader calls = %d, want 2 because errors are not cached", got)
	}
}

func TestForgetAllowsNewCall(t *testing.T) {
	var group singleflight.Group
	firstStarted := make(chan struct{})
	releaseFirst := make(chan struct{})
	firstDone := make(chan struct{})

	go func() {
		defer close(firstDone)
		_, _, _ = group.Do("key", func() (any, error) {
			close(firstStarted)
			<-releaseFirst
			return "old", nil
		})
	}()
	<-firstStarted
	group.Forget("key")

	second := group.DoChan("key", func() (any, error) {
		return "new", nil
	})
	if result := <-second; result.Val != "new" || result.Err != nil {
		t.Fatalf("second call result = %+v, want new", result)
	}
	close(releaseFirst)
	<-firstDone
}

func TestArticleQueryKeyIncludesResultDimensions(t *testing.T) {
	base := ArticleQuery{Tenant: "tenant-a", Locale: "zh-CN", ID: 1}
	variants := []ArticleQuery{
		{Tenant: "tenant-b", Locale: base.Locale, ID: base.ID},
		{Tenant: base.Tenant, Locale: "en-US", ID: base.ID},
		{Tenant: base.Tenant, Locale: base.Locale, ID: 2},
	}
	for _, variant := range variants {
		if base.key() == variant.key() {
			t.Fatalf("key collision: %+v and %+v", base, variant)
		}
	}
}

func newTestArticleService(t *testing.T, loader ArticleLoader) *ArticleService {
	t.Helper()
	service, err := NewArticleService(loader, time.Minute, time.Second)
	if err != nil {
		t.Fatalf("NewArticleService() error = %v", err)
	}
	return service
}

func testQuery(id int) ArticleQuery {
	return ArticleQuery{Tenant: "tenant-a", Locale: "zh-CN", ID: id}
}

func TestNewArticleServiceValidatesConfiguration(t *testing.T) {
	loader := func(context.Context, ArticleQuery) (Article, error) { return Article{}, nil }
	tests := []struct {
		name        string
		loader      ArticleLoader
		cacheTTL    time.Duration
		loadTimeout time.Duration
	}{
		{name: "nil loader", cacheTTL: time.Second, loadTimeout: time.Second},
		{name: "invalid cache TTL", loader: loader, loadTimeout: time.Second},
		{name: "invalid load timeout", loader: loader, cacheTTL: time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewArticleService(tt.loader, tt.cacheTTL, tt.loadTimeout); err == nil {
				t.Fatal("NewArticleService() should reject invalid configuration")
			}
		})
	}
}

func TestFindValidatesInput(t *testing.T) {
	service := newTestArticleService(t, func(context.Context, ArticleQuery) (Article, error) {
		return Article{}, nil
	})
	tests := []struct {
		name  string
		ctx   context.Context
		query ArticleQuery
	}{
		{name: "nil context", query: testQuery(1)},
		{name: "empty tenant", ctx: context.Background(), query: ArticleQuery{Locale: "zh-CN", ID: 1}},
		{name: "empty locale", ctx: context.Background(), query: ArticleQuery{Tenant: "tenant-a", ID: 1}},
		{name: "invalid ID", ctx: context.Background(), query: ArticleQuery{Tenant: "tenant-a", Locale: "zh-CN"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := service.Find(tt.ctx, tt.query); err == nil {
				t.Fatal("Find() should reject invalid input")
			}
		})
	}
}

func TestArticleServiceConcurrentCacheMiss(t *testing.T) {
	var loads atomic.Int32
	loaderStarted := make(chan struct{})
	releaseLoader := make(chan struct{})
	service := newTestArticleService(t, func(_ context.Context, query ArticleQuery) (Article, error) {
		if loads.Add(1) == 1 {
			close(loaderStarted)
		}
		<-releaseLoader
		return Article{ID: query.ID, Title: "shared"}, nil
	})

	const callers = 10
	start := make(chan struct{})
	var done sync.WaitGroup
	errorsCh := make(chan error, callers)
	for range callers {
		done.Go(func() {
			<-start
			_, err := service.Find(context.Background(), testQuery(1))
			errorsCh <- err
		})
	}
	close(start)
	<-loaderStarted
	time.Sleep(10 * time.Millisecond)
	close(releaseLoader)
	done.Wait()
	close(errorsCh)

	for err := range errorsCh {
		if err != nil {
			t.Fatalf("Find() error = %v", err)
		}
	}
	if got := loads.Load(); got != 1 {
		t.Fatalf("loader calls = %d, want 1", got)
	}
}
