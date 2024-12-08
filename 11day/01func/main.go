package main

import "fmt"

// 函数存在的意义
// 函数的本质是对一段代码的封装
/*
	优点：
		可以封装逻辑代码，可重复调用
		让整体代码结构更易阅读，更简洁
*/
// 函数的定义与使用
/*
	func 函数名(传参列表) (返回值列表) {
		函数体
	}
*/

// 形式1: 正常的函数 传参列表 返回值列表 和 return 函数体
// a b 声明的传参列表变量，c 为声明的返回列表变量
func code1(a int, b int) (c int) {
	// 这个是返回结果 c 的函数体 即代码片段
	return a + b
}

// 对照形式1 ：return 也可以不用在后面加上返回变量
// 声明的返回列表变量 c 可以直接使用，不需要在代码块中再声明
func code1_1(a int, b int) (c int) {
	// 这个是函数体执行代码块
	c = a + b
	// 这个是返回结果 c 可以不写
	return
}

// 形式2: 没有参数列表，没有返回列表
func code2() {
	fmt.Printf("你好世界\n")
}

// 形式3：没有参数但返回固定值
// 可以看到返回列表 int 可以不声明变量名
func code4() int {
	result := 1
	return result
}

// 形式4：有参数列表 没有返回列表
func code3(a string, b string) {
	fmt.Printf("%s%s\n", a, b)
}

// 形式5: 多变量多类型参数列表，多返回列表
// a 是 bool类型参数列表变量 ; b 和 c 是 字符串类型参数列表变量 ; d e 和 f 是 整型类型参数列表变量 ; g h 是 返回整型类型返回列表变量
func code5(a bool, b, c string, d, e, f int) (g, h int) {
	if a {
		fmt.Printf("条件为真，接收到字符串: %s, %s。开始执行加法和减法运算\n", b, c)
		g, h = d+e+f, d-e-f
	} else {
		fmt.Printf("条件为假，接收到字符串: %s, %s。开始执行乘法和除法运算\n", b, c)
		if e != 0 && f != 0 {
			g, h = d*e*f, d/e/f
		} else {
			fmt.Println("错误：e 或 f 为 0，无法进行除法运算")
			g, h = 0, 0
		}
	}

	return g, h
}

// 形式6：任意个数参数列表，有返回列表
// a 布尔类型参数列表； b ...string 即传入多个变量并返回slice[]string类型参数列表；c 为返回字符串类型返回列表
// 其中需要注意的是 b 切片类型的参数列表不能放在首位
func code6(a bool, b ...string) (c string) {
	if a {
		fmt.Printf("条件为真，返回最后一个元素: %v, 类型为: %T\n", b[len(b)-1], b)
		c = b[len(b)-1]
	} else {
		fmt.Printf("条件为假，返回第一个元素: %v, 类型为: %T\n", b[0], b)
		c = b[0]
	}
	return
}

func main() {
	a := 1
	b := 2
	c := code1(a, b)
	d := code1_1(a, b)
	fmt.Printf("%d+%d=%d\n", a, b, c)
	fmt.Printf("%d+%d=%d\n", a, b, d)
	fmt.Println("*****************************")
	code2()
	fmt.Println("*****************************")
	e := "你好"
	f := "世界"
	code3(e, f)
	fmt.Println("*****************************")
	fmt.Printf("额，我也不知道这个无参有返回值有什么作用？%d\n", code4())
	fmt.Println("*****************************")
	g := true
	h, i := "喵喵", "嗷嗷"
	j, k, l := 34, 21, 13 // 必须指定所有变量的值

	// 调用 code5 并接收两个返回值
	m, n := code5(g, h, i, j, k, l)
	fmt.Printf("最终返回：%d,%d\n", m, n)
	g, h, i, j, k, l = false, "呜呜", "嘎嘎", 144, 89, 55
	// 调用 code5 并接收两个返回值
	m, n = code5(g, h, i, j, k, l)
	fmt.Printf("最终返回：%d,%d\n", m, n)
	fmt.Println("*****************************")
	// 调用 code6 传入多个字符串，接受一个返回值
	i = code6(g, h, i, h, i)
	fmt.Printf("接收到的字符串为：%s\n", i)
	g, h, i = true, "喵喵", "嗷嗷"
	i = code6(g, h, i, h, i, i)
	fmt.Printf("接收到的字符串为：%s\n", i)
}
