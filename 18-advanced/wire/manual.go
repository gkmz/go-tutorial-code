package wireexample

// InitializeAppManually 使用普通 Go 代码装配依赖，并返回聚合清理函数。
func InitializeAppManually(config Config) (*App, func(), error) {
	repository, cleanup, err := NewMemoryRepository(config)
	if err != nil {
		return nil, nil, err
	}
	service := NewService(repository, config)
	app := NewApp(service)
	return app, cleanup, nil
}
