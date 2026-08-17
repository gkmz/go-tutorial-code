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

// SearchTree 是一个由调用方定义排序规则的泛型二叉搜索树。
type SearchTree[T any] struct {
	compare func(T, T) int
	root    *searchNode[T]
}

// searchNode 保存二叉搜索树的一个节点。
type searchNode[T any] struct {
	value T
	left  *searchNode[T]
	right *searchNode[T]
}

// NewSearchTree 使用 compare 创建一棵空的二叉搜索树。
// compare 返回负数、零或正数，分别表示第一个值小于、等于或大于第二个值。
func NewSearchTree[T any](compare func(T, T) int) *SearchTree[T] {
	if compare == nil {
		panic("compare function must not be nil")
	}
	return &SearchTree[T]{compare: compare}
}

// Insert 插入 value；如果 value 已存在，则保持树结构不变。
func (tree *SearchTree[T]) Insert(value T) {
	tree.root = insertNode(tree.root, value, tree.compare)
}

func insertNode[T any](node *searchNode[T], value T, compare func(T, T) int) *searchNode[T] {
	if node == nil {
		return &searchNode[T]{value: value}
	}
	comparison := compare(value, node.value)
	if comparison < 0 {
		node.left = insertNode(node.left, value, compare)
	} else if comparison > 0 {
		node.right = insertNode(node.right, value, compare)
	}
	return node
}

// Contains 判断 value 是否存在于树中。
func (tree *SearchTree[T]) Contains(value T) bool {
	for node := tree.root; node != nil; {
		comparison := tree.compare(value, node.value)
		if comparison == 0 {
			return true
		}
		if comparison < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}
	return false
}

// Delete 删除 value；如果 value 不存在，则保持树结构不变。
func (tree *SearchTree[T]) Delete(value T) {
	tree.root = deleteNode(tree.root, value, tree.compare)
}

func deleteNode[T any](node *searchNode[T], value T, compare func(T, T) int) *searchNode[T] {
	if node == nil {
		return nil
	}
	comparison := compare(value, node.value)
	if comparison < 0 {
		node.left = deleteNode(node.left, value, compare)
		return node
	}
	if comparison > 0 {
		node.right = deleteNode(node.right, value, compare)
		return node
	}

	if node.left == nil {
		return node.right
	}
	if node.right == nil {
		return node.left
	}

	// 两个子节点都存在时，用右子树中的最小节点替换当前值。
	successor := node.right
	for successor.left != nil {
		successor = successor.left
	}
	node.value = successor.value
	node.right = deleteNode(node.right, successor.value, compare)
	return node
}
