// Package exercises 提供标准库章节练习的参考实现。
package exercises

import (
	"encoding/json"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

// NormalizeText 将连续空白规范化为单个空格。
func NormalizeText(value string) string {
	return strings.Join(strings.Fields(value), " ")
}

// ParseNumber 解析十进制整数并返回错误。
func ParseNumber(value string) (int, error) {
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return number, nil
}

// Event 表示带 JSON 标签的事件。
type Event struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Secret    string    `json:"-"`
}

// EncodeEvent 编码事件并隐藏 Secret。
func EncodeEvent(event Event) ([]byte, error) { return json.Marshal(event) }

// StableSearch 稳定排序 values 并返回 target 的索引。
func StableSearch(values []int, target int) int {
	sort.Ints(values)
	index := sort.SearchInts(values, target)
	if index == len(values) || values[index] != target {
		return -1
	}
	return index
}

// RandomSequence 使用固定种子生成可重复的随机序列。
func RandomSequence(seed int64, count int) []int {
	random := rand.New(rand.NewSource(seed))
	result := make([]int, count)
	for i := range result {
		result[i] = random.Intn(100)
	}
	return result
}
