// Package initialization 对比包级隐式初始化和显式构造函数。
package initialization

import "errors"

var implicitRegistry = map[string]string{}

func init() {
	// init 会在导入包时执行，调用方无法传入配置或接收初始化错误。
	implicitRegistry["default"] = "implicit"
}

// Service 保存由调用方显式提供的配置。
type Service struct {
	name string
}

// NewService 校验配置并显式创建 Service。
func NewService(name string) (*Service, error) {
	if name == "" {
		return nil, errors.New("service name must not be empty")
	}
	return &Service{name: name}, nil
}

// Name 返回 Service 的配置名称。
func (s *Service) Name() string {
	return s.name
}

// ImplicitValue 返回由 init 注册的值，仅用于对照隐式副作用。
func ImplicitValue() string {
	return implicitRegistry["default"]
}
