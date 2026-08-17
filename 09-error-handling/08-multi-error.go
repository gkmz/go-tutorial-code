package main

import (
	"errors"
	"fmt"
	"strings"
)

// 实战建议 5：错误聚合

// MultiError 聚合多个错误
type MultiError struct {
	Errors []error
}

func (m *MultiError) Error() string {
	if len(m.Errors) == 0 {
		return ""
	}
	if len(m.Errors) == 1 {
		return m.Errors[0].Error()
	}

	var msgs []string
	for i, err := range m.Errors {
		msgs = append(msgs, fmt.Sprintf("%d. %s", i+1, err.Error()))
	}
	return fmt.Sprintf("multiple errors occurred:\n%s", strings.Join(msgs, "\n"))
}

func (m *MultiError) Add(err error) {
	if err != nil {
		m.Errors = append(m.Errors, err)
	}
}

func (m *MultiError) HasErrors() bool {
	return len(m.Errors) > 0
}

// User 结构体
type User struct {
	Name  string
	Age   int
	Email string
}

// 验证用户
func validateUser(user *User) error {
	var errs MultiError

	if user.Name == "" {
		errs.Add(errors.New("name is required"))
	}
	if len(user.Name) < 2 {
		errs.Add(errors.New("name must be at least 2 characters"))
	}
	if user.Age < 0 {
		errs.Add(errors.New("age must be positive"))
	}
	if user.Age > 150 {
		errs.Add(errors.New("age must be less than 150"))
	}
	if user.Email == "" {
		errs.Add(errors.New("email is required"))
	}
	if !strings.Contains(user.Email, "@") {
		errs.Add(errors.New("email must contain @"))
	}

	if errs.HasErrors() {
		return &errs
	}
	return nil
}

// 批量处理
func processUsers(users []*User) error {
	var errs MultiError

	for i, user := range users {
		if err := validateUser(user); err != nil {
			errs.Add(fmt.Errorf("user %d: %w", i, err))
		}
	}

	if errs.HasErrors() {
		return &errs
	}
	return nil
}

func main() {
	fmt.Println("=== 错误聚合示例 ===")

	// 示例 1：单个用户验证
	fmt.Println("\n--- 单个用户验证 ---")
	user1 := &User{
		Name:  "",
		Age:   -5,
		Email: "invalid",
	}
	if err := validateUser(user1); err != nil {
		fmt.Println("Validation errors:")
		fmt.Println(err)
	}

	// 示例 2：正常用户
	fmt.Println("\n--- 正常用户 ---")
	user2 := &User{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	if err := validateUser(user2); err != nil {
		fmt.Println("Validation errors:", err)
	} else {
		fmt.Println("User is valid")
	}

	// 示例 3：批量处理
	fmt.Println("\n--- 批量处理 ---")
	users := []*User{
		{Name: "Alice", Age: 30, Email: "alice@example.com"},
		{Name: "", Age: -5, Email: "invalid"},
		{Name: "Bob", Age: 25, Email: "bob@example.com"},
		{Name: "C", Age: 200, Email: ""},
	}

	if err := processUsers(users); err != nil {
		fmt.Println("Batch processing errors:")
		fmt.Println(err)
	} else {
		fmt.Println("All users are valid")
	}
}
