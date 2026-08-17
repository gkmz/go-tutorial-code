// Package httpclient 提供带有限重试能力的 HTTP Client 参考实现。
package httpclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client 封装可复用的 http.Client 和重试次数。
type Client struct {
	httpClient *http.Client
	maxRetry   int
}

// New 创建一个 HTTP Client；调用方应长期复用返回值。
func New(httpClient *http.Client, maxRetry int) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 5 * time.Second}
	}
	if maxRetry < 0 {
		maxRetry = 0
	}
	return &Client{httpClient: httpClient, maxRetry: maxRetry}
}

// Do 执行请求并返回响应。
// 默认只重试 GET、HEAD 和 OPTIONS；返回非空响应时由调用方负责关闭 Body。
func (c *Client) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	if req == nil {
		return nil, fmt.Errorf("nil request")
	}
	retryableMethod := isRetryableMethod(req.Method)
	var lastErr error
	for attempt := 0; attempt <= c.maxRetry; attempt++ {
		current, err := cloneRequest(ctx, req, attempt)
		if err != nil {
			return nil, err
		}
		resp, requestErr := c.doOnce(current, retryableMethod)
		if resp != nil {
			return resp, nil
		}
		lastErr = requestErr
		if !retryableMethod || attempt == c.maxRetry {
			break
		}
		if err := waitBackoff(ctx, attempt); err != nil {
			return nil, err
		}
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	return nil, lastErr
}

func isRetryableMethod(method string) bool {
	return method == http.MethodGet || method == http.MethodHead || method == http.MethodOptions
}

func cloneRequest(ctx context.Context, req *http.Request, attempt int) (*http.Request, error) {
	current := req.Clone(ctx)
	if attempt == 0 || req.Body == nil {
		return current, nil
	}
	if req.GetBody == nil {
		return nil, fmt.Errorf("request body is not replayable")
	}
	body, err := req.GetBody()
	if err != nil {
		return nil, fmt.Errorf("recreate request body: %w", err)
	}
	current.Body = body
	return current, nil
}

func (c *Client) doOnce(req *http.Request, retryableMethod bool) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < http.StatusInternalServerError || !retryableMethod {
		return resp, nil
	}
	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
	return nil, fmt.Errorf("unexpected status: %s", resp.Status)
}

func waitBackoff(ctx context.Context, attempt int) error {
	timer := time.NewTimer(time.Duration(attempt+1) * 10 * time.Millisecond)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
