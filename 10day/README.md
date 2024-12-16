## Go 语言 `map` 笔记

### 1. 什么是 `map`？

- **定义**：`map` 是 Go 语言中的一种内置数据结构，表示无序的键值对集合。它类似于其他编程语言中的字典或哈希表。
- **特点**：
  - 无序：`map` 中的元素没有特定的顺序。
  - 唯一键：`map` 中的键是唯一的，同一个 `map` 中不能有重复的键。
  - 根据键访问值：可以通过键快速地访问到对应的值。

### 2. 声明和初始化

#### 2.1 声明 `map`

```go
var m map[string]int
```

- 声明一个 `map` 变量 `m`，但尚未初始化，`m` 为 `nil`。

#### 2.2 使用 `make` 初始化 `map`

```go
m = make(map[string]int)
```

- 使用 `make` 函数初始化 `map`。可以指定初始容量，但这不是强制的。

#### 2.3 使用字面量初始化 `map`

```go
m := map[string]int{
	"cat":  12,
	"dog":  8,
	"bird": 3,
}
```

- 使用字面量初始化 `map`，并直接赋值。

### 3. 基本操作

#### 3.1 添加和更新元素

```go
m["key"] = value
```

- 如果键 `key` 已经存在，更新其对应的值为 `value`；如果不存在，则插入新的键值对。

#### 3.2 访问元素

```go
value := m["key"]
```

- 通过键访问值，如果键不存在，返回值类型的零值（例如，`int` 的零值是 `0`）。

#### 3.3 删除元素

```go
delete(m, "key")
```

- 删除 `map` 中指定键的键值对。如果键不存在，`delete` 不会产生错误。

#### 3.4 检查键是否存在

```go
value, ok := m["key"]
if ok {
	fmt.Println("Key exists with value:", value)
} else {
	fmt.Println("Key does not exist")
}
```

- `ok` 为 `true` 表示键存在，为 `false` 表示键不存在。

### 4. 遍历 `map`

#### 4.1 使用 `for range`

```go
for key, value := range m {
	fmt.Println(key, value)
}
```

- 遍历 `map` 中的键值对，顺序是无序的。

#### 4.2 只遍历键

```go
for key := range m {
	fmt.Println(key)
}
```

- 只遍历 `map` 的键。

#### 4.3 只遍历值

```go
for _, value := range m {
	fmt.Println(value)
}
```

- 只遍历 `map` 的值。

### 5. 使用 `map` 和 `slice`

#### 5.1 `map` 中的 `slice`

```go
slicesMap := make(map[string][]int)
slicesMap["key"] = []int{1, 2, 3}
```

- 在 `map` 中存储 `slice` 类型的值。

#### 5.2 `slice` 中的 `map`

```go
slices := make([]map[int]string, 2)
slices[0] = make(map[int]string)
slices[0][1] = "value1"
slices[1] = make(map[int]string)
slices[1][2] = "value2"
```

- 在 `slice` 中存储 `map` 类型的值。注意初始化 `slice` 中的每个 `map` 元素。

### 6. `map` 的排序

- Go 语言中的 `map` 是无序的，如果需要排序，通常先将 `map` 的键或值提取到 `slice` 中，然后对 `slice` 进行排序。

#### 6.1 按键排序

```go
keys := make([]string, 0, len(m))
for k := range m {
	keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
	fmt.Println(k, m[k])
}
```

#### 6.2 按值排序

```go
type kv struct {
	Key   string
	Value int
}

var sortedPairs []kv
for k, v := range m {
	sortedPairs = append(sortedPairs, kv{k, v})
}
sort.Slice(sortedPairs, func(i, j int) bool {
	return sortedPairs[i].Value < sortedPairs[j].Value
})
for _, kv := range sortedPairs {
	fmt.Println(kv.Key, kv.Value)
}
```

### 7. 高级用法

#### 7.1 嵌套 `map`

```go
nestedMap := make(map[string]map[string]int)
nestedMap["A"] = make(map[string]int)
nestedMap["A"]["B"] = 1
```

- `map` 中的值可以是另一个 `map`，即嵌套 `map`。

#### 7.2 `map` 的零值

- `nil` 的 `map` 不支持任何操作，包括读取、写入和遍历。初始化后的 `map` 才能进行这些操作。

```go
var m map[string]int
fmt.Println(m == nil) // 输出: true
```