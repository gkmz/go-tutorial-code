// Package shadow 提供变量遮蔽练习的参考实现。
package shadow

import "fmt"

// LoadCombined 依次加载基础值和附加值，并保留每一层错误。
func LoadCombined(load, loadExtra func() (int, error)) (int, error) {
	value, err := load()
	if err != nil {
		return 0, fmt.Errorf("load base value: %w", err)
	}

	extra, err := loadExtra()
	if err != nil {
		return 0, fmt.Errorf("load extra value: %w", err)
	}
	return value + extra, nil
}
