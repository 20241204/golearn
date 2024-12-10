
# Go 语言流程控制语句总结1

## `if` 语句

`if` 语句用于条件判断，可以包括初始化语句：

```go
if x := compute(); x > 0 {
    fmt.Println("Positive")
} else {
    fmt.Println("Non-positive")
}
```

- **语法**：
  ```go
  if 条件 {
      // 条件为真时执行的代码
  } else if 另一个条件 {
      // 另一个条件为真时执行的代码
  } else {
      // 所有条件都不满足时执行的代码
  }
  ```

## `for` 循环

Go 语言中只有 `for` 循环，用于迭代：

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// 传统的 while 样式
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}

// 无限循环
for {
    fmt.Println("Infinite loop")
}
```

- **语法**：
  ```go
  for 初始化; 条件; 迭代 {
      // 循环体
  }
  ```