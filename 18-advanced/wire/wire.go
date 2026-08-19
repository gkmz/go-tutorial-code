//go:build wireinject

package wireexample

import "github.com/google/wire"

// InitializeApp 声明 App 的 Wire 依赖图。
func InitializeApp(config Config) (*App, func(), error) {
	wire.Build(
		NewMemoryRepository,
		wire.Bind(new(Repository), new(*memoryRepository)),
		NewService,
		NewApp,
	)
	return nil, nil, nil
}
