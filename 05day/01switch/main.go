package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		// switch 判断语句
		// 可以判断单值或多值
		// default为case都不符合时输出
		switch i {
		case 0:
			fmt.Printf("判断变量为%d\n", i)
		case 1:
			fmt.Printf("判断变量为%d\n", i)
		case 2, 3:
			fmt.Printf("判断变量为%d\n", i)
		default:
			fmt.Printf("判断变量为其他值%d\n", i)
		}
	}
	fmt.Println("*****************************")
	// switch 判断语句
	// 可以短声明初始化判断单值或多值
	// default为case都不符合时输出
	switch i := 9; i {
	case 0:
		fmt.Printf("判断变量为%d\n", i)
	case 1:
		fmt.Printf("判断变量为%d\n", i)
	case 2, 3:
		fmt.Printf("判断变量为%d\n", i)
	default:
		fmt.Printf("判断变量为其他值%d\n", i)
	}
	fmt.Println("*****************************")

	// switch 判断语句
	// 可以判断条件
	// default为case都不符合时输出
	i := 8
	switch {
	case i <= 0:
		fmt.Printf("判断变量%d小于等于0\n", i)
	case i <= 1:
		fmt.Printf("判断变量为%d小于等于1\n", i)
	case i <= 2:
		fmt.Printf("判断变量为%d小于等于2\n", i)
	default:
		fmt.Printf("判断变量大于其他值，为: %d\n", i)
	}
}
