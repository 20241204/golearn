package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

var a uint64

func main() {
	a = math.MaxUint64
	fmt.Printf("无符号 uint64 变量 a 值: %d\n", a)
	fmt.Printf("无符号 uint64 变量 a 类型: %T\n", a)
	fmt.Printf("无符号 uint64 变量 a 类型 reflect.TypeOf 返回值: %v\n", reflect.TypeOf(a))
	fmt.Printf("无符号 uint64 变量 a 占用字节大小: %d\n", unsafe.Sizeof(a))
	// math.MaxUint64 是一个未指定类型的常量, 它的值超出了 int 类型的范围, 直接用于 fmt.Printf 时可能会导致溢出错误
	// 将 math.MaxUint64 转换为 uint64 确保 math.MaxUint64 的值在打印时不会因为溢出而导致错误
	fmt.Printf("无符号 uint64 变量 a 取值范围: 0 ~ %d\n", uint64(math.MaxUint64))
}
