## Go 语言学习笔记

### 变量作用域

**1. 全局变量和局部变量**

- **全局变量**：在包级别声明的变量，对包内所有函数可见。
- **局部变量**：在函数内部声明的变量，仅对函数内部可见。

**示例：**
```go
package main

import "fmt"

var globalVar = "I am global" // 全局变量

func main() {
    localVar := "I am local" // 局部变量
    fmt.Println(globalVar)   // 访问全局变量
    fmt.Println(localVar)    // 访问局部变量
}
```

**2. 变量隐藏**

- **局部变量会隐藏同名的全局变量**。在局部作用域中，访问的是局部变量。

**示例：**
```go
package main

import "fmt"

var x = "global"

func foo() {
    var x = "local"
    fmt.Println(x) // 输出 "local"
}

func main() {
    foo()
    fmt.Println(x) // 输出 "global"
}
```

**补充：**

```go
// 变量作用域
// 全局变量、局部变量、函数内变量和代码块内变量的优先级从高到低。
// 代码块内变量的作用域仅限于代码块内。
```

### 函数的类型

**1. 函数类型的定义**

- Go 语言中的函数是一种类型，可以定义函数类型并作为参数传递。

**示例：**
```go
package main

import "fmt"
import "reflect"

// 定义函数类型
type IntOperation func(int, int) int

// 实现加法函数
func add(a, b int) int {
    return a + b
}

// 实现减法函数
func subtract(a, b int) int {
    return a - b
}

func performOperation(op IntOperation, a, b int) int {
    return op(a, b)
}

func main() {
    fmt.Println(performOperation(add, 5, 3))       // 输出 8
    fmt.Println(performOperation(subtract, 5, 3))  // 输出 2

    // 打印函数类型
    fmt.Printf("%T\n", add)  // 输出 func(int, int) int
    fmt.Printf("%T\n", performOperation) // 输出 func(IntOperation, int, int) int
}
```

**补充：**

```go
// 函数类型示例
// func(int, int) int 表示一个接收两个 int 类型参数并返回 int 类型的函数
// func() int 表示一个无参数并返回 int 类型的函数
// func(func() int) func(int) int 表示一个接收无参数返回 int 的函数并返回接收 int 参数返回 int 的函数
```

### 匿名函数

**1. 定义和使用**

- 匿名函数是没有名字的函数，通常用于一次性调用或者作为参数传递给其他函数。

**示例：**
```go
package main

import "fmt"

func main() {
    // 定义并立即调用匿名函数
    func(message string) {
        fmt.Println(message)
    }("Hello, Anonymous Function!")

    // 将匿名函数赋值给变量并调用
    greeting := func(name string) string {
        return "Hello, " + name
    }
    fmt.Println(greeting("World"))

    // 获取匿名函数类型
    anonymousFunc := func() { fmt.Println("Hello, this is an anonymous function") }
    fmt.Printf("%T\n", anonymousFunc) // 输出 func()
}
```

**补充：**

```go
// 匿名函数
// 可以使用变量存储匿名函数，或者直接执行匿名函数。
// 使用 reflect 包可以获取匿名函数的类型信息。
// 匿名函数通常用于一次性操作或者作为其他函数的参数。
```

### 闭包

**1. 闭包的定义**

- 闭包是一个函数和它捕获的外部变量的组合。闭包可以访问其外部函数的局部变量。

**示例：**
```go
package main

import "fmt"

// 闭包示例：计数器
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    counter := makeCounter()
    fmt.Println(counter()) // 输出 1
    fmt.Println(counter()) // 输出 2
    fmt.Println(counter()) // 输出 3
}
```

**2. 闭包的应用**

- 闭包可以用于返回一个函数，该函数可以访问定义闭包时的环境变量。

**示例：**
```go
package main

import "fmt"

// 闭包应用之一：构建一个函数并返回
func code3(code2 func(int, int), a, b int) func() {
    fmt.Println("这里是精心构建的闭包函数 code3")
    return func() {
        fmt.Println("这里是闭包函数 code3 返回的匿名函数")
        code2(a, b)
    }
}

// 带有运算目的的函数
func code2(x, y int) {
    fmt.Println("这是个可以用来计算的普通函数 code2")
    fmt.Printf("%d\n", x+y)
}

func main() {
    // 使用闭包来调用不同类型的函数
    code1(code3(code2, 10, 20)) // 传递闭包函数并调用
}

// 普通函数调用闭包
func code1(code func()) {
    fmt.Println("这是一个可以执行传入函数的普通函数 code1")
    code()
}
```

**补充：**

```go
// 闭包
// 闭包是函数和它捕获的外部变量的组合。
// 通过闭包可以保持外部函数的状态。
// 例子：在 makeCounter 和 code3 中，闭包用于返回和传递函数。
// 闭包用于创建状态保持函数以及将不同类型的函数作为参数传递。
```