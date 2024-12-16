package main

// 导入 fmt 打印工具
import "fmt"

// 变量命名规范(任选其一)
// 小驼峰 aaBbCc
// 大驼峰 AaBbCc
// 单词之间 _ 连接 aa_bb_cc

// 函数外全局变量作用域被所有函数使用
// 函数外只能以关键字开头声明
// 函数外 := 不能使用
// 批量声明变量
var (
	varGlobal string = "hello world!"
)

// main 函数
func main() {
	// 打印 你好世界
	fmt.Println("你好世界！")
	// 变量，可变的值，声明的变量必须被调用，要不然就报错
	// 函数内的局部变量作用域只能在函数内使用
	// 声明初始化方法1 定义并赋值 var 变量名 变量类型 = 变量值
	var name string = "golang"
	// 声明初始化方法2 先声明再赋值
	// var 变量名 变量类型
	// 变量名 = 变量值
	var age, birthday int
	age = 17
	birthday = 2007

	// 声明初始化方法3 类型推导 var 变量名 = 变量值
	var like = "sport"

	// 声明初始化方法4 短变量赋值 变量名 := 变量值
	like1 := "sport"

	// 声明初始化方法5 批量定义变量
	var (
		like2 = "run"
		like3 = "play game"
		like4 = "eat"
	)
	// 全局变量初始化
	varGlobal = "hello,"
	// fmt 打印
	fmt.Println(varGlobal, name, age, birthday, like, like1, like2, like3, like4)
	// 还有一种特殊的定义变量 匿名变量 用 _ 接收参数
	// 匿名变量不占用命名空间，不分配内存，匿名变量可重复
	// _ 多用于占位，表示可忽略
}
