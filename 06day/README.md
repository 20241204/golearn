# Golang 运算符笔记

## 1. 算术运算符
用于执行基本的数学运算。

- `+` ：加法运算符。例：`a + b`
- `-` ：减法运算符。例：`a - b`
- `*` ：乘法运算符。例：`a * b`
- `/` ：除法运算符。例：`a / b`
- `%` ：取模运算符（求余）。例：`a % b`

**注意**：除法 `/` 会进行整数除法，结果为整数。如果需要得到浮点数结果，确保操作数至少有一个是浮点数。

## 2. 赋值运算符
用于将右侧的值赋给左侧的变量。

- `=` ：简单赋值。例：`a = b`
- `+=` ：加后赋值。例：`a += b` 等同于 `a = a + b`
- `-=` ：减后赋值。例：`a -= b`
- `*=` ：乘后赋值。例：`a *= b`
- `/=` ：除后赋值。例：`a /= b`
- `%=` ：取模后赋值。例：`a %= b`

## 3. 自增/自减运算符
用于增加或减少变量的值。

- `++` ：自增运算符。将变量的值增加 `1`。例：`a++`
- `--` ：自减运算符。将变量的值减少 `1`。例：`a--`

**注意**：在 Go 语言中，自增和自减运算符只能作为单独的语句使用，不能嵌入到更复杂的表达式中。

## 4. 关系运算符
用于比较两个值，返回布尔值（`true` 或 `false`）。

- `==` ：相等。例：`a == b`
- `!=` ：不相等。例：`a != b`
- `>` ：大于。例：`a > b`
- `<` ：小于。例：`a < b`
- `>=` ：大于等于。例：`a >= b`
- `<=` ：小于等于。例：`a <= b`

## 5. 逻辑运算符
用于组合条件表达式，返回布尔值。

- `&&` ：逻辑与（AND）。例：`a && b`
- `||` ：逻辑或（OR）。例：`a || b`
- `!` ：逻辑非（NOT）。例：`!a`

**注意**：在 `&&` 和 `||` 运算中，Go 语言支持短路求值，即如果左侧表达式已经确定整个表达式的结果，右侧表达式将不再计算。

## 6. 位运算符
用于按位操作整型数据。

- `&` ：按位与（AND）。例：`a & b`
- `|` ：按位或（OR）。例：`a | b`
- `^` ：按位异或（XOR）。例：`a ^ b`
- `&^` ：按位清除（AND NOT）。例：`a &^ b`，清除 `a` 中 `b` 对应位为 `1` 的位。
- `<<` ：左移。将位左移指定数量的位数。例：`a << 2`
- `>>` ：右移。将位右移指定数量的位数。例：`a >> 2`

**注意**：在位运算中，左移 `<<` 和右移 `>>` 运算符用于对二进制位进行操作，通常用来快速乘法和除法。

## 7. 取址运算符与指针运算符

在 Go 语言中，取址运算符（`&`）和指针运算符（`*`）用于操作指针。指针是指向某个变量内存地址的变量。

### 取址运算符 `&`

取址运算符 `&` 用于获取一个变量的内存地址。这个运算符的结果是一个指针，指向该变量的内存地址。

**示例**：
```go
package main

import "fmt"

func main() {
    a := 42
    p := &a // p 是指向变量 a 的指针
    fmt.Println(p)  // 输出变量 a 的内存地址，例如：0xc0000100a0
}
```

**解释**：
- `&a` 返回变量 `a` 的内存地址，并将其赋值给指针变量 `p`。
- `p` 的类型是 `*int`，表示它是一个指向 `int` 类型变量的指针。

### 指针运算符 `*`

指针运算符 `*` 用于间接访问指针指向的值，也称为“解引用”。通过这个运算符，可以访问指针指向的变量的值。

**示例**：
```go
package main

import "fmt"

func main() {
    a := 42
    p := &a     // p 是指向变量 a 的指针
    fmt.Println(*p)  // 输出 42，这是变量 a 的值
    *p = 21     // 通过指针 p 修改变量 a 的值
    fmt.Println(a)   // 输出 21
}
```

**解释**：
- `*p` 表示解引用指针 `p`，即访问指针 `p` 指向的变量 `a` 的值。
- `*p = 21` 通过指针 `p` 修改了变量 `a` 的值，因此 `a` 的值变为 `21`。

### 指针的零值与空指针

在 Go 语言中，指针的零值是 `nil`，表示它不指向任何内存地址。

**示例**：
```go
var p *int
fmt.Println(p) // 输出 <nil>
```

### 指针的比较

指针可以进行比较，两个指针相等当且仅当它们指向同一个变量或都为 `nil`。

**示例**：
```go
var p1, p2 *int
fmt.Println(p1 == p2) // 输出 true，因为两个指针都为 nil
```

### 指针与函数

在 Go 语言中，函数参数传递是值传递，即传递的是值的副本。如果希望在函数内部修改原变量的值，可以通过传递指针实现。

**示例**：
```go
package main

import "fmt"

func increment(n *int) {
    *n++ // 修改指针 n 指向的变量的值
}

func main() {
    a := 42
    increment(&a) // 传递变量 a 的地址
    fmt.Println(a) // 输出 43
}
```

**解释**：
- `increment` 函数接收一个 `*int` 类型的指针，通过解引用 `*n` 来修改 `n` 指向的变量的值。
- `increment(&a)` 传递 `a` 的内存地址给函数，函数内对 `*n` 的修改会影响到变量 `a` 的值。
- 取址运算符 `&` 和指针运算符 `*` 是 Go 语言中操作指针的基本工具。`&` 用于获取变量的内存地址，`*` 用于通过指针访问或修改变量的值。理解和使用指针有助于在 Go 语言中实现高效的内存管理和更灵活的函数参数传递。

## 8. 操作符优先级
在表达式中，运算符有优先级。例如，乘法 `*` 和除法 `/` 的优先级高于加法 `+` 和减法 `-`。你可以使用括号 `()` 来明确运算顺序。

### 优先级顺序（从高到低）：
1. `*`, `/`, `%`, `<<`, `>>`, `&`, `&^`
2. `+`, `-`, `|`, `^`
3. `==`, `!=`, `<`, `<=`, `>`, `>=`
4. `&&`
5. `||`

**示例**：
```go
a := 5 + 3*2  // 结果为 11，因为乘法优先级高于加法
b := (5 + 3) * 2  // 结果为 16，因为括号改变了计算顺序
```

# 注意
**负数位运算**
### 负数运算：位运算的具体流程

假设我们有两个整数 `x = -23` 和 `y = -12`，我们想对它们进行按位与（`&`）、按位或（`|`）、按位异或（`^`）、按位清除（`&^`）、左移（`<<`）、右移（`>>`）等操作。我们将详细展示每个操作的 64 位二进制表示及其结果。

### 步骤1：将整数转换为 64 位二进制表示

在计算机中，整数通常是以二进制形式存储的。为了演示复数运算，我们需要先将两个整数 `x = -23` 和 `y = -12` 转换为 64 位的二进制表示。

- **`x = -23`：**  
  1. 首先，计算 `23` 的二进制表示：`00000000 00000000 00000000 00000000 00000000 00000000 00000000 00010111`。
  2. 取反（求反码）：`11111111 11111111 11111111 11111111 11111111 11111111 11111111 11101000`。
  3. 补码表示（反码加1）：`11111111 11111111 11111111 11111111 11111111 11111111 11111111 11101001`。

- **`y = -12`：**  
  1. 首先，计算 `12` 的二进制表示为：`00000000 00000000 00000000 00000000 00000000 00000000 00000000 00001100`。
  2. 取反（求反码）：`11111111 11111111 11111111 11111111 11111111 11111111 11111111 11110011`。
  3. 补码表示（反码加1）：`11111111 11111111 11111111 11111111 11111111 11111111 11111111 11110100`。

### 步骤2：进行二进制运算

现在我们可以对 `x` 和 `y` 进行各种二进制运算。

1. **按位与运算 `x & y`**：
   ```
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11101001  (-23)
   &
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11110100  (-12)
   -----------------------------------------------------
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11100000  (-32)
   ```

2. **按位或运算 `x | y`**：
   ```
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11101001  (-23)
   |
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11110100  (-12)
   -----------------------------------------------------
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11111101  (-3)
   ```

3. **按位异或运算 `x ^ y`**：
   ```
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11101001  (-23)
   ^
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11110100  (-12)
   -----------------------------------------------------
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00011101  (29)
   ```

4. **按位清除运算 `x &^ y`**：
   ```
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11101001  (-23)
   &^
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11110100  (-12)
   -----------------------------------------------------
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00001001  (9)
   ```

5. **左移运算 `x << 2`**：
   ```
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11101001  (-23)
   <<
   2
   -----------------------------------------------------
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 10100100  (-92)
   ```

6. **右移运算 `y >> 2`**：
   ```
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11110100  (-12)
   >>
   2
   -----------------------------------------------------
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11111101  (-3)
   ```

### 步骤3：将二进制结果转换为十进制

#### 步骤3-1：判断是否为负数
- 如果是 `0`，那么该数是正数。
  **按位异或运算 `x ^ y`结果**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00011101  (29)
   ```

   则可以直接得到 `2^4 + 2^3 + 2^2 + 2^0` = `29`

  **按位清除运算 `x &^ y`结果**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00001001  (9)
   ```

   则可以直接得到 `2^3 + 2^0` = `9`

- 如果二进制的最高位（即最左侧位）是 `1`，那么该数是负数。
- 理解负数在计算机中的二进制表示需要了解补码的概念。
- 负数在计算机中通常使用补码表示，而补码的计算方法是：
  **将数的绝对值的二进制表示取反（求反码），然后加1**
- 可以通过逆过程将二进制补码转换回十进制负数。

#### 步骤3-2：求补码对应的反码
- 反码是将二进制的每一位取反（即 `0` 变 `1`，`1` 变 `0`）。

  **按位与运算 `x & y`结果**：

   ```
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11100000  (-32)
   ```

  **按位或运算 `x | y`结果**：

   ```
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11111101  (-3)
   ```

  **左移运算 `x << 2`结果**：

   ```
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 10100100  (-92)
   ```

  **右移运算 `y >> 2`结果**：

   ```
   11111111 11111111 11111111 11111111 11111111 11111111 11111111 11111010  (-3)
   ```
 
- 取反后得到反码：

  **按位与运算 `x & y`结果反码**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00011111
   ```

  **按位或运算 `x | y`结果反码**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000010
   ```

  **左移运算 `x << 2`结果反码**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 01011011
   ```

  **右移运算 `y >> 2`结果反码**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000101
   ```

#### 步骤3-3：反码加1得到原码（即该负数的绝对值的二进制表示）
- 在反码的基础上加 1：

  **按位与运算 `x & y`结果反码+1**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00011111
   +                                                              00000001
   ---------------------------------------------------------
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00100000
   ```

  **按位或运算 `x | y`结果反码+1**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000010
   +                                                              00000001
   ---------------------------------------------------------
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000011
   ```

  **左移运算 `x << 2`结果反码+1**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 01011011
   +                                                              00000001
   ---------------------------------------------------------
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 01011100
   ```

  **右移运算 `y >> 2`结果反码+1**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000101
   +                                                              00000001
   ---------------------------------------------------------
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000110
   ```

- 得到原码

  **按位与运算 `x & y`结果原码**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00100000
   ```

  **按位或运算 `x | y`结果原码**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000011
   ```

  **左移运算 `x << 2`结果原码**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 01011100
   ```

  **右移运算 `y >> 2`结果原码**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000110
   ```

#### 步骤3-4：将原码转换为十进制

  **按位与运算 `x & y`结果原码**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00100000
   ```
   对应的十进制值是 `2^5` -> `32`

  **按位或运算 `x | y`结果原码**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000011
   ```
   对应的十进制值是 `2^1 + 2^0` -> `3`

  **左移运算 `x << 2`结果原码**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 01011100
   ```
   对应的十进制值是 `2^6 + 2^4 + 2^3 + 2^2` -> `92`

  **右移运算 `y >> 2`结果原码**：

   ```
   00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000011
   ```
   对应的十进制值是 `2^1 + 2^0` -> `3`


#### 步骤3-5：添加负号
- 因为最初确定该为负数

  **按位与运算 `x & y`结果**
  `-32`

  **按位或运算 `x | y`结果**
  `-3`

  **左移运算 `x << 2`结果**
  `-92`

  **右移运算 `y >> 2`结果**
  `-3`

#### 最终结果 
  **按位与运算 `x & y`结果**：`-32`

  **按位或运算 `x | y`结果**：`-3`

  **按位异或运算 `x ^ y`结果**：`29`

  **按位清除运算 `x &^ y`结果**：`9`

  **左移运算 `x << 2`结果**：`-92`

  **右移运算 `y >> 2`结果**：`-3`

### 总结

通过上述过程，你可以看到如何将负数和正数转换为 64 位二进制表示，并进行各种位运算。负数的位运算与正数有一些不同，主要体现在补码的使用上。