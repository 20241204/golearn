package main

import "fmt"

// 闭包
// 闭包是函数和它捕获的外部变量的组合。
// 通过闭包可以保持外部函数的状态。
// 闭包函数应用之一，让不同类型的函数作为参数列表调用时
func code1(code func()) {
	fmt.Println("这是一个可以执行传入函数的普通函数 code1")
	code()
}

// 这是个带有运算目的的函数
func code2(x, y int) {
	fmt.Println("这是个可以用来计算的普通函数 code2")
	fmt.Printf("%d\n", x+y)
}

/*
		闭包用于返回和传递函数。
		需求是 code1 函数调用 code2 函数
	    这个在常规使用中是不可能的，因为 code1 函数和 code2 函数类型不同
	    闭包函数 code3 对 code2 构建就可以实现这个需求
	    这个函数参数列表传入 code2 函数和 code2 包含的 int 参数列表
*/
func code3(code2 func(int, int), a, b int) func() {
	fmt.Println("这里是精心构建的闭包函数 code3")
	// 返回一个匿名函数直接调用参数列表 code2 函数传入 int 参数列表 a,b
	// 返回列表是一个没有参数列表没有返回列表的函数
	return func() {
		fmt.Println("这里是闭包函数 code3 返回的匿名函数")
		code2(a, b)
	}
}

// 闭包用于返回和传递函数。
// 闭包函数应用之二就是用于返回一个函数
func code4(a int) func(int, int) int {
	fmt.Printf("闭包函数 code4 被调用，返回 func(int,int)int 类型函数，传入 %d\n", a)
	// 匿名函数返回一个 func(int,int)int 类型的函数
	return func(b, c int) int {
		// 如果被调用则执行
		fmt.Println("匿名函数执行运算")
		return a + b + c
	}
}

// 闭包函数应用之双重禁忌
func code5(a int) (b, c func(int) int) {
	return func(i int) int {
			fmt.Printf("执行加法运算%d+%d\n", a, i)
			return a + i
		}, func(i int) int {
			fmt.Printf("再执行减法法运算%d-%d\n", a, i)
			return a - i
		}
}
func main() {
	/* 在 code1 函数中传入闭包函数 code3
	   在 code3 中传入 code2 运算函数
	   实现了 code1 函数导入了 code2 函数参数列表并执行
	*/
	code1(code3(code2, 10, 20))
	// 执行 code4 函数返回运算函数并执行运算
	a := code4(1)(1, 2)
	fmt.Printf("%d\n", a)
	// 执行 code5 函数双重禁忌
	b, c := code5(10)
	fmt.Printf("%d,%d\n", b(5), c(4))
}
