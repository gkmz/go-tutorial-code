package main

import "fmt"

// 自定义错误类型

// 定义错误类型
type ValidationError struct {
	Field string
	Value interface{}
	Msg   string
}

// 实现 error 接口
func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field '%s': %s (value: %v)",
		e.Field, e.Msg, e.Value)
}

// 使用自定义错误
func validateAge(age int) error {
	if age < 0 || age > 150 {
		return ValidationError{
			Field: "age",
			Value: age,
			Msg:   "age must be between 0 and 150",
		}
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return ValidationError{
			Field: "name",
			Value: name,
			Msg:   "name cannot be empty",
		}
	}
	return nil
}

// 检查自定义错误
func checkValidation(err error) {
	if err != nil {
		fmt.Println("Error:", err)

		// 类型断言
		if ve, ok := err.(ValidationError); ok {
			fmt.Printf("  Field: %s\n", ve.Field)
			fmt.Printf("  Value: %v\n", ve.Value)
			fmt.Printf("  Message: %s\n", ve.Msg)
		}
	}
}

func main() {
	fmt.Println("=== 自定义错误类型 ===")

	// 测试年龄验证
	fmt.Println("\n--- Validate Age ---")
	err := validateAge(200)
	checkValidation(err)

	// 测试名称验证
	fmt.Println("\n--- Validate Name ---")
	err = validateName("")
	checkValidation(err)

	// 正常情况
	fmt.Println("\n--- Valid Input ---")
	err = validateAge(25)
	if err == nil {
		fmt.Println("Age is valid")
	}
}
