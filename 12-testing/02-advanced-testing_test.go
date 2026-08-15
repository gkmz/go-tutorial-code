package math

import (
	"fmt"
	"os"
	"testing"
)

// 1. Setup 与 Teardown (TestMain)
// TestMain 会在所有测试运行前后执行，适合初始化数据库、清理临时文件等全局操作。
func TestMain(m *testing.M) {
	fmt.Println(">>> [Setup] 全局资源初始化...")

	// 运行包下的所有测试
	exitCode := m.Run()

	fmt.Println(">>> [Teardown] 全局资源清理...")
	os.Exit(exitCode)
}

// 2. 并行测试 (t.Parallel)
// 并行测试可以显著缩短测试耗时，但要注意避免共享变量导致的竞态冲突。
func TestParallel(t *testing.T) {
	t.Run("SubTest1", func(t *testing.T) {
		t.Parallel() // 标记为可并行
		fmt.Println("Running SubTest1")
	})
	t.Run("SubTest2", func(t *testing.T) {
		t.Parallel()
		fmt.Println("Running SubTest2")
	})
}

// 3. 测试辅助函数 (t.Helper)
// 使用 t.Helper() 后，测试失败时报出的行号会显示调用方（即 Test 函数内）的行号，而不是辅助函数内部。
func assertEqual(t *testing.T, got, want int) {
	t.Helper() // 显式标记为 Helper
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestWithHelper(t *testing.T) {
	assertEqual(t, 1+1, 2)
}

// 4. 面向接口的 Mock 实战
// 被测试的业务逻辑依赖一个接口，而不是具体的数据库实现。
type UserRepository interface {
	GetName(id int) string
}

func GetGreeting(repo UserRepository, id int) string {
	name := repo.GetName(id)
	return "Hello, " + name
}

// 在测试中定义一个 Mock 实现
type MockRepo struct{}

func (m *MockRepo) GetName(id int) string {
	if id == 1 {
		return "极客老墨"
	}
	return "Guest"
}

func TestGetGreeting(t *testing.T) {
	mock := &MockRepo{}
	got := GetGreeting(mock, 1)
	want := "Hello, 极客老墨"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// 避坑引导：
// - 使用 -race 运行此文件：go test -v -race 02-advanced-testing_test.go
// - 并行测试时，注意循环变量的闭包坑（Go 1.22 已修复，但老版本仍需注意）。
// - 尽量使用外部测试包（package main_test）来强制进行黑盒测试。
