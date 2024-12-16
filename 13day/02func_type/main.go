package main

import (
	"fmt"
	"reflect"
)

// 以下的函数以函数类型为参数调用也被称为高阶函数
// 创建一个 func() int 表示一个无参数并返回 int 类型的函数
func code1() int {
	return 1
}

// 创建一个 func(int) int 表示一个接收 int 类型参数并返回 int 类型的函数
func code2(a int) int {
	return 1
}

// 创建一个 func(func() int) func(int) int 表示一个接收无参数返回 int 的函数并返回接收 int 参数返回 int 的函数
func code3(a func() int) func(int) int {
	return code2
}
func main() {
	fmt.Println("一个无参数列表有返回列表的函数类型为")
	// 获取 code1 函数的类型
	fmt.Printf("%T\n%v\n", code1, reflect.TypeOf(code1))
	fmt.Println("一个有参数列表有返回列表的函数类型为")
	// 获取 code2 函数的类型
	fmt.Printf("%T\n%v\n", code2, reflect.TypeOf(code2))
	fmt.Println("一个参数列表为 无参数列表有返回列表函数 返回列表为 有参数列表有返回列表函数 的函数类型为")
	// 获取 code3 函数的类型
	fmt.Printf("%T\n%v\n", code3, reflect.TypeOf(code3))
	// 获取 code3 函数导入参数列表 code1 函数后的返回列表 code2 类型
	a := code3(code1)
	fmt.Println("一个参数列表为 无参数列表有返回列表函数 返回列表为 有参数列表有返回列表函数 的函数导入一个函数后的返回列表为 有参数列表有返回列表函数 的类型为")
	fmt.Printf("%T\n%v\n", a, reflect.TypeOf(a))
}
