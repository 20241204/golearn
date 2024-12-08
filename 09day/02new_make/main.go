package main

import (
	"fmt"
)

var (
	a *int
	b int
)

// new make 函数
func main() {
	// 声明的 指针变量一定要初始化，指针变量只能存地址，否则值为 nil
	// 如果此时打印 *a 会得到空指针一场，因为 nil 类型地址没有值
	fmt.Printf("指针变量 a \n\t值(地址的值): %v \t类型: %T\n", a, a)
	// 存入已知变量内存地址
	a = &b
	fmt.Printf("将变量 b 的地址赋值给指针变量 a 后(即指针变量 a 的初始化)\n\t值(地址的值): %v \t类型: %T\t 值(地址)指向的值(即变量 b 的值): %d\n", a, a, *a)
	b = 10
	fmt.Printf("将变量 b 的值修改后指针变量 a 的参数\n\t值(地址的值): %v \t类型: %T\t 值(地址)指向的值(即变量 b 值): %d\n", a, a, *a)
	// 同时也可以用指针变量 a 影响变量 b 的值
	*a = 20
	fmt.Printf("修改指针变量 a 的地址值影响变量 b\n\t值(地址的值): %v \t类型: %T\t 值(地址)指向的值(即变量 b 值): %d\n", a, a, *a)
	// new 函数可以申请内存用于创建 基本数据类型的 指针变量 非基本数据类型比如 *string 指针变量也可以创建
	// 初始化 int类型 的 指针类型 的变量
	var c *int = new(int)
	fmt.Printf("指针变量 c 通过 new 的声明初始化\n\t值(地址的值): %v \t类型: %T\t 值(地址)指向的值: %d\n", c, c, *c)
	// make 函数也可以申请内存初始化变量
	// 用于创建 slice map 和 channel 的变量而非指针变量，
	//因为都是引用类型的变量，所以 make 创建的变量类型是变量自身
	d := make([]int, 5, 10)
	fmt.Printf("变量 d 通过 make 的声明初始化\n\t值: %v \t类型: %T\t长度: %d\t容量: %d\n", d, d, len(d), cap(d))
}
