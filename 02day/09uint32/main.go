package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

var a uint32

func main() {
	a = 4294967295
	fmt.Printf("无符号 uint32 变量 a 值: %d\n", a)
	fmt.Printf("无符号 uint32 变量 a 类型: %T\n", a)
	fmt.Printf("无符号 uint32 变量 a 类型 reflect.TypeOf 返回值: %v\n", reflect.TypeOf(a))
	fmt.Printf("无符号 uint32 变量 a 占用字节大小: %d\n", unsafe.Sizeof(a))
	fmt.Printf("无符号 uint32 变量 a 取值范围: 0 ~ %d\n", math.MaxUint32)
}
