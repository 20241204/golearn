package main

import "fmt"

/*
变量作用域分为三种
全局变量
局部变量

	函数内变量
	代码块变量

对于函数，作用域优先级顺序

	局部变量->全局变量
*/
var (
	a string = "全局变量"
)

func code1() {
	// 变量作用域
	// 全局变量、局部变量、函数内变量和代码块内变量的优先级从高到低。
	a := "函数内定义变量"
	fmt.Printf("%s\n", a)
	for a := 0; a < 1; a++ {
		fmt.Printf("代码块内变量%d\n", a)
	}
}

func main() {
	a := "局部变量" // 局部变量优先
	fmt.Printf("%s\n", a)
	code1()

}
