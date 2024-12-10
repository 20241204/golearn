package main

import "fmt"

/*
*练习
 1. 求数组[1,3,5,7,8]所有元素的和
 2. 找出数组中和为指定值的两个元素下标，
    比如从数组[1,3,5,7,8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
*/

var (
	arrayInt = [...]int{1, 3, 5, 7, 8} // 声明初始化 int 类型数组
	sum      int                       // 加和变量
)

const (
	FLAGNUM int = 8 // 指定数组任意两个元素之和为8
)

func code1() { //1. 求数组[1,3,5,7,8]所有元素的和
	// for range 匿名变量（哑光变量）遍历数组
	for _, v := range arrayInt {
		sum += v
	}
	fmt.Printf("数组%v元素之和为：%d", arrayInt, sum)
}

func code2() { //2. 找出数组[1,3,5,7,8]中和为指定值8的两个元素下标，
	// for range 匿名变量（哑光变量）嵌套遍历数组
	for i, v1 := range arrayInt {
		for j, v2 := range arrayInt {
			if v1+v2 == FLAGNUM {
				fmt.Printf("数组%v中任意两个元素和为%d的索引和值: 索引(%d,%d) -> 值(%d,%d)\n", arrayInt, FLAGNUM, i, j, v1, v2)
			}
		}
	}
}

func main() {
	code1()
	fmt.Println("\n************************************************")
	code2()
}
