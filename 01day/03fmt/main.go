package main

import "fmt"

func main() {
	fmt.Print("Print不会换行---->")
	fmt.Println("Println会换行---->")
	// %s 输出变量的字符串值
	// %d 输出变量的整型值
	fmt.Printf("我是%s,我%d岁了，Printf会调用格式打印", "glang", 19)
}
