package main

import "fmt"

func main() {
	// 正常循环
	for a := 1; a < 6; a++ {
		fmt.Println("打印数:", a)
	}
	fmt.Println("*********************")

	// 没有初始值或结束表达式
	// 此时的 for 语句只有一个条件判断
	a := 0
	for a < 6 {
		fmt.Println("打印数:", a)
		// 内部的自增表达式可以作为结束表达式
		a++
	}
	fmt.Println("*********************")

	// for range 打印 索引,值
	// 索引按字节计数：在 UTF-8 编码中，中文字符占用 3 个字节，因此索引是 3 的倍数。
	// 使用 %c 可以正确输出 rune 类型的字符。
	for i, v := range "abcd未满18岁~" {
		fmt.Printf("打印索引: %d -> 值: %c\n", i, v)
	}
	fmt.Println("*********************")

	// 嵌套循环
	for a := 1; a < 4; a++ {
		for b := 1; b < 4; b++ {
			// 打印外循环数和内循环数
			fmt.Printf("外: %d->内: %d\t", a, b)
		}
		fmt.Println()
	}

	// 死循环
	// 注意：这会导致无限输出，确保使用时有适当的退出条件
	// for {
	//     fmt.Println("死亡循环打印")
	// }
}
