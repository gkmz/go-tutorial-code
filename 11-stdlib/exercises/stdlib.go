// Package exercises 提供标准库章节练习的参考实现。
package exercises

import (
	"bytes"
	"cmp"
	"context"
	cryptorand "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
)

// NormalizeText 将连续的 Unicode 空白规范化为单个空格。
func NormalizeText(value string) string {
	return strings.Join(strings.Fields(value), " ")
}

// ParsedNumbers 保存解析后的整数和浮点数。
type ParsedNumbers struct {
	Integer int64
	Decimal float64
}

// ParseNumbers 按十进制解析一个 64 位整数和一个 64 位浮点数。
func ParseNumbers(integerText, decimalText string) (ParsedNumbers, error) {
	integer, err := strconv.ParseInt(integerText, 10, 64)
	if err != nil {
		return ParsedNumbers{}, fmt.Errorf("解析整数 %q: %w", integerText, err)
	}

	decimal, err := strconv.ParseFloat(decimalText, 64)
	if err != nil {
		return ParsedNumbers{}, fmt.Errorf("解析浮点数 %q: %w", decimalText, err)
	}

	return ParsedNumbers{Integer: integer, Decimal: decimal}, nil
}

// ParseNumber 解析十进制 int，保留给最小转换练习使用。
func ParseNumber(value string) (int, error) {
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return number, nil
}

// ParseIntervalInLocation 在指定时区解析两个本地时间并计算间隔。
func ParseIntervalInLocation(layout, startText, endText string, location *time.Location) (time.Duration, error) {
	if location == nil {
		return 0, errors.New("时区不能为空")
	}

	start, err := time.ParseInLocation(layout, startText, location)
	if err != nil {
		return 0, fmt.Errorf("解析开始时间: %w", err)
	}
	end, err := time.ParseInLocation(layout, endText, location)
	if err != nil {
		return 0, fmt.Errorf("解析结束时间: %w", err)
	}
	return end.Sub(start), nil
}

// ErrWaitTimeout 表示等待结果超过了指定时限。
var ErrWaitTimeout = errors.New("等待结果超时")

// WaitForResult 等待结果、超时或 Context 取消，并返回最先发生的事件。
func WaitForResult[T any](ctx context.Context, result <-chan T, timeout time.Duration) (T, error) {
	var zero T
	if ctx == nil {
		return zero, errors.New("context 不能为空")
	}

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case value := <-result:
		return value, nil
	case <-timer.C:
		return zero, ErrWaitTimeout
	case <-ctx.Done():
		return zero, ctx.Err()
	}
}

// Event 表示带 JSON 标签的事件。
type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	Secret      string    `json:"-"`
}

// EncodeEvent 编码事件，并根据标签省略空描述和敏感字段。
func EncodeEvent(event Event) ([]byte, error) {
	return json.Marshal(event)
}

// DecodeEventStrict 只接受字段已知且仅包含一个 JSON 值的事件文档。
func DecodeEventStrict(data []byte) (Event, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()

	var event Event
	if err := decoder.Decode(&event); err != nil {
		return Event{}, fmt.Errorf("解码事件: %w", err)
	}

	// 第二次解码必须直接到达 EOF，避免静默接受拼接的多个 JSON 值。
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		if err == nil {
			return Event{}, errors.New("JSON 后存在额外值")
		}
		return Event{}, fmt.Errorf("检查 JSON 尾部: %w", err)
	}
	return event, nil
}

// Record 是稳定排序练习使用的记录。
type Record struct {
	Key      int
	Sequence int
}

// StableSortRecords 按 Key 稳定排序副本，不修改输入切片。
func StableSortRecords(records []Record) []Record {
	result := slices.Clone(records)
	slices.SortStableFunc(result, func(left, right Record) int {
		return cmp.Compare(left.Key, right.Key)
	})
	return result
}

// FirstRecordIndex 返回已按 Key 升序排列的记录中第一个目标键的位置。
func FirstRecordIndex(records []Record, target int) int {
	index, found := slices.BinarySearchFunc(records, target, func(record Record, key int) int {
		return cmp.Compare(record.Key, key)
	})
	if !found {
		return -1
	}
	return index
}

// StableSearch 原地排序 values 并返回 target 的第一个索引。
func StableSearch(values []int, target int) int {
	slices.Sort(values)
	index, found := slices.BinarySearch(values, target)
	if !found {
		return -1
	}
	return index
}

// SortedCopySearch 返回排序后的副本和 target 的第一个索引，不修改输入。
func SortedCopySearch(values []int, target int) ([]int, int) {
	sorted := slices.Clone(values)
	slices.Sort(sorted)
	index, found := slices.BinarySearch(sorted, target)
	if !found {
		return sorted, -1
	}
	return sorted, index
}

// RandomSequence 使用局部固定种子生成可重复的非安全随机序列。
func RandomSequence(seed int64, count int) []int {
	if count <= 0 {
		return []int{}
	}

	random := rand.New(rand.NewSource(seed))
	result := make([]int, count)
	for i := range result {
		result[i] = random.Intn(100)
	}
	return result
}

// TempFileRoundTrip 在指定目录写入并读回文件内容。
func TempFileRoundTrip(directory, name string, content []byte) ([]byte, error) {
	path := filepath.Join(directory, name)
	if err := os.WriteFile(path, content, 0o600); err != nil {
		return nil, fmt.Errorf("写入临时文件: %w", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取临时文件: %w", err)
	}
	return data, nil
}

// SecureToken 生成包含 byteCount 个安全随机字节的 URL 安全令牌。
func SecureToken(byteCount int) (string, error) {
	if byteCount <= 0 {
		return "", errors.New("随机字节数必须大于零")
	}

	buffer := make([]byte, byteCount)
	if _, err := cryptorand.Read(buffer); err != nil {
		return "", fmt.Errorf("生成安全随机数: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(buffer), nil
}
