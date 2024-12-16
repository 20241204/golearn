package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

var a uint

func main() {
	a = math.MaxUint64
	fmt.Printf("无符号整型 uint 变量 a 的值: %d\n", a)
	fmt.Printf("无符号整型 uint 变量 a 的类型: %T\n", a)
	fmt.Printf("无符号整型 uint 变量 a 的 reflect.TypeOf 返回值: %v\n", reflect.TypeOf(a))
	fmt.Printf("无符号整型 uint 变量 a 占用的字节大小: %d\n", unsafe.Sizeof(a))
	// ^uint(0) 会给出 uint 类型的最大值, 不一定与 math.MaxUint64 或 math.MaxUint32 相同
	// 具体取决于 uint 在当前系统上的定义。适用于平台上 uint 的最大值, 在不同的平台上可能会有不同的结果。
	// 这个表达式用来计算当前平台上 uint 类型的最大值
	fmt.Printf("无符号整型 uint 变量 a 的取值范围: 0 ~ %d\n", ^uint(0))
	// math.MaxUint64 是一个未指定类型的常量, 它的值超出了 int 类型的范围, 直接用于 fmt.Printf 时可能会导致溢出错误
	// 将 math.MaxUint64 转换为 uint64 确保 math.MaxUint64 的值在打印时不会因为溢出而导致错误
	fmt.Printf("64 位操作系统无符号整型 uint 变量 a 的取值范围: 0 ~ %d\n", uint64(math.MaxUint64))
	fmt.Printf("32 位操作系统无符号整型 uint 变量 a 的取值范围: 0 ~ %d\n", math.MaxUint32)
}
