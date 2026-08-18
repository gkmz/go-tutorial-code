// Package exercises 提供反射章节代码练习的参考实现。
package exercises

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// CopyMatchingFields 按字段名复制类型可直接赋值的导出字段。
func CopyMatchingFields(destination, source any) error {
	destinationValue := reflect.ValueOf(destination)
	if !destinationValue.IsValid() || destinationValue.Kind() != reflect.Pointer || destinationValue.IsNil() {
		return errors.New("destination must be a non-nil pointer")
	}
	destinationValue = destinationValue.Elem()
	if destinationValue.Kind() != reflect.Struct {
		return errors.New("destination must point to a struct")
	}

	sourceValue, err := structValue(source)
	if err != nil {
		return fmt.Errorf("source: %w", err)
	}
	for index := 0; index < destinationValue.NumField(); index++ {
		destinationField := destinationValue.Field(index)
		if !destinationField.CanSet() {
			continue
		}
		name := destinationValue.Type().Field(index).Name
		sourceField := sourceValue.FieldByName(name)
		if !sourceField.IsValid() || !sourceField.CanInterface() {
			continue
		}
		if !sourceField.Type().AssignableTo(destinationField.Type()) {
			return fmt.Errorf("field %q has incompatible type %s", name, sourceField.Type())
		}
		destinationField.Set(sourceField)
	}
	return nil
}

// MissingRequired 返回嵌套结构体中 required 字段的路径。
func MissingRequired(value any) ([]string, error) {
	reflected := reflect.ValueOf(value)
	if !reflected.IsValid() {
		return nil, errors.New("value must not be nil")
	}
	visited := make(map[visit]struct{})
	return collectMissing(reflected, "", visited), nil
}

type visit struct {
	typeOf  reflect.Type
	pointer uintptr
}

func collectMissing(value reflect.Value, prefix string, visited map[visit]struct{}) []string {
	for value.IsValid() && value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return nil
		}
		current := visit{typeOf: value.Type(), pointer: value.Pointer()}
		if _, exists := visited[current]; exists {
			return nil
		}
		visited[current] = struct{}{}
		defer delete(visited, current)
		value = value.Elem()
	}
	if !value.IsValid() || value.Kind() != reflect.Struct {
		return nil
	}

	var missing []string
	typ := value.Type()
	for index := 0; index < value.NumField(); index++ {
		fieldType := typ.Field(index)
		fieldValue := value.Field(index)
		if !fieldValue.CanInterface() {
			continue
		}
		name := fieldName(fieldType)
		path := name
		if prefix != "" {
			path = prefix + "." + name
		}
		if fieldType.Tag.Get("validate") == "required" && fieldValue.IsZero() {
			missing = append(missing, path)
			continue
		}
		missing = append(missing, collectMissing(fieldValue, path, visited)...)
	}
	return missing
}

func fieldName(field reflect.StructField) string {
	name := strings.Split(field.Tag.Get("json"), ",")[0]
	if name == "" || name == "-" {
		return field.Name
	}
	return name
}

func structValue(value any) (reflect.Value, error) {
	reflected := reflect.ValueOf(value)
	if !reflected.IsValid() {
		return reflect.Value{}, errors.New("value must not be nil")
	}
	if reflected.Kind() == reflect.Pointer {
		if reflected.IsNil() {
			return reflect.Value{}, errors.New("value must not be a nil pointer")
		}
		reflected = reflected.Elem()
	}
	if reflected.Kind() != reflect.Struct {
		return reflect.Value{}, errors.New("value must be a struct")
	}
	return reflected, nil
}
