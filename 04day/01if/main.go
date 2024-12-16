package main

import "fmt"

var a int

// 鬼畜打印
func main() {
	a = 17
	// 单分支结构判断
	if a < 18 {
		fmt.Println("欸呀人家还未满18岁~")
	}
	fmt.Println("*********************")

	a = 19
	// 多分支结构判断
	if a < 18 {
		fmt.Println("欸呀人家还未满18岁~")
	} else if a <= 80 {
		fmt.Println("欸呀人家已满18岁~但小于80岁~")
	} else {
		fmt.Println("欸呀人家不知道有多少岁~")
	}
	fmt.Println("*********************")

	// 声明内部变量判断作用域仅在if语句之内
	if a := 80; a == 80 {
		fmt.Println("欸呀人家刚满80岁~")
	} else {
		fmt.Println("欸呀人家不知道有多少岁~")
	}
	fmt.Println("*********************")

	// 嵌套判断
	if a := 80; a == 80 {
		fmt.Println("欸呀人家刚满80岁~")
		if b := 18; b == 18 {
			fmt.Println("欸呀人家的人家刚满18岁~")
		}
	}
	fmt.Println("*********************")

}
