//go:build cgo

package main

/*
#include <stdlib.h>
#include <string.h>

static int add(int a, int b) {
	return a + b;
}

static char* duplicate_message(const char* input) {
	size_t length = strlen(input);
	char* output = malloc(length + 1);
	if (output == NULL) {
		return NULL;
	}
	memcpy(output, input, length + 1);
	return output;
}

static int duplicate_message_checked(const char* input, char** output) {
	if (input == NULL || output == NULL) {
		return 1;
	}

	*output = duplicate_message(input);
	if (*output == NULL) {
		return 2;
	}
	return 0;
}
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

// add 调用 C 函数完成整数相加，用于观察最小跨语言调用边界。
func add(left, right int) int {
	return int(C.add(C.int(left), C.int(right)))
}

// duplicateMessage 将 Go 字符串复制到 C 堆，再把 C 返回的字符串复制回 Go。
func duplicateMessage(message string) (string, error) {
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))

	cResult := C.duplicate_message(cMessage)
	if cResult == nil {
		return "", errors.New("C allocation failed")
	}
	defer C.free(unsafe.Pointer(cResult))

	return C.GoString(cResult), nil
}

// duplicateMessageWithStatus 将 C 状态码映射为带语义的 Go 错误。
func duplicateMessageWithStatus(message string) (string, error) {
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))

	var cResult *C.char
	switch status := C.duplicate_message_checked(cMessage, &cResult); status {
	case 0:
		defer C.free(unsafe.Pointer(cResult))
		return C.GoString(cResult), nil
	case 1:
		return "", errors.New("C received an invalid argument")
	case 2:
		return "", errors.New("C allocation failed")
	default:
		return "", fmt.Errorf("C duplicate_message failed with status %d", status)
	}
}
