package exercises

import "sync"

// CleanupLog 以并发安全方式记录资源清理顺序。
type CleanupLog struct {
	mu      sync.Mutex
	entries []string
}

// NewCleanupLog 创建清理顺序记录器。
func NewCleanupLog() *CleanupLog {
	return &CleanupLog{}
}

// Add 追加一条清理记录。
func (l *CleanupLog) Add(entry string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.entries = append(l.entries, entry)
}

// Entries 返回清理记录的副本，避免调用方修改内部切片。
func (l *CleanupLog) Entries() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return append([]string(nil), l.entries...)
}

// ResourceA 表示先创建的底层资源。
type ResourceA struct{}

// NewResourceA 创建底层资源及其清理函数。
func NewResourceA(log *CleanupLog) (*ResourceA, func()) {
	return &ResourceA{}, func() {
		log.Add("resource-a")
	}
}

// ResourceB 表示依赖 ResourceA 的上层资源。
type ResourceB struct {
	resourceA *ResourceA
}

// NewResourceB 创建上层资源及其清理函数。
func NewResourceB(resourceA *ResourceA, log *CleanupLog) (*ResourceB, func()) {
	return &ResourceB{resourceA: resourceA}, func() {
		log.Add("resource-b")
	}
}

// ExerciseApp 表示完成装配后的练习应用。
type ExerciseApp struct {
	resourceB *ResourceB
}

// NewExerciseApp 创建练习应用。
func NewExerciseApp(resourceB *ResourceB) *ExerciseApp {
	return &ExerciseApp{resourceB: resourceB}
}

// Ready 报告依赖图是否已经完整装配。
func (a *ExerciseApp) Ready() bool {
	return a != nil && a.resourceB != nil && a.resourceB.resourceA != nil
}
