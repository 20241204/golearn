package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

var a complex128

func main() {
	// 初始化复数
	// 使用 complex(1.0, 2.0) 创建了一个 complex128 类型的复数 a，其中 1.0 是实部，2.0 是虚部。
	a = complex(1.0, 2.0)

	fmt.Printf("变量 a 值: %v\n", a)
	fmt.Printf("变量 a 类型: %T\n", a)
	fmt.Printf("变量 a 类型 reflect.TypeOf 返回值: %v\n", reflect.TypeOf(a))
	fmt.Printf("变量 a 占用字节: %d\n", unsafe.Sizeof(a))

	// 输出 complex128 的组成部分 float64 的取值范围
	// 复数类型的值：通过 complex 函数创建复数，complex128 由两个 float64 组成。
	// 内存占用：使用 unsafe.Sizeof() 可以获取复数类型的内存占用。
	// 取值范围：复数的取值范围通常由其组成部分 float64 的取值范围决定。
	// 输出 complex128 类型的取值范围，使用科学计数法
	fmt.Printf("变量 a 的实部和虚部的取值范围（科学计数法）: %.10e ~ %.10e\n", math.SmallestNonzeroFloat64, math.MaxFloat64)
	// 以标准浮点数形式输出 complex128 的取值范围
	fmt.Printf("变量 a 的实部和虚部的取值范围（标准浮点数）: %f ~ %f\n", math.SmallestNonzeroFloat64, math.MaxFloat64)
}
