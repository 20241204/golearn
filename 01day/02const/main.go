package main

// 导入 fmt 打印工具
import "fmt"

// main 函数
func main() {
	// 常量: 不可改变的值
	// 声明初始化方法1 const 常量名 常量类型 = 常量值
	const name string = "golang"
	// 声明初始化方法2 const 常量名 = 常量值
	const age = 19
	// 声明初始化方法3 批量定义常量
	const (
		birthday int = 2009
		like         = "sport"
		like2        = "run"
		like3        = "play game"
		like4        = "eat"
	)
	// 声明初始化方法4 批量定义常量同等值
	const (
		like5 = "unknown" // unknown
		like6             // unknown
		like7             // unknown
	)
	// 打印常量
	fmt.Println(name, age, birthday, like, like2, like3, like4, like5, like6, like7)
	// 声明初始化方法5 iota 常量计数器 + 批量声明初始化常量
	// iota 常量计数器只能在常量表达式中使用
	// 第一个 iota 起始值是 0 此后常量表达式内没增加一行就默认 +1
	const (
		like8  = iota // 0
		like9  = iota // 1
		like10 = iota // 2
	)
	// 打印常量
	fmt.Println(like8, like9, like10)
	// 声明初始化方法6 iota 常量计数器 + 批量声明初始化常量
	// 不会影响计数
	const (
		like11 = iota // 0
		like12        // 1
		like13        // 2
	)
	// 打印常量
	fmt.Println(like11, like12, like13)
	// 声明初始化方法7 iota 常量计数器 + 批量声明初始化常量
	// _ 可以占位丢弃 iota 计数
	const (
		like14 = iota // 0
		_             // 1
		like15        // 2
	)
	// 打印常量
	fmt.Println(like14, like15)
	// 声明初始化方法8 iota 常量计数器 + 批量声明初始化常量
	// iota 新增一行常量声明初始化计数 +1
	// 非开头赋予 iota 正常计数不会被其他常量影响计数
	const (
		like16 = iota // 0
		like17 = 101  // 101
		like18 = iota // 2
		like19        // 3
	)
	// 打印常量
	fmt.Println(like16, like17, like18, like19)
	// 声明初始化方法9 iota 常量计数器 + 批量声明初始化常量
	// iota 新增一行常量声明初始化计数 +1
	// 常量行数不增加则不计数
	const (
		like20, like21 = iota + 0, iota + 1 // like20: 0+0 like21: 0+1
		like22, like23 = iota + 2, iota + 3 // like22: 1+2 like23: 1+3
	)
	// 打印常量
	fmt.Println(like20, like21, like22, like23)
	// 声明初始化方法10 iota 常量计数器 + 批量声明初始化常量
	// 定义数量级 可以这样理解，具体的细节：
	// 左移操作 `1 << (10 * iota)` 是将数字 1 左移 `10 * iota` 位，等于 `2` 的 `10 * iota` 次方。例如，`1 << 10` 是 `2` 的 10 次方，即 `1024`。
	// 这样一来，`1 << (10 * iota)` 就等于 `1024` 的 `iota` 次方。例如：
	// 当 `iota` 为 1 时，`1 << 10` 相当于 `1024^1`，即 `1024`。
	// 当 `iota` 为 2 时，`1 << 20` 相当于 `1024^2`，即 `1048576`。
	// 当 `iota` 为 3 时，`1 << 30` 相当于 `1024^3`，即 `1073741824`。
	// 所以，简化说，`1 << (10 * iota)` 实际上是 `1024` 的 `iota` 次方。
	const (
		_      = iota             // 0
		like24 = 1 << (10 * iota) // 1 左移运算 10 * 1 相当于 1024^1
		like25 = 1 << (10 * iota) // 1 左移运算 10 * 2 相当于 1024^1
		like26 = 1 << (10 * iota) // 1 左移运算 10 * 3 相当于 1024^1
		like27 = 1 << (10 * iota) // 1 左移运算 10 * 4 相当于 1024^1
	)
	// 打印常量
	fmt.Println(like24, like25, like26, like27)

}
