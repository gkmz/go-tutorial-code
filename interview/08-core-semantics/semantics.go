// Package semantics 提供 Go 语言核心语义的最小可运行实验。
package semantics

import (
	"errors"
	"fmt"
	"reflect"
)

// AppendValue 展示 Slice 头按值传递；调用方需要接收返回的 Slice。
func AppendValue(values []int, value int) []int {
	return append(values, value)
}

// MutateElement 展示 Slice 副本仍可能修改共享底层数组。
func MutateElement(values []int, index, value int) {
	if index >= 0 && index < len(values) {
		values[index] = value
	}
}

// TypedNilError 返回一个动态类型存在、动态值为 nil 的 Interface。
func TypedNilError() error {
	var err *customError
	return err
}

type customError struct{ message string }

func (e *customError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return e.message
}

// DeferEvaluation 返回 defer 参数求值与闭包读取变量的差异。
func DeferEvaluation() (argumentValue, closureValue int) {
	argumentValue = 1
	defer func(value int) { argumentValue = value }(argumentValue)
	defer func() { closureValue = argumentValue }()
	argumentValue = 2
	return argumentValue, closureValue
}

// WrapNotFound 将底层错误包装成可通过 errors.Is 判断的错误。
func WrapNotFound(resource string) error {
	return fmt.Errorf("load %s: %w", resource, ErrNotFound)
}

// ErrNotFound 表示资源不存在。
var ErrNotFound = errors.New("not found")

// MapValues 使用泛型将输入 Slice 映射为新的 Slice。
func MapValues[T any, R any](values []T, mapper func(T) R) []R {
	result := make([]R, len(values))
	for i, value := range values {
		result[i] = mapper(value)
	}
	return result
}

// InspectPointer 返回反射实验中常用的类型、Kind 和可设置状态。
func InspectPointer(value any) (typeName string, kind reflect.Kind, canSet bool) {
	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Pointer && !rv.IsNil() {
		return rv.Type().String(), rv.Elem().Kind(), rv.Elem().CanSet()
	}
	return rv.Type().String(), rv.Kind(), rv.CanSet()
}
