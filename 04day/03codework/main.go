package main

import "fmt"

/*
打印九九乘法表
*/
func code1() { // 打印9个数字
	for i := 1; i < 10; i++ {
		fmt.Printf("打印1~9: %d\n", i)
	}
}
func code2() { // 需要两个数相乘
	for i := 1; i < 10; i++ {
		fmt.Printf("打印1~9相乘: %dx%d=%d\n", i, i, i*i)
	}
}

func code3() { // 嵌套两个for循环实现9x9矩阵乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			// 打印 9x9 乘法表矩阵
			fmt.Printf("%dx%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func code4() { // 嵌套两个for循环实现9x9正三角乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			// 打印 9x9 乘法表矩阵
			fmt.Printf("%dx%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
func main() {
	code1()
	fmt.Println("*********************")

	code2()
	fmt.Println("*********************")

	code3()
	fmt.Println("*********************")

	code4()
	fmt.Println("*********************")
}
