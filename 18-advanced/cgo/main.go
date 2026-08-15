package main

/*
#include <stdlib.h>
static int add(int a, int b) { return a + b; }
*/
import "C"

import "fmt"

func main() {
	// CGO 会引入跨语言调用成本，指针传递还必须遵守 CGO 指针规则。
	fmt.Println(C.add(2, 3))
}
