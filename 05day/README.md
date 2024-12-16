# Go 语言流程控制语句总结2

## `switch` 语句

`switch` 语句用于多分支选择，每个分支无需显式使用 `break` 语句，因为 Go 会自动终止当前分支的执行。

```go
switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
default:
    fmt.Println("Other day")
}
```

- **语法**：
  ```go
  switch 表达式 {
  case 值1:
      // 执行代码
  case 值2:
      // 执行代码
  default:
      // 执行代码
  }
  ```

- **类型选择**：
  `switch` 语句也可以用于类型选择，特别是在处理接口类型时非常有用：
  ```go
  switch v := interfaceValue.(type) {
  case int:
      fmt.Println("int", v)
  case string:
      fmt.Println("string", v)
  default:
      fmt.Println("unknown", v)
  }
  ```

## `goto` 语句

`goto` 语句用于无条件跳转到程序中的另一部分。尽管 `goto` 能用于实现复杂的流程控制，但应尽量避免使用，以保持代码的可读性。

```go
func main() {
    i := 0
LOOP:
    fmt.Println(i)
    i++
    if i < 5 {
        goto LOOP
    }
}
```

- **语法**：
  ```go
  goto 标签
  ...
  标签:
      // 执行的代码
  ```

- **注意**：
  - 使用 `goto` 可能会让代码变得难以理解，因此应慎重使用。

## `continue` 语句

`continue` 语句用于跳过当前循环中的剩余语句，直接进入下一次循环。

```go
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if j == 3 {
            continue // 跳过 j == 3 时的输出
        }
        fmt.Printf("i: %d, j: %d\n", i, j)
    }
}
```

- **语法**：
  ```go
  continue
  ```

- **特点**：
  - 在多层循环中，`continue` 只会影响它所在的那个循环。

## `break` 语句

`break` 语句用于终止当前的循环或 `switch` 语句。

### 1. 用于终止单层循环

```go
for i := 0; i < 5; i++ {
    if i == 3 {
        break // 跳出循环，当 i == 3 时
    }
    fmt.Println(i)
}
```

### 2. 用于跳出多层循环（使用标签）

```go
outerLoop:
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if j == 3 {
            break outerLoop // 跳出整个循环
        }
        fmt.Printf("i: %d, j: %d\n", i, j)
    }
}
```

- **语法**：
  ```go
  break [标签]
  ```

- **特点**：
  - 在带标签的 `break` 语句中，可以终止多个循环。

## `defer` 语句

`defer` 语句用于在函数退出前执行延迟调用的函数。通常用于资源释放，如关闭文件、解锁互斥量等。

```go
func main() {
    defer fmt.Println("World") // 延迟执行
    fmt.Println("Hello")
}
```

- **语法**：
  ```go
  defer 函数调用
  ```

- **特点**：
  - `defer` 语句的执行顺序是“后进先出”（LIFO），即最后一个 `defer` 语句会最先执行。
