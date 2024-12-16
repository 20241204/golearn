package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

var a float64

func main() {
	// 将值修改为 float64 能表示的最大值
	a = 1.7976931348623157e+308
	// 输出变量 a 的值，保留小数点后7位
	fmt.Printf("变量 a 值: %.7f\n", a)
	// 输出变量 a 的类型
	fmt.Printf("变量 a 类型: %T\n", a)
	// 使用 reflect.TypeOf 输出变量 a 的类型信息
	fmt.Printf("变量 a 类型 reflect.TypeOf 返回值: %v\n", reflect.TypeOf(a))
	// 输出变量 a 占用的内存大小
	fmt.Printf("变量 a 占用字节大小: %d\n", unsafe.Sizeof(a))
	// 输出 float64 类型的取值范围，使用科学计数法
	fmt.Printf("变量 a 以科学计数法显示浮点数取值范围: %.10e ~ %.10e\n", math.SmallestNonzeroFloat64, math.MaxFloat64)
	// 以标准浮点数形式输出 float64 的取值范围
	fmt.Printf("变量 a 以标准浮点数显示取值范围: %f ~ %f \n", math.SmallestNonzeroFloat64, math.MaxFloat64)
}
