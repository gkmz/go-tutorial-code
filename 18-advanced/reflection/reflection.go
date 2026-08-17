package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// User 是反射示例使用的结构体。
type User struct {
	Name  string `json:"name" validate:"required"`  // Name 是用户名称。
	Age   int    `json:"age"`                       // Age 是用户年龄。
	Email string `json:"email" validate:"required"` // Email 是用户邮箱。
}

// FieldDescription 保存一个可导出结构体字段的名称、JSON 名称和值。
type FieldDescription struct {
	Name     string // Name 是 Go 结构体字段名称。
	JSONName string // JSONName 是 json 标签声明的字段名称。
	Value    any    // Value 是字段当前值。
}

// Describe 返回结构体中可导出字段的描述。
// value 可以是结构体或指向结构体的非 nil 指针。
func Describe(value any) ([]FieldDescription, error) {
	reflected, err := structValue(value)
	if err != nil {
		return nil, err
	}

	typ := reflected.Type()
	fields := make([]FieldDescription, 0, reflected.NumField())
	for index := 0; index < reflected.NumField(); index++ {
		fieldValue := reflected.Field(index)
		fieldType := typ.Field(index)
		if !fieldValue.CanInterface() {
			continue
		}

		jsonName := strings.Split(fieldType.Tag.Get("json"), ",")[0]
		if jsonName == "" {
			jsonName = fieldType.Name
		}
		if jsonName == "-" {
			continue
		}

		fields = append(fields, FieldDescription{
			Name:     fieldType.Name,
			JSONName: jsonName,
			Value:    fieldValue.Interface(),
		})
	}
	return fields, nil
}

// SetField 修改 target 指向的结构体字段。
// value 的动态类型必须可以直接赋值给目标字段。
func SetField(target any, fieldName string, value any) error {
	reflected := reflect.ValueOf(target)
	if !reflected.IsValid() || reflected.Kind() != reflect.Pointer || reflected.IsNil() {
		return errors.New("target must be a non-nil pointer")
	}

	element := reflected.Elem()
	if element.Kind() != reflect.Struct {
		return errors.New("target must point to a struct")
	}

	field := element.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("field %q does not exist", fieldName)
	}
	if !field.CanSet() {
		return fmt.Errorf("field %q is not settable", fieldName)
	}

	newValue := reflect.ValueOf(value)
	if !newValue.IsValid() {
		return fmt.Errorf("nil cannot be assigned to field %q of type %s", fieldName, field.Type())
	}
	if !newValue.Type().AssignableTo(field.Type()) {
		return fmt.Errorf("value of type %s cannot be assigned to field %q of type %s", newValue.Type(), fieldName, field.Type())
	}

	field.Set(newValue)
	return nil
}

// ValidateRequired 检查结构体中带有 validate:"required" 标签的字段是否为零值。
// value 可以是结构体或指向结构体的非 nil 指针。
func ValidateRequired(value any) error {
	reflected, err := structValue(value)
	if err != nil {
		return err
	}

	typ := reflected.Type()
	var missing []string
	for index := 0; index < reflected.NumField(); index++ {
		fieldType := typ.Field(index)
		if fieldType.Tag.Get("validate") != "required" {
			continue
		}
		fieldValue := reflected.Field(index)
		if !fieldValue.CanInterface() {
			continue
		}
		if fieldValue.IsZero() {
			missing = append(missing, fieldType.Name)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("required fields are empty: %s", strings.Join(missing, ", "))
	}
	return nil
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
		return reflect.Value{}, fmt.Errorf("value must be a struct, got %s", reflected.Kind())
	}
	return reflected, nil
}
