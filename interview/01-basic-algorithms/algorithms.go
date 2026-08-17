// Package algorithms 提供 Go 面试中常见的基础算法参考实现。
package algorithms

import "sort"

// TwoSum 返回数组中和为 target 的两个元素下标。
// 如果不存在满足条件的两个元素，返回 nil。
func TwoSum(values []int, target int) []int {
	seen := make(map[int]int, len(values))
	for i, value := range values {
		if j, ok := seen[target-value]; ok {
			return []int{j, i}
		}
		seen[value] = i
	}
	return nil
}

// Interval 表示一个闭区间 [Start, End]。
type Interval struct {
	Start int
	End   int
}

// MergeIntervals 合并重叠或首尾相接的区间。
// 输入不会被修改，返回结果按 Start 升序排列。
func MergeIntervals(input []Interval) []Interval {
	if len(input) == 0 {
		return nil
	}
	intervals := append([]Interval(nil), input...)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start == intervals[j].Start {
			return intervals[i].End < intervals[j].End
		}
		return intervals[i].Start < intervals[j].Start
	})

	merged := make([]Interval, 0, len(intervals))
	for _, current := range intervals {
		if current.End < current.Start {
			current.Start, current.End = current.End, current.Start
		}
		if len(merged) == 0 || current.Start > merged[len(merged)-1].End {
			merged = append(merged, current)
			continue
		}
		if current.End > merged[len(merged)-1].End {
			merged[len(merged)-1].End = current.End
		}
	}
	return merged
}

// LongestUniqueSubstring 返回字符串中不含重复字符的最长子串长度。
// 题目按 Unicode Rune 计数，而不是按 UTF-8 字节计数。
func LongestUniqueSubstring(input string) int {
	runes := []rune(input)
	last := make(map[rune]int, len(runes))
	left, longest := 0, 0
	for right, r := range runes {
		if previous, ok := last[r]; ok && previous >= left {
			left = previous + 1
		}
		last[r] = right
		if width := right - left + 1; width > longest {
			longest = width
		}
	}
	return longest
}

// TreeNode 是二叉树节点。
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrder 按层序遍历二叉树，并返回每一层的节点值。
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Value)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
