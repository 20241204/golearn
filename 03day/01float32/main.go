package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

var a float32

func main() {
	// 赋值一个接近 float32 最大值的值，确保在范围内
	a = 3.4028235e+38
	// 输出变量 a 的值
	// %f默认会输出6位小数，如果希望展示更高精度，可以指定精度如%.7f
	fmt.Printf("变量 a 值: %.7f\n", a)
	// 输出变量 a 的类型
	fmt.Printf("变量 a 类型: %T\n", a)
	// 使用 reflect.TypeOf 输出变量 a 的类型信息
	fmt.Printf("变量 a 类型 reflect.TypeOf 返回值: %v\n", reflect.TypeOf(a))
	// 输出变量 a 占用的内存大小
	fmt.Printf("变量 a 占用字节大小: %d\n", unsafe.Sizeof(a))
	// 输出 float32 类型的取值范围
	// %.10e 的具体含义
	// %e：以科学计数法表示浮点数。
	// %.10e：表示以科学计数法表示浮点数，同时保留小数点后 10 位数字。
	// 在编程中，这通常写作 a.eb，其中 a 是系数，e 表示以 10 为底的指数，b 是指数值。例如：
	// 3.14e+02 表示 3.14 × 10^2，即 314。
	// 2.71828e-03 表示 2.71828 × 10^-3，即 0.00271828。
	fmt.Printf("变量 a 以科学计数法表示浮点数表示取值范围: %.10e ~ %.10e\n", math.SmallestNonzeroFloat32, math.MaxFloat32)
	// 以标准浮点数形式输出 float32 的取值范围
	fmt.Printf("变量 a 以标准浮点数表示取值范围: %f ~ %f\n", math.SmallestNonzeroFloat32, math.MaxFloat32)
}
