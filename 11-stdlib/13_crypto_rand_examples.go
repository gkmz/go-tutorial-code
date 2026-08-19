package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

// secureToken 生成包含 byteCount 个安全随机字节的 URL 安全令牌。
func secureToken(byteCount int) (string, error) {
	if byteCount <= 0 {
		return "", errors.New("随机字节数必须大于零")
	}

	buffer := make([]byte, byteCount)
	if _, err := rand.Read(buffer); err != nil {
		return "", fmt.Errorf("生成安全随机数: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(buffer), nil
}

func main() {
	token, err := secureToken(32)
	if err != nil {
		fmt.Println("生成令牌失败:", err)
		return
	}
	fmt.Println("URL-safe token:", token)
}
