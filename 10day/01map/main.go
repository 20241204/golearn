package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var (
	a map[string]int
	r *rand.Rand
)

func init() {
	// 创建一个使用当前时间作为种子的局部随机数生成器
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// map 是由内置 hash 散列实现
// map 是无序的，这也意味着循环遍历的key,value取值是随机的
// map 是通过 key, value 的形式存值的
func main() {
	// 初始化方式1：先声明再 make(map) 初始化
	a = make(map[string]int)
	a["cat"] = 12

	// 初始化方式2：简短声明并 make(map,num) 初始化
	// num 可以有也可以没有,只是知道有这个参数就行
	b := make(map[int]string, 1)
	b[1] = "dog"
	fmt.Printf("map 变量 a 的值: %v,类型: %T\n", a, a)
	fmt.Printf("map 变量 b 的值: %v,类型: %T\n", b, b)

	// map 取值方式1：通过key取value
	c := make(map[string]float32, 3)
	c["语文"] = 120.4
	c["英语"] = 130.5
	c["数学"] = 140.3
	fmt.Printf("map 变量 c 的值：%v,类型：%T\n\t语文:%.2f分\n\t英语:%.2f分\n\t数学:%.2f分\n", c, c, c["语文"], c["英语"], c["数学"])

	// map 取值方式2：通过 for rang 遍历取 key,value
	fmt.Println("map 变量 c 的key和value值：")
	for k, v := range c {
		fmt.Printf("\t%s,%.2f\n", k, v)
	}

	// map 取值方式3：通过 for rang 遍历取 key，再通过 key 取 value
	fmt.Println("map 变量 c 的key和key对应的值：")
	for k := range c {
		fmt.Printf("\t%s,%.2f\n", k, c[k])
	}

	// map 取值方式4：通过 for rang 遍历取 value
	fmt.Println("map 变量 c 的 value 值：")
	for _, v := range c {
		fmt.Printf("\t%.2f\n", v)
	}

	// 判断 map 是否存在 key
	value, ok := c["英语"]
	if ok {
		fmt.Printf("map 变量 c 的 key 英语 存在，值为%.2f\n", value)
	} else {
		fmt.Println("map 变量 c 的 key 英语 不存在")

	}
	fmt.Println("********************************")
	// 删除 map 中的值
	delete(c, "语文")
	// 删除一个不存在的 key 也不会报错
	delete(c, "变态")
	fmt.Printf("map 变量 c 删除了 key 语文 后值为%v\n", c)
	fmt.Println("********************************")

	// map 类型的 slice 需要初始化
	d := make([]map[int]float32, 2, 2)
	// 初始化 map 防止 nil 报错
	d[0] = make(map[int]float32, 1)
	d[0][10] = 123.9
	d[1] = make(map[int]float32, 1)
	d[1][100] = 13.9
	fmt.Printf("打印切片变量 d 值:%v,类型:%T\n\t切片索引0对应值:%v,类型:%T\n\t切片索引1对应值:%v,类型:%T\n\t切片索引0对应值的key对应的value:%.2f,类型:%T\n\t切片索引1对应值的key对应的value:%.2f,类型:%T\n", d, d, d[0], d[0], d[1], d[1], d[0][10], d[0][10], d[1][100], d[1][100])
	fmt.Println("********************************")

	// slice 类型的 map 需要初始化
	e := make(map[string][]int, 2)
	fmt.Printf("切片类型的map变量 e 值和类型: %v,%T\n", e, e)
	// 初始化 slice 防止 nil 报错
	e["英语"] = make([]int, 1, 1)
	e["英语"][0] = 100
	fmt.Printf("初始化后切片类型的map变量 e 值和类型: %v,%T\n\tmap 类型切片的值,类型：%d,%T\n\t切片索引0的值,类型：%d,%T\n", e, e, e["英语"], e["英语"], e["英语"][0], e["英语"][0])
	fmt.Println("********************************")

	// map 有序排列只是了解
	// 声明初始化一个 map
	scoreMap := make(map[string]float32, 101)
	// 声明初始化一个只存 map key 的 slice
	keySlice := make([]string, 0, 101)
	for i := 0; i < 101; i++ {
		// 随机生成学生号码
		// [0,100)+1 -> [1,101)
		min := 1
		max := 100
		// 生成 1 到 100 之间的随机浮点数
		randomNum := r.Intn(max-min+1) + min
		// 随机生成分数
		// [0,150.99)+1 -> [1,151.99)
		min = 1
		max = 150
		// 生成 1.00 到 150.99 之间的随机浮点数
		randomScore := float32(min) + r.Float32()*(float32(max)-float32(min)+0.99)
		key := fmt.Sprintf("学生%03d", randomNum)
		scoreMap[key] = randomScore
		keySlice = append(keySlice, key)
	}

	// for range 遍历 slice
	fmt.Println("排序前 map")
	for _, key := range keySlice {
		fmt.Printf("%s,%.2f", key, scoreMap[key])
	}
	// 对 slice 也就是 key 进行排序
	fmt.Println()
	fmt.Println("排序后 map")
	sort.Strings(keySlice)
	// for range 遍历 slice
	for _, key := range keySlice {
		fmt.Printf("%s,%.2f", key, scoreMap[key])
	}
}
