package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

// 全局声明变量
var (
	a int // 十六进制变量
	b int // 八进制变量
	c int // 二进制变量
	d int // 十进制变量
)

// 整型默认根据平台位数判断
func main() {
	// 初始化十六进制变量
	a = 0xF
	// %o 输出变量值的八进制，加上 0 前缀
	fmt.Printf("十六进制0x%x的八进制数值: 0%o", a, a)
	fmt.Printf("\n十六进制0x%x的十进制数值: %d", a, a)
	// %b 输出变量值的二进制
	fmt.Printf("\n十六进制0x%x的二进制数值: %b", a, a)

	// 初始化八进制变量
	b = 017
	// %x 输出变量值的十六进制，加上 0x 前缀
	fmt.Printf("\n八进制0%o的十六进制数值: 0x%x", b, b)
	fmt.Printf("\n八进制0%o的十进制数值: %d", b, b)
	fmt.Printf("\n八进制0%o的二进制数值: %b", b, b)

	// 初始化二进制变量
	c = 0b1111
	fmt.Printf("\n二进制%b的十六进制数值: 0x%x", c, c)
	fmt.Printf("\n二进制%b的八进制数值: 0%o", c, c)
	fmt.Printf("\n二进制%b的十进制数值: %d", c, c)

	// 初始化十进制变量
	d = 15
	fmt.Printf("\n十进制%d的十六进制表示: 0x%x", d, d)
	fmt.Printf("\n十进制%d的八进制表示: 0%o", d, d)
	fmt.Printf("\n十进制%d的二进制表示: %b", d, d)

	// %T 输出变量的类型
	fmt.Printf("\n十进制%d的类型: %T", d, d)

	// %v 可以用于格式化和打印任何类型的变量, 会智能地选择合适的格式来输出内容
	// 无论是基本类型、结构体、数组、切片, 还是其他复杂类型。
	fmt.Printf("\n十进制%d的reflect.TypeOf()类型返回值: %v", d, reflect.TypeOf(d))

	// 8 bits位 -> 1 Byte字节, 它可以表示 2^8 = 256 种不同的值
	// 对于 int 类型, 在 32 位系统上输出 4, 在 64 位系统上输出 8
	fmt.Printf("\n十进制%d占用Byte字节大小: %d", d, unsafe.Sizeof(d))

	// 这段代码将会输出 int 类型在 64 位系统上可能的最小值和最大值（-9223372036854775808 ~ 9223372036854775807）
	fmt.Printf("\n十进制%d的64位系统数据值允许范围: %d ~ %d", d, math.MinInt64, math.MaxInt64)

	// 这段代码将会输出 int 类型在 32 位系统上可能的最小值和最大值（-2147483648 ~ 2147483647）
	fmt.Printf("\n十进制%d的32位系统数据值允许范围: %d ~ %d", d, math.MinInt32, math.MaxInt32)

	// 将 int 类型或其他整数类型强制转换为 int8 类型
	// int8 类型的值范围是 -128 到 127。如果将一个超出这个范围的整数值转换为 int8，会发生溢出
	// 强制转换时，确保目标类型的范围，以避免意外的值变化
	e := int8(18)
	fmt.Println("\n默认 int 数值强转为 int8 :", e) // 输出: 18
}
