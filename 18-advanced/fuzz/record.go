package fuzz

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ErrInputTooLarge 表示解析输入超过调用方允许的大小。
var ErrInputTooLarge = errors.New("input exceeds size limit")

// Record 是 JSON 解析与往返属性示例使用的数据结构。
type Record struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// Validate 检查 Record 是否满足示例业务约束。
func (r Record) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name must not be empty")
	}
	if r.Count < 0 {
		return errors.New("count must not be negative")
	}
	return nil
}

// ParseRecord 在大小限制内解析并校验 JSON 记录。
func ParseRecord(input []byte, maxBytes int) (Record, error) {
	if maxBytes <= 0 {
		return Record{}, errors.New("maxBytes must be positive")
	}
	if len(input) > maxBytes {
		return Record{}, ErrInputTooLarge
	}
	var record Record
	if err := json.Unmarshal(input, &record); err != nil {
		return Record{}, fmt.Errorf("decode record: %w", err)
	}
	if err := record.Validate(); err != nil {
		return Record{}, fmt.Errorf("validate record: %w", err)
	}
	return record, nil
}

// EncodeRecord 校验并编码 JSON 记录。
func EncodeRecord(record Record) ([]byte, error) {
	if err := record.Validate(); err != nil {
		return nil, fmt.Errorf("validate record: %w", err)
	}
	return json.Marshal(record)
}
