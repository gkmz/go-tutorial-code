package race

import "sync"

// Registry 使用读写锁保护服务地址映射。
type Registry struct {
	mu       sync.RWMutex
	services map[string]string
}

// NewRegistry 创建一个空的服务注册表。
func NewRegistry() *Registry {
	return &Registry{services: make(map[string]string)}
}

// Register 写入或更新服务地址。
func (r *Registry) Register(name, address string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.services[name] = address
}

// Lookup 查询服务地址。
func (r *Registry) Lookup(name string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	address, ok := r.services[name]
	return address, ok
}
