// calculator/calculator_test.go
package calculator

import "testing"

func TestAdd(t *testing.T) {
	calc := New()
	result := calc.Add(10, 20)

	if result != 30 {
		t.Errorf("Add(10, 20) = %f; want 30", result)
	}
}

func TestDivide(t *testing.T) {
	calc := New()

	// 正常情况
	result, err := calc.Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) returned error: %v", err)
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %f; want 5", result)
	}

	// 除零错误
	_, err = calc.Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) should return error")
	}
}

func TestHistory(t *testing.T) {
	calc := New()
	calc.Add(1, 2)
	calc.Subtract(5, 3)

	history := calc.History()
	if len(history) != 2 {
		t.Errorf("History length = %d; want 2", len(history))
	}
}
