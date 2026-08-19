package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// 完整示例：综合使用标准库

// User 表示可以持久化为 JSON 的用户。
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// String 返回不包含敏感字段的用户摘要。
func (u User) String() string {
	return fmt.Sprintf("User{ID: %d, Name: %s, Email: %s, Created: %s}",
		u.ID, u.Name, u.Email, u.CreatedAt.Format("2006-01-02"))
}

// 保存用户到文件
func saveUser(user User, filename string) error {
	data, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("write file error: %w", err)
	}

	return nil
}

// 从文件加载用户
func loadUser(filename string) (*User, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read file error: %w", err)
	}

	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return &user, nil
}

// processName 清理用户名中的首尾空白和连续空白。
func processName(name string) string {
	return strings.Join(strings.Fields(name), " ")
}

func main() {
	// 创建用户
	user := User{
		ID:        1,
		Name:      processName("  alice smith  "),
		Email:     "alice@example.com",
		CreatedAt: time.Now(),
	}

	// 打印用户信息
	fmt.Println(user)

	// 保存到文件
	filename := "user.json"
	err := saveUser(user, filename)
	if err != nil {
		fmt.Println("Save error:", err)
		return
	}
	fmt.Println("User saved to", filename)

	// 从文件加载
	loadedUser, err := loadUser(filename)
	if err != nil {
		fmt.Println("Load error:", err)
		return
	}
	fmt.Println("Loaded user:", loadedUser)

	// 计算账户年龄
	age := time.Since(loadedUser.CreatedAt)
	fmt.Printf("Account age: %v\n", age.Round(time.Second))

	// 清理示例生成的文件，并报告清理失败。
	if err := os.Remove(filename); err != nil {
		fmt.Println("Cleanup error:", err)
		return
	}
	fmt.Println("Cleaned up")
}
