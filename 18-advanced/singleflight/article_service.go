package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/sync/singleflight"
)

// ArticleQuery 包含所有会影响文章查询结果的维度。
type ArticleQuery struct {
	Tenant string
	Locale string
	ID     int
}

// Article 是示例服务返回的不可变文章快照。
type Article struct {
	ID    int
	Title string
}

// ArticleLoader 表示真正访问数据库或下游服务的加载函数。
type ArticleLoader func(context.Context, ArticleQuery) (Article, error)

// ArticleService 使用本地缓存和 singleflight 合并同 key 的重叠查询。
type ArticleService struct {
	group       singleflight.Group
	cache       *ttlCache[ArticleQuery, Article]
	loader      ArticleLoader
	cacheTTL    time.Duration
	loadTimeout time.Duration
}

// NewArticleService 创建文章服务并校验缓存和共享加载的超时配置。
func NewArticleService(loader ArticleLoader, cacheTTL, loadTimeout time.Duration) (*ArticleService, error) {
	if loader == nil {
		return nil, errors.New("article loader must not be nil")
	}
	if cacheTTL <= 0 {
		return nil, errors.New("cache TTL must be greater than zero")
	}
	if loadTimeout <= 0 {
		return nil, errors.New("load timeout must be greater than zero")
	}
	return &ArticleService{
		cache:       newTTLCache[ArticleQuery, Article](),
		loader:      loader,
		cacheTTL:    cacheTTL,
		loadTimeout: loadTimeout,
	}, nil
}

// Find 返回文章快照。每个等待者可以独立取消，但共享加载使用单独的超时预算。
func (s *ArticleService) Find(ctx context.Context, query ArticleQuery) (Article, error) {
	if ctx == nil {
		return Article{}, errors.New("context must not be nil")
	}
	if err := query.validate(); err != nil {
		return Article{}, err
	}
	if article, ok := s.cache.get(query); ok {
		return article, nil
	}

	resultCh := s.group.DoChan(query.key(), func() (any, error) {
		// 等待首个加载期间，其他调用可能已经填充缓存，所以进入共享函数后再查一次。
		if article, ok := s.cache.get(query); ok {
			return article, nil
		}

		// 保留首个请求的值，但不让它的短 Deadline 直接取消所有等待者共享的加载。
		loadCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), s.loadTimeout)
		defer cancel()

		article, err := s.loader(loadCtx, query)
		if err != nil {
			return Article{}, err
		}
		s.cache.set(query, article, s.cacheTTL)
		return article, nil
	})

	select {
	case <-ctx.Done():
		return Article{}, ctx.Err()
	case result := <-resultCh:
		if result.Err != nil {
			return Article{}, result.Err
		}
		article, ok := result.Val.(Article)
		if !ok {
			return Article{}, fmt.Errorf("unexpected article result type %T", result.Val)
		}
		return article, nil
	}
}

func (q ArticleQuery) validate() error {
	if q.Tenant == "" {
		return errors.New("tenant must not be empty")
	}
	if q.Locale == "" {
		return errors.New("locale must not be empty")
	}
	if q.ID <= 0 {
		return errors.New("article ID must be greater than zero")
	}
	return nil
}

func (q ArticleQuery) key() string {
	values := url.Values{}
	values.Set("article", strconv.Itoa(q.ID))
	values.Set("locale", q.Locale)
	values.Set("tenant", q.Tenant)
	return values.Encode()
}
