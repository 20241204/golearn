package main

import "fmt"

func main() {
	// continue 跳过单循环
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			// 当内循环值为 3 则跳过此值打印继续执行下一个值打印
			if j == 3 {
				fmt.Printf("值为3跳过内循环并继续\t")
				continue
			}
			fmt.Printf("外:%d->内:%d\t", i, j)
		}
		fmt.Println()
	}
	fmt.Println("*************************************")

	// break 跳出单循环
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			// 当内循环值为 3 则跳过此值打印并退出内循环
			if j == 3 {
				fmt.Printf("值为3跳过并退出内循环\t")
				break
			}
			fmt.Printf("外:%d->内:%d\t", i, j)
		}
		fmt.Println()
	}
	fmt.Println("*************************************")

	// break flag标记 跳出双循环
	flag := false
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			// 当内循环值为 3 则跳过此值打印并退出内循环
			if j == 3 {
				fmt.Printf("值为3跳过并退出内循环\t")
				flag = true
				break
			}
			fmt.Printf("外:%d->内:%d\t", i, j)
		}
		if flag {
			fmt.Printf("flag值为true跳过并退出外循环\t")
			break
		}
		fmt.Println()
	}
	fmt.Println("*************************************")

	// break 直接跳出双循环标签语句
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				fmt.Printf("值为3跳过并直接退出双循环标签\t")
				break outerLoop
			}
			fmt.Printf("外:%d->内:%d\t", i, j)
		}
		fmt.Println()
	}
	fmt.Println("*************************************")

	// goto 跳转至标签语句
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			// 当内循环值为 3 则跳转至A标签 lable
			if j == 3 {
				fmt.Printf("值为3跳转至A lable\t")
				goto A
			}
			fmt.Printf("外:%d->内:%d\t", i, j)
		}
		fmt.Println()
	}
A:
	fmt.Println("结束")
}
