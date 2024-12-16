package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

var (
	a string
	b string
)

func main() {
	// 初始化赋值字符串使用 "" 或 `` 括起来
	// "" 中的内容支持转义，不会保留原有格式
	a = "|你好你好世界|"
	// `` 中的内容不支持转义，保留格式内容
	b = `|你好你好
世界|`

	fmt.Printf("变量 a 值: %s\n", a)
	fmt.Printf("变量 b 值: %s\n", b)
	fmt.Printf("变量 a 类型: %T\n", a)
	fmt.Printf("变量 b 类型: %T\n", b)
	fmt.Printf("变量 a 类型 reflect.TypeOf 返回值: %v\n", reflect.TypeOf(a))
	fmt.Printf("变量 b 类型 reflect.TypeOf 返回值: %v\n", reflect.TypeOf(b))
	// 输出字符串的字节大小实际内存占用量可以通过 len() 函数获得，它返回字符串的字节长度
	// unsafe.Sizeof() 返回的是字符串结构体的大小（通常是 16 字节，包含指针和长度），而不是字符串内容的大小
	// 获取实际字符串字节 len
	fmt.Printf("变量 a 实际占用字节: %d, 字符串类型结构体占用字节大小: %d\n", len(a), unsafe.Sizeof(a))
	fmt.Printf("变量 b 实际占用字节: %d, 字符串类型结构体占用字节大小: %d\n", len(b), unsafe.Sizeof(b))
	// 字符串类型为复合类型没有取值范围，因此不需要类似的输出
	fmt.Println("变量 a 和 b 的取值范围不可定义。")
	fmt.Println("**********************")

	// 字符串变量相关常用转义功能
	// 使用 \\ 转义 \
	fmt.Println("|\\|转义效果")
	fmt.Println("**********************")

	// 使用 \" 转义 "
	fmt.Println("|\"|转义效果")
	fmt.Println("**********************")

	// 使用 \t 实现 4个空格 效果
	fmt.Println("四个空格|\t|转义效果")
	fmt.Println("**********************")

	// 使用 \r 实现 光标回到行头 效果
	fmt.Println("光标回到行头|\r|转义效果")
	fmt.Println("**********************")

	// 使用 \n 实现换行效果
	fmt.Println("换行|\n|转义效果")
	fmt.Println("**********************")

	// %#v 输出变量值并根据值类型将转义字符和标点符号一起输出
	fmt.Printf("变量 a 带转义和标点符号值: %#v\n", a)
	fmt.Println("**********************")
	fmt.Printf("变量 b 带转义和标点符号值: %#v\n", b)
	fmt.Println("**********************")

	// 字符串变量相关常用函数
	// 拼接字符串 +
	a = a + "+++" + b
	fmt.Println("字符串变量 + 拼接 a b :", a)
	fmt.Println("**********************")
	// 还原变量 a
	a = "|你好你好世界|"

	// 拼接字符串 fmt.Sprintf
	a = fmt.Sprintf("%s+++%s", a, b)
	fmt.Println("字符串变量 fmt.Sprintf 拼接 a b :", a)
	fmt.Println("**********************")

	// 切分 strings.Split 与 加入 strings.Join
	a = "|ambmcmdmemfmgmhmimjmkmlm|"
	b = `|z.x.v.f.n.b.r.w.w.a.f.b
	s.f.w.f.f.v.d.s.s.c.s.
	|`
	// 根据 m 切分字符串变量 a
	c := strings.Split(a, "m")
	fmt.Printf("%s 根据 m 切分数据值: %v \n", a, c)
	fmt.Printf("%s 根据 m 切分数据带转义和标点符号值: %#v \n", a, c)
	fmt.Println("**********************")

	// 根据 ! 加入到切分后变量
	d := strings.Join(c, "!")
	fmt.Printf("%s 根据 m 加入到切分后变量值: %v \n", a, d)
	fmt.Printf("%s 根据 m 加入到切分后变量带转义和标点符号值: %#v \n", a, d)
	fmt.Println("**********************")

	// 根据 . 切分字符串变量 b
	c = strings.Split(b, ".")
	fmt.Printf("%s 根据 . 切分数据值: %v \n", b, c)
	fmt.Printf("%s 根据 . 切分数据带转义和标点符号值: %#v \n", b, c)
	fmt.Println("**********************")

	// 根据 ~ 加入到切分后变量
	d = strings.Join(c, "~")
	fmt.Printf("%s 根据 ~ 加入到切分后变量值: %v \n", b, d)
	fmt.Printf("%s 根据 ~ 加入到切分后变量带转义和标点符号值: %#v \n", b, d)
	fmt.Println("**********************")

	// 判断全部字符串之中存在部分字符串
	// 判断在变量 a 中存在 "bm"
	e := strings.Contains(a, "bm")
	fmt.Printf("判断在 %s 中存在 bm : %t\n", a, e)
	fmt.Println("**********************")

	// 判断在变量 e 中存在 "a.f"
	e = strings.Contains(b, "a.f")
	fmt.Printf("判断在 %s 中存在 a.f : %t\n", b, e)
	fmt.Println("**********************")

	// 判断前后缀 strings.HasPrefix strings.HasSuffix
	// 判断字符串中存在前缀 "|amb"
	f := strings.HasPrefix(a, "|amb")
	fmt.Printf("判断在 %s 中存在前缀 |amb : %t\n", a, f)
	fmt.Println("**********************")

	/*
		判断字符串中存在后缀 c.s.
		|
	*/
	f = strings.HasSuffix(b, `c.s.
	|`)
	fmt.Printf("判断在 %s 中存在后缀 c.s.\n\t| : %t\n", b, f)
	fmt.Println("**********************")

	// 获取字符串索引 strings.Index 最后一次出现的索引值 strings.LastIndex
	// 获取在字符串中 fmg 的索引值
	g := strings.Index(a, "fmg")
	fmt.Printf("判断在 %s 中 fmg 的索引值 : %d\n", a, g)
	fmt.Println("**********************")
	g = strings.Index(b, "s.")
	fmt.Printf("判断在 %s 中 s. 最后一次出现的索引值 : %d\n", b, g)
	fmt.Println("**********************")
}
