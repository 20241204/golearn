## Go 语言 Slice（切片）深入学习笔记

### 1. 什么是 Slice？

- **Slice** 是 Go 语言中一个非常灵活且功能强大的数据结构，用于表示一个具有相同类型元素的序列。
- **Slice** 是一种引用类型，它对底层数组的一个连续片段进行抽象，包含长度（`len`）和容量（`cap`）两个属性。
- **Slice** 并不直接存储数据，而是对底层数组的引用。因此，对一个切片的修改会影响到底层数组以及其他共享这个底层数组的切片。

### 2. Slice 的创建方式

#### 2.1 使用 `make` 函数

`make` 是 Go 提供的内置函数，用于动态创建切片，指定长度和容量。

```go
s := make([]int, 5, 10) // 创建长度为5，容量为10的int类型切片
```

- 这行代码创建了一个底层数组包含 10 个元素的切片 `s`，但只初始化了前 5 个元素。

#### 2.2 从数组或其他切片中创建

可以通过切片表达式，从一个数组或另一个切片中创建新的切片。

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4] // 创建包含 arr 索引 1 到 3 的新切片
```

- 注意：新切片 `s` 与数组 `arr` 共享相同的底层数组。

#### 2.3 直接声明和初始化

也可以直接使用字面量初始化一个切片。

```go
s := []int{1, 2, 3} // 创建并初始化一个切片
```

- 这种方式下，切片的长度和容量都等于初始化时的元素个数。

### 3. Slice 的基本属性

#### 3.1 长度 (`len`)

切片中元素的个数。

#### 3.2 容量 (`cap`)

从切片的起始位置到底层数组末尾的元素数量。

#### 3.3 示例代码

```go
s := []int{1, 2, 3, 4, 5}
fmt.Println("长度:", len(s)) // 输出: 5
fmt.Println("容量:", cap(s)) // 输出: 5

s = s[:3] // 截取前3个元素
fmt.Println("长度:", len(s)) // 输出: 3
fmt.Println("容量:", cap(s)) // 输出: 5
```

- 通过缩减切片的长度，容量保持不变，因为底层数组未变。

### 4. Slice 的操作

#### 4.1 添加元素 (`append`)

- `append` 函数用于向切片末尾追加元素，当容量不足时，会自动扩容并返回一个新切片。

```go
s := []int{1, 2, 3}
s = append(s, 4, 5) // 在切片末尾添加元素
```

- **注意**：`append` 操作可能会触发底层数组的重新分配。

```plaintext
初始切片: [1, 2, 3]
             ───────
追加后:  [1, 2, 3, 4, 5, _]
```

#### 4.2 拷贝切片 (`copy`)

- `copy` 函数用于将一个切片的内容复制到另一个切片中。

```go
src := []int{1, 2, 3}
dest := make([]int, len(src))
copy(dest, src) // 将 src 内容拷贝到 dest
```

- **注意**：目标切片的长度应不小于源切片，否则只会拷贝部分内容。

#### 4.3 截取切片

- 可以通过 `s[i:j]` 表达式来截取一个新的切片。

```go
s := []int{1, 2, 3, 4, 5}
sub := s[1:4] // 截取索引 1 到索引 3 的元素，sub 为 {2, 3, 4}
```

- 截取后的切片与原切片共享相同的底层数组。

```plaintext
底层数组:  [1, 2, 3, 4, 5]
             ──────────────
切片s:    [2, 3, 4]       
```

### 5. 切片的高级特性

#### 5.1 切片是引用类型

- 切片在函数传递时是引用传递，修改切片会影响到原来的数据。

```go
func modify(s []int) {
    s[0] = 10
}

s := []int{1, 2, 3}
modify(s)
fmt.Println(s) // 输出: [10, 2, 3]
```

- **深刻理解**：所有的切片操作都只是在操作底层数组的不同片段。

#### 5.2 扩容机制

- 当切片的长度增加超过容量时，Go 会自动扩容，新的容量通常是之前的两倍（当切片容量小于 256 时）。

```go
s := []int{1, 2, 3}
s = append(s, 4, 5, 6, 7, 8, 9, 10) // 触发扩容
fmt.Println("长度:", len(s)) // 输出: 10
fmt.Println("容量:", cap(s)) // 输出: 12 (原始容量为 3，扩容后为 12)
```

- **注意**：扩容后，新切片的底层数组是一个全新的数组，旧切片仍然引用原来的底层数组。

### 6. 性能考虑

#### 6.1 减少扩容次数

- 为了减少扩容次数，可以在创建切片时合理估计容量。

```go
s := make([]int, 0, 100) // 预分配容量为 100，避免频繁扩容
```

#### 6.2 扩容行为
- 当切片的长度增加超过容量时，Go 运行时会创建一个新的、更大的底层数组，并将原有元素复制到新数组中。
- 新数组的容量通常是原数组容量的两倍，但具体实现可能有所不同。

```plaintext
初始容量 5: [1, 2, 3, 4, 5]
             ───────────────
扩容到10:  [1, 2, 3, 4, 5, _, _, _, _, _]
```

#### 6.3 扩容算法
- Go 的扩容算法确保了切片在需要扩容时不会浪费内存，同时避免频繁的重新分配。
- go version go1.23.0 windows/amd64 版本的源码部分参考 GOROOT\src\runtime\slice.go

该算法在 `runtime` 包中定义，核心代码如下：

```go
//	newLen = new length (= oldLen + num)
...
// nextslicecap computes the next appropriate slice length.
func nextslicecap(newLen, oldCap int) int {
	newcap := oldCap
	doublecap := newcap + newcap
	if newLen > doublecap {
		return newLen
	}

	const threshold = 256
	if oldCap < threshold {
		return doublecap
	}
	for {
		// Transition from growing 2x for small slices
		// to growing 1.25x for large slices. This formula
		// gives a smooth-ish transition between the two.
		newcap += (newcap + 3*threshold) >> 2

		// We need to check `newcap >= newLen` and whether `newcap` overflowed.
		// newLen is guaranteed to be larger than zero, hence
		// when newcap overflows then `uint(newcap) > uint(newLen)`.
		// This allows to check for both with the same comparison.
		if uint(newcap) >= uint(newLen) {
			break
		}
	}

	// Set newcap to the requested cap when
	// the newcap calculation overflowed.
	if newcap <= 0 {
		return newLen
	}
	return newcap
}
```

**算法说明**：
- **小切片扩容**：如果长度在 512 范围内变化，扩容后的容量时原来容量的 2 倍（`doublecap`）。
- **大切片扩容**：如果长度超过 512 范围外变化，扩容后的容量将随着扩容次数变多按照原来容量的约 1.6 倍 1.5 倍 1.4 倍逐级递减。这个算法在切片容量较大时避免过多的内存浪费。
- **容量溢出处理**：在容量溢出时，算法将容量设置为需要的最小值 `newLen`。

#### 6.4 使用 `copy` 进行深拷贝

- 当你不希望修改切片时影响到其他共享的切片时，可以使用 `copy` 函数进行深拷贝。

```go
original :=

 []int{1, 2, 3}
copy := make([]int, len(original))
copy(copy, original) // 深拷贝
```

### 7. Slice 的零值

- 切片的零值是 `nil`，一个 `nil` 切片的长度和容量都是 0。
- **注意**：和 `nil` 不同的是，一个空切片 `[]int{}` 的长度也是 0，但它有一个有效的底层数组，容量通常为初始化时的大小。

### 8. 注意事项

1. **避免陷入共享数据的陷阱**：
   - 由于切片与底层数组共享数据，因此对切片的修改会影响其他切片或数组。
   - 若需要独立的数据副本，使用 `copy` 函数。

2. **使用 `==` 比较切片**：
   - 切片之间不能用 `==` 运算符比较（`nil` 切片除外）。若要判断切片是否相等，需要比较每个元素。

3. **容量管理**：
   - 合理管理切片容量，以避免频繁扩容带来的性能开销。
   - 在需要大量追加数据时，最好提前估算切片容量，并在创建时通过 `make` 进行分配。

### 9. 总结

- Go 语言的切片提供了非常灵活和高效的方式来处理动态数组，但它也是一种复杂的引用类型，使用时需要注意它的底层数组共享机制、扩容行为以及性能问题。

- **最佳实践**：在创建切片时尽量预估好容量以减少扩容的次数；对外部传递的切片要慎重，确保不会无意中修改底层数据。


### 10. 习题与实践

#### 10.1 基础习题

##### 题目1：创建一个切片，添加 100 个整数，观察容量变化

```go
package main

import "fmt"

func main() {
    var s []int
    for i := 1; i <= 100; i++ {
        s = append(s, i)
        fmt.Printf("添加元素 %d 后，长度: %d, 容量: %d\n", i, len(s), cap(s))
    }
}
```

**答案解析**：
- 这段代码初始化了一个空切片 `s`，然后通过 `append` 函数向其中添加 100 个整数。
- 每次添加一个整数后，代码打印当前的长度和容量。
- 你会看到切片的容量按需扩容，通常是之前容量的两倍。

##### 题目2：练习使用 `copy` 函数实现两个切片之间的深拷贝

```go
package main

import "fmt"

func main() {
    original := []int{1, 2, 3, 4, 5}
    copySlice := make([]int, len(original))

    copy(copySlice, original) // 深拷贝

    fmt.Println("Original slice:", original)
    fmt.Println("Copied slice:", copySlice)

    // 修改 copySlice 不会影响 original
    copySlice[0] = 100
    fmt.Println("After modification:")
    fmt.Println("Original slice:", original)
    fmt.Println("Copied slice:", copySlice)
}
```

**答案解析**：
- 这段代码创建了一个 `original` 切片，并通过 `copy` 函数将其内容深拷贝到 `copySlice`。
- 修改 `copySlice` 后，`original` 不受影响，验证了深拷贝的效果。

#### 10.2 高级习题

##### 题目1：编写一个函数，接收一个切片和一个整数 `n`，返回一个新的切片，包含原切片中所有大于 `n` 的元素

```go
package main

import "fmt"

func filterGreaterThan(s []int, n int) []int {
    var result []int
    for _, v := range s {
        if v > n {
            result = append(result, v)
        }
    }
    return result
}

func main() {
    s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    filtered := filterGreaterThan(s, 5)
    fmt.Println("Filtered slice:", filtered) // 输出: [6 7 8 9 10]
}
```

**答案解析**：
- `filterGreaterThan` 函数遍历输入切片 `s`，将所有大于 `n` 的元素添加到结果切片中。
- 结果切片通过 `append` 函数动态增长，并在函数结束时返回。

##### 题目2：实现一个动态数组，自动扩容并支持各种基本操作（添加、删除、查找）

```go
package main

import "fmt"

type DynamicArray struct {
    data []int
}

func NewDynamicArray() *DynamicArray {
    return &DynamicArray{
        data: make([]int, 0),
    }
}

func (da *DynamicArray) Add(element int) {
    da.data = append(da.data, element)
}

func (da *DynamicArray) Remove(index int) {
    if index < 0 || index >= len(da.data) {
        fmt.Println("Index out of bounds")
        return
    }
    da.data = append(da.data[:index], da.data[index+1:]...)
}

func (da *DynamicArray) Get(index int) int {
    if index < 0 || index >= len(da.data) {
        fmt.Println("Index out of bounds")
        return -1
    }
    return da.data[index]
}

func (da *DynamicArray) Find(element int) int {
    for i, v := range da.data {
        if v == element {
            return i
        }
    }
    return -1
}

func main() {
    da := NewDynamicArray()

    // 添加元素
    da.Add(1)
    da.Add(2)
    da.Add(3)
    fmt.Println("Array:", da.data)

    // 获取元素
    fmt.Println("Element at index 1:", da.Get(1))

    // 查找元素
    fmt.Println("Index of element 3:", da.Find(3))

    // 删除元素
    da.Remove(1)
    fmt.Println("Array after removal:", da.data)
}
```

**答案解析**：
- `DynamicArray` 结构体封装了一个切片，并提供了 `Add`、`Remove`、`Get` 和 `Find` 操作。
- `Add` 方法通过 `append` 函数自动扩容并添加新元素。
- `Remove` 方法通过切片截取删除指定索引的元素。
- `Get` 方法和 `Find` 方法分别用于访问和查找元素。
