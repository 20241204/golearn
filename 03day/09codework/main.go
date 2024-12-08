package main

import (
	"fmt"
	"math"
	"unicode"
)

/*
练习题
1. 常量声明及练习iota定义数量级(1024)
2.编写代码分别定义一个整型、浮点型、布尔型、字符串型变量

	使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。

3. 从网上获取办法编写代码统计出字符串"hello沙河hello小王子hello"中汉字的数量。
*/

const (
	_  = iota
	TB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	KB = 1 << (10 * iota)
	B  = 1 << (10 * iota)
)

var (
	varA int
	varB float64
	varC bool
	varD string
)

func code1() { // 练习1: iota 常量计数器数量级声明初始化

	// 数量级打印
	fmt.Printf("数量级打印: %dB => %dKB => %dMB => %dGB => %dTB\n", B, KB, MB, GB, TB)
}

func code2() { // 练习2:声明初始化多个类型变量，并打印值和类型
	// 初始化不同类型变量
	varA = math.MaxInt
	varB = math.MaxFloat64
	varD = "地球欢迎你！为你开天辟地！Golang中的魅力充满着朝气！"
	fmt.Printf("变量值%d的数据类型: %T\n", varA, varA)
	fmt.Printf("变量值%f的数据类型: %T\n", varB, varB)
	fmt.Printf("变量值%t的数据类型: %T\n", varC, varC)
	fmt.Printf("变量值\"%s\"的数据类型: %T\n", varD, varD)
}

func code3() { // 练习3：编写代码统计出字符串"hello沙河hello小王子hello"中汉字的数量。
	// 初始化字符串变量
	// 计数器
	varA = 0
	// 待处理字符串
	varD = "hello沙河hello小王子hello"
	// 迭代字符串
	for _, varDByte := range varD {
		// 判断字符是否在汉字的 Unicode 码点范围内
		if varDByte >= '\u4e00' && varDByte <= '\u9fff' {
			varA++
		}
	}
	fmt.Printf("\"%s\"中汉字的数量: %d\n", varD, varA)
}

// 第二种写法
func code4() { // 练习3：编写代码统计出字符串"hello沙河hello小王子hello"中汉字的数量。
	// 初始化字符串变量
	// 计数器
	varA = 0
	// 待处理字符串
	varD = "hello沙河hello小王子hello"
	// 迭代字符串
	for _, varDByte := range varD {
		// 判断字符是否在汉字 Unicode 的 RangeTable 为 Han
		if unicode.Is(unicode.Han, varDByte) {
			varA++
		}
	}
	fmt.Printf("\"%s\"中汉字的数量: %d\n", varD, varA)
}

func main() {
	code1()
	code2()
	code3()
	code4()
}
