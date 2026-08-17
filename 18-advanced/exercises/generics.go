// Package exercises 提供高级篇代码类练习的参考实现。
package exercises

// Reverse 返回输入切片的反转副本，不修改原切片。
func Reverse[T any](values []T) []T {
	result := make([]T, len(values))
	for i := range values {
		result[len(values)-1-i] = values[i]
	}
	return result
}

// Queue 是一个使用切片保存元素的泛型先进先出队列。
type Queue[T any] struct {
	values []T
}

// Enqueue 将元素加入队尾。
func (q *Queue[T]) Enqueue(value T) {
	q.values = append(q.values, value)
}

// Dequeue 移除并返回队首元素；队列为空时返回 false。
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.values) == 0 {
		var zero T
		return zero, false
	}
	value := q.values[0]
	var zero T
	q.values[0] = zero
	q.values = q.values[1:]
	return value, true
}

// Reduce 从 initial 开始聚合切片中的元素。
func Reduce[T any, U any](values []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, value := range values {
		result = fn(result, value)
	}
	return result
}

// Set 保存可比较类型的唯一值。
type Set[T comparable] struct {
	values map[T]struct{}
}

// Add 插入一个值。
func (s *Set[T]) Add(value T) {
	if s.values == nil {
		s.values = make(map[T]struct{})
	}
	s.values[value] = struct{}{}
}

// Remove 删除一个值。
func (s *Set[T]) Remove(value T) {
	delete(s.values, value)
}

// Contains 判断集合中是否存在指定值。
func (s *Set[T]) Contains(value T) bool {
	_, ok := s.values[value]
	return ok
}

// GroupBy 根据 keyFn 将元素分组。
func GroupBy[T any, K comparable](values []T, keyFn func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, value := range values {
		key := keyFn(value)
		result[key] = append(result[key], value)
	}
	return result
}
