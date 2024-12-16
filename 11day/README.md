### Go语言函数定义与使用笔记

#### 1. 基本函数定义

```go
func 函数名(参数列表) (返回值列表) {
    // 函数体
    return 返回值
}
```

- **参数列表**：包含函数的参数及其数据类型。可以传递零个、一个或多个参数。多个参数之间使用逗号分隔。
- **返回值列表**：指定函数返回值的类型，支持返回零个、一个或多个值。

#### 2. 函数的几种常见形式

1. **有参数，有返回值**

    ```go
    func add(a int, b int) int {
        return a + b
    }
    ```

    - **参数**：`a int, b int`，即函数接收两个整型参数`a`和`b`。
    - **返回值**：函数返回一个整型值，代表`a`与`b`的和。

2. **无参数，无返回值**

    ```go
    func printHello() {
        fmt.Println("Hello, World!")
    }
    ```

    - **说明**：函数没有参数，也没有返回值。它只是执行一些操作，如打印信息。

3. **有参数，无返回值**

    ```go
    func greet(name string) {
        fmt.Printf("Hello, %s!\n", name)
    }
    ```

    - **说明**：函数接收一个字符串参数`name`，但没有返回值。

4. **无参数，有返回值**

    ```go
    func getNumber() int {
        return 42
    }
    ```

    - **说明**：函数不接收任何参数，但返回一个整型值。

5. **多参数，多返回值**

    ```go
    func calculate(a int, b int) (sum int, product int) {
        sum = a + b
        product = a * b
        return
    }
    ```

    - **参数**：`a int, b int`，函数接收两个整型参数。
    - **返回值**：返回两个整型值，分别表示和与积。

6. **可变参数函数**

    ```go
    func printStrings(strs ...string) {
        for _, str := range strs {
            fmt.Println(str)
        }
    }
    ```

    - **说明**：函数接收可变数量的字符串参数，`strs ...string`表示传入的参数被视为一个字符串切片。

#### 3. 参数传递

- **值传递**：函数参数默认是按值传递的，即将参数的副本传递给函数，函数内的修改不会影响外部变量。

    ```go
    func increment(a int) int {
        a = a + 1
        return a
    }
    ```

- **指针传递**：通过传递参数的指针，可以在函数内修改外部变量的值。

    ```go
    func incrementPointer(a *int) {
        *a = *a + 1
    }
    ```

    - **说明**：`a *int`表示传递的是指向整型变量的指针，`*a`表示解引用以访问指针指向的值。

#### 4. 命名返回值

- 可以在函数定义中为返回值命名，这样在函数体内可以直接操作这些变量，最终通过`return`语句返回它们。

    ```go
    func divide(a, b float64) (result float64, err error) {
        if b == 0 {
            err = fmt.Errorf("division by zero")
            return
        }
        result = a / b
        return
    }
    ```

    - **说明**：`result`和`err`是命名返回值，这使得函数体内的代码更清晰，并且无需在`return`时显式列出返回的变量。

#### 5. 匿名函数与闭包

- **匿名函数**：可以定义没有名字的函数并直接调用，通常用于简短的、一次性的功能实现。

    ```go
    func main() {
        func(msg string) {
            fmt.Println(msg)
        }("Hello, Go!")
    }
    ```

- **闭包**：匿名函数可以捕获并使用其外部作用域中的变量，即使该变量的生命周期已结束。

    ```go
    func adder() func(int) int {
        sum := 0
        return func(x int) int {
            sum += x
            return sum
        }
    }
    ```

    - **说明**：`adder`函数返回一个闭包，该闭包持有并操作外部变量`sum`。

#### 6. 递归函数

- 递归函数是指一个函数调用自身。这种方式适用于分解复杂问题，但要注意设置递归终止条件，以避免无限递归。

    ```go
    func factorial(n int) int {
        if n == 0 {
            return 1
        }
        return n * factorial(n-1)
    }
    ```

    - **说明**：`factorial`函数递归地计算阶乘，`n == 0`是递归终止条件。

#### 7. 延迟执行 `defer`

- `defer`语句用于延迟函数的执行，直到外层函数返回时才执行。常用于资源的释放或清理工作。

    ```go
    func readFile(filename string) {
        file, err := os.Open(filename)
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
        // 对文件进行操作
    }
    ```

    - **说明**：`defer file.Close()`确保即使在操作文件时发生错误，文件也会被正确关闭。

#### 8. 函数类型与高阶函数

- **函数类型**：函数在Go语言中也是一种类型，可以赋值给变量或作为参数传递。

    ```go
    func operate(a, b int, op func(int, int) int) int {
        return op(a, b)
    }
    ```

    - **说明**：`op`是一个参数类型为函数的参数，代表执行的操作。

- **高阶函数**：接收函数作为参数，或者返回一个函数的函数称为高阶函数。

    ```go
    func main() {
        add := func(x, y int) int { return x + y }
        result := operate(2, 3, add)
        fmt.Println(result)
    }
    ```

    - **说明**：`operate`是一个高阶函数，它接收一个函数作为参数，并在内部调用它。
