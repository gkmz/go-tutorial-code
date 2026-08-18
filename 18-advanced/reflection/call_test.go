package main

import (
	"errors"
	"testing"
)

type handler struct{}

// Handle 返回带前缀的输入，特殊输入用于测试错误传播。
func (handler) Handle(input string) (string, error) {
	if input == "fail" {
		return "", errors.New("requested failure")
	}
	return "handled:" + input, nil
}

// Wrong 故意提供不匹配的签名，用于验证调用前检查。
func (handler) Wrong(int) string { return "wrong" }

func TestCallStringMethod(t *testing.T) {
	got, err := CallStringMethod(handler{}, "Handle", "request")
	if err != nil || got != "handled:request" {
		t.Fatalf("CallStringMethod() = %q, %v", got, err)
	}
	if _, err := CallStringMethod(handler{}, "Handle", "fail"); err == nil {
		t.Fatal("CallStringMethod() did not return the method error")
	}
}

func TestCallStringMethodRejectsInvalidSignatures(t *testing.T) {
	for _, name := range []string{"Missing", "Wrong"} {
		if _, err := CallStringMethod(handler{}, name, "request"); err == nil {
			t.Fatalf("CallStringMethod(%q) error = nil", name)
		}
	}
	var receiver *handler
	if _, err := CallStringMethod(receiver, "Handle", "request"); err == nil {
		t.Fatal("CallStringMethod() accepted a typed nil receiver")
	}
}
