package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

var (
	a rune
	b string
)

func main() {
	// 初始化赋值字符变量 a
	// rune 是 int32 的别名，表示一个 Unicode 代码点（字符）。
	// 用于表示 Unicode 字符，包括汉字、特殊符号等。
	a = 'A'
	// %c 输出变量字符值
	fmt.Printf("字符变量 rune 变量 a 值: %c\n", a)
	fmt.Printf("字符变量 rune 变量 a int32值: %d\n", a)
	// %U 输出变量 Unicode 码点值
	fmt.Printf("字符变量 rune 变量 a Unicode码点值: %U\n", a)
	fmt.Printf("字符变量 rune 变量 a 类型: %T\n", a)
	fmt.Printf("字符变量 rune 变量 a 类型 reflect.TypeOf 返回值: %v\n", reflect.TypeOf(a))
	fmt.Printf("字符变量 rune 变量 a 占用字节大小: %d\n", unsafe.Sizeof(a))
	fmt.Printf("字符变量 rune 变量 a 取值范围: %d ~ %d \n", 0, 0x10FFFF)

	// 定义并打印多个字符
	b = "你好世界! 我是Golang"
	fmt.Printf("字符串变量 b 值: %s\n", b)

	// 使用 range 遍历字符串
	fmt.Println("\n使用 range 遍历字符串并输出每个字符的 int32 值和 Unicode 码点:")
	for _, i := range b {
		fmt.Printf("字符值: %c\tint32: %d\tUnicode码点: %U\n", i, i, i)
	}

	// 使用切片截取字符串的一部分
	fmt.Println("\n使用切片截取字符串的一部分:")
	slice := b[6:12] // 这里 b[6:12] 截取了字符串 b 的一部分
	fmt.Printf("切片值: %s\n", slice)

	// 将字符串转换为 []rune 切片
	runeSlice := []rune(b)
	fmt.Println("\n将字符串转换为 []rune 切片:")
	fmt.Printf("rune 切片值: %v\n", runeSlice)
	fmt.Printf("rune 切片长度: %d\n", len(runeSlice))
	// 切片容量 cap
	fmt.Printf("rune 切片容量: %d\n", cap(runeSlice))

	// 修改切片中的内容，字符串是不可以修改的，这里是创建了新的字符串
	runeSlicePart := runeSlice[3:6] // 取出一部分切片
	runeSlicePart[1] = '？'          // 修改第二个元素
	b = string(runeSlice)           // 将修改后的切片转换回新字符串
	fmt.Printf("修改后的新字符串值: %s\n", b)
}
