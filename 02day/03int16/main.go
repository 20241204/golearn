package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

var a int16

func main() {
	a = 32767
	fmt.Printf("int16 变量 a 的数据值: %d\n", a)
	fmt.Printf("int16 变量 a 的类型: %T\n", a)
	fmt.Printf("int16 变量 a 的 reflect.TypeOf 返回值: %v\n", reflect.TypeOf(a))
	fmt.Printf("int16 变量 a 占用的字节大小: %d\n", unsafe.Sizeof(a))
	fmt.Printf("int16 变量 a 的数值允许范围: %d ~ %d\n", math.MinInt16, math.MaxInt16)
}
