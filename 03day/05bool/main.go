package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

var a bool

func main() {
	// bool 类型变量默认是false
	// bool 类型不参与运算
	// bool 不参与类型转换
	// 与 python 不同 int 无法强转 bool 类型
	// 使用 %t 格式化输出布尔值
	fmt.Printf("变量 a 值: %t\n", a)
	// 输出变量 a 的类型
	fmt.Printf("变量 a 类型: %T\n", a)
	// 使用 reflect.TypeOf 输出变量 a 的类型信息
	fmt.Printf("变量 a 类型 reflect.TypeOf 返回值: %v\n", reflect.TypeOf(a))
	// 输出变量 a 占用的内存大小
	fmt.Printf("变量 a 占用字节大小: %d\n", unsafe.Sizeof(a))
	// 输出布尔类型的取值范围（可以通过注释描述，或者简单输出 true 和 false）
	fmt.Printf("变量 a 可能的取值范围: %t ~ %t\n", false, true)
}
