package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

// 声明 int32 类型的变量
var a int32

func main() {
	a = 2147483647

	// 打印 int32 变量的数值
	fmt.Printf("int32 变量 a 的数值: %d\n", a)
	// 打印 int32 变量的类型
	fmt.Printf("int32 变量 a 的类型: %T\n", a)
	// 打印 int32 变量的 reflect.TypeOf 返回值
	fmt.Printf("int32 变量 a 的 reflect.TypeOf 返回值: %v\n", reflect.TypeOf(a))
	// 打印 int32 变量的字节大小
	fmt.Printf("int32 变量 a 占用的字节大小: %d\n", unsafe.Sizeof(a))
	// 打印 int32 变量的数值范围
	fmt.Printf("int32 变量 a 的数值允许范围: %d ~ %d\n", math.MinInt32, math.MaxInt32)
}
