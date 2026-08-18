package main

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	stringType = reflect.TypeFor[string]()
	errorType  = reflect.TypeFor[error]()
)

// CallStringMethod 调用签名为 func(string) (string, error) 的导出方法。
func CallStringMethod(receiver any, methodName, input string) (string, error) {
	value := reflect.ValueOf(receiver)
	if !value.IsValid() {
		return "", errors.New("receiver must not be nil")
	}
	if (value.Kind() == reflect.Pointer || value.Kind() == reflect.Interface) && value.IsNil() {
		return "", errors.New("receiver must not contain a nil value")
	}
	method := value.MethodByName(methodName)
	if !method.IsValid() {
		return "", fmt.Errorf("method %q does not exist", methodName)
	}

	methodType := method.Type()
	if methodType.NumIn() != 1 || methodType.In(0) != stringType ||
		methodType.NumOut() != 2 || methodType.Out(0) != stringType || methodType.Out(1) != errorType {
		return "", fmt.Errorf("method %q must have signature func(string) (string, error)", methodName)
	}

	// 完整校验签名后再调用，避免把参数错误变成 reflect.Call 的 panic。
	results := method.Call([]reflect.Value{reflect.ValueOf(input)})
	output := results[0].String()
	if results[1].IsNil() {
		return output, nil
	}
	return output, results[1].Interface().(error)
}
