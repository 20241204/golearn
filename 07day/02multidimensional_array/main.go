package main

import "fmt"

// 声明多维数组变量
var multidimensional_array_int [3][3]int

// multidimensional array 多维数组练习
func main() {
	// 初始化赋值
	multidimensional_array_int = [3][3]int{
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
		[3]int{7, 8, 9},
	}

	// 简短声明初始化赋值
	multidimensional_array_float32 := [3][3]float32{
		{1.99, 2.99, 3.99},
		{4.99, 5.99, 6.99},
		{7.99, 8.99, 9.99},
	}
	fmt.Printf("int类型多维数组: %v, 类型: %T\n", multidimensional_array_int, multidimensional_array_int)
	fmt.Printf("float32类型多维数组: %v, 类型: %T\n", multidimensional_array_float32, multidimensional_array_float32)
	fmt.Println("********************************")

	// 遍历二维数组
	fmt.Println("遍历int类型多维数组, 索引, 值, 类型: ")
	for i := 0; i < len(multidimensional_array_int); i++ {
		fmt.Printf("第 %d 行: %v\n", i, multidimensional_array_int[i])
		fmt.Println("\t内部数组遍历, 索引, 值, 类型: ")
		for j := 0; j < len(multidimensional_array_int[i]); j++ {
			fmt.Printf("\t(%d,%d): %v, 类型: %T\n", i, j, multidimensional_array_int[i][j], multidimensional_array_int[i][j])
		}
	}

	fmt.Println("遍历float32类型多维数组, 索引, 值, 类型: ")
	for i, row := range multidimensional_array_float32 {
		fmt.Printf("第 %d 行: %v\n", i, row)
		fmt.Println("\t内部数组遍历, 索引, 值, 类型: ")
		for j, value := range row {
			fmt.Printf("\t(%d,%d): %v, 类型: %T\n", i, j, value, value)
		}
	}
}
