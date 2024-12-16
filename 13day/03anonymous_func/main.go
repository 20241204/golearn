package main

import "fmt"

// 匿名函数
// 可以使用变量存储匿名函数，或者直接执行匿名函数。
// 使用 reflect 包可以获取匿名函数的类型信息
var (
	a func() = func() { fmt.Println("你好我是返回函数变量的匿名函数") }
)

// 匿名函数常用在函数内部定义函数
func main() {
	// 调用匿名函数变量执行匿名函数 func()(){}()
	a()
	// 直接执行匿名函数 func()(){}() 一次性不用赋予变量
	func() { fmt.Println("你好这个是匿名函数") }()
	// 获取匿名函数类型
	fmt.Printf("%T\n", func() { fmt.Println("你好这个是匿名函数") })
}
