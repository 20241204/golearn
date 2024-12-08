package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

var (
	a byte
	b string
)

func main() {
	// 初始化赋值字符变量 a
	// byte 是 uint8 的别名，
	// byte 主要用于表示 ASCII 字符或二进制数据。与 rune 表示 Unicode 代码点不同
	// byte 仅占用 1 个字节，因此只能表示 0-255 之间的数字（即单个字节)
	a = 'A' // 'A' 对应的 ASCII 码是 65
	fmt.Printf("字符变量 byte 变量 a 值: %c\n", a)
	fmt.Printf("字符变量 byte 变量 a 数值: %d\n", a)
	fmt.Printf("字符变量 byte 变量 a 类型: %T\n", a)
	fmt.Printf("字符变量 byte 变量 a 类型 reflect.TypeOf 返回值: %v\n", reflect.TypeOf(a))
	fmt.Printf("字符变量 byte 变量 a 占用字节大小: %d\n", unsafe.Sizeof(a))
	// byte 的取值范围
	fmt.Printf("字符变量 byte 变量 a 取值范围: %d ~ %d \n", 0, 255)

	// 定义并打印多个字符，由于是单字节无法使用中文字符串
	b = "Hello World. I'm Golang!"
	fmt.Printf("字符串变量 b 值: %s\n", b)

	// 遍历字符串并打印每个字符的 byte 值和 ASCII 码
	fmt.Println("\n遍历字符串并输出每个字符的 byte 值和 ASCII 码:")
	for i := 0; i < len(b); i++ {
		fmt.Printf("字符值: %c\tbyte值: %d\tASCII码: %d\n", b[i], b[i], b[i])
	}

	// 使用切片截取字符串的一部分
	fmt.Println("\n使用切片截取字符串的一部分:")
	slice := b[6:12] // 这里 b[6:12] 截取了字符串 b 的一部分
	fmt.Printf("切片值: %s\n", slice)

	// 将字符串转换为 []byte 切片
	byteSlice := []byte(b)
	fmt.Println("\n将字符串转换为 []byte 切片:")
	fmt.Printf("byte 切片值: %v\n", byteSlice)
	fmt.Printf("byte 切片长度: %d\n", len(byteSlice))
	fmt.Printf("byte 切片容量: %d\n", cap(byteSlice))

	// 修改切片中的内容，字符串是不可以修改的，这里是创建了新的字符串
	byteSlice[1] = 'G'
	b = string(byteSlice) // 将修改后的切片转换回字符串
	fmt.Printf("修改后的新字符串值: %s\n", b)
}
