// Package wireexample 演示手工依赖装配与 Wire 生成的编译期依赖注入。
package wireexample

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

// ErrRepositoryClosed 表示仓库资源已经被清理函数关闭。
var ErrRepositoryClosed = errors.New("repository is closed")

// Config 描述示例应用的构造参数。
type Config struct {
	DSN            string
	GreetingPrefix string
}

// Repository 描述 Service 所需的最小持久化能力。
type Repository interface {
	SaveGreeting(context.Context, string) error
	Count() int
}

type memoryRepository struct {
	mu     sync.RWMutex
	closed bool
	items  []string
}

// NewMemoryRepository 创建仓库，并返回用于关闭资源的清理函数。
func NewMemoryRepository(config Config) (*memoryRepository, func(), error) {
	if strings.TrimSpace(config.DSN) == "" {
		return nil, nil, errors.New("repository DSN must not be empty")
	}
	repository := &memoryRepository{}
	cleanup := func() {
		repository.mu.Lock()
		defer repository.mu.Unlock()
		repository.closed = true
	}
	return repository, cleanup, nil
}

func (r *memoryRepository) SaveGreeting(ctx context.Context, message string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.closed {
		return ErrRepositoryClosed
	}
	r.items = append(r.items, message)
	return nil
}

func (r *memoryRepository) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.items)
}

// Service 负责生成问候消息并通过 Repository 保存结果。
type Service struct {
	repository Repository
	prefix     string
}

// NewService 创建业务服务。
func NewService(repository Repository, config Config) *Service {
	prefix := strings.TrimSpace(config.GreetingPrefix)
	if prefix == "" {
		prefix = "Hello"
	}
	return &Service{repository: repository, prefix: prefix}
}

// Greet 生成并保存问候消息。
func (s *Service) Greet(ctx context.Context, name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", errors.New("name must not be empty")
	}
	message := fmt.Sprintf("%s, %s!", s.prefix, name)
	if err := s.repository.SaveGreeting(ctx, message); err != nil {
		return "", fmt.Errorf("save greeting: %w", err)
	}
	return message, nil
}

// App 是完成依赖装配后的应用入口。
type App struct {
	service *Service
}

// NewApp 创建应用入口。
func NewApp(service *Service) *App {
	return &App{service: service}
}

// Greet 调用业务服务生成问候消息。
func (a *App) Greet(ctx context.Context, name string) (string, error) {
	return a.service.Greet(ctx, name)
}

// GreetingCount 返回当前已经保存的问候数量。
func (a *App) GreetingCount() int {
	return a.service.repository.Count()
}
