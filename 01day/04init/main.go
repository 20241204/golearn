package main

import "fmt"

// 函数外只能声明变量、常量、函数和类型并会优先加载
var name string

// init 在函数中优先加载，通常用于初始变量赋值等操作
// init 可以定义多个
func init() {
	fmt.Println("init函数执行---->此时name:", name)
	name = "golang"
	fmt.Println("name初始化赋值---->此时name:", name)
}

// 多个 init 函数按顺序执行
func init() {
	// 函数内定义变量无法被其他函数调用
	age := 19
	fmt.Print("init函数执行---->此时age:", age)
}

// main 函数，go程序唯一入口
// 一个 package 包有且只能有一个 main 函数
// 最后执行 main 函数
func main() {
	// \n 为换行转义字符
	fmt.Printf("\nmain函数执行---->name: %s", name)
}
