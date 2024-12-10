package main

import (
	"fmt"
	"sort"
)

// 切片的声明初始化
var (
	a []int
	b [20]float32
)

func main() {
	// 切片本质是对数组的封装抽象
	// 方法一：先声明再初始化
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// 方法二：直接声明初始化
	var str []string = []string{"你好", "我好", "它也好", "好不好", "用了才知道", "差不差", "不用不知道", "知不知道", "不知道"}
	// 方法三：简短声明初始化，自动推断
	chars := []rune{'你', '我', '它', '好', '用', '差', '1', 'a', 'b', 'C', 'W'}
	// 方法四：make函数创建初始化
	// 创建一个长度为5，容量为10的byte类型切片
	bytes := make([]byte, 5, 10)
	// 方法五：以数组为基础创建切片
	b = [20]float32{1: 1.99, 3: 2.98, 5: 3.97, 7: 4.96, 9: 5.95, 11: 6.94, 13: 7.93, 15: 8.92, 17: 9.91, 19: 0.90}
	// 按索引获取切片变量
	// 左闭右开，即包含左边索引对应的值，不包含右边索引对应的值
	newSlice1 := b[0:20]
	// 从头按索引获取切片变量
	newSlice2 := b[:10]
	// 从某个索引到结尾获取切片变量
	newSlice3 := b[9:]
	// 方法六：以切片为基础创建切片
	newSlice4 := newSlice3[5:]

	fmt.Printf("a 初始切片值和类型: %d,%T\n", a, a)
	fmt.Printf("str 初始切片值和类型: %s,%T\n", str, str)
	fmt.Printf("rune 初始切片值和类型: %c,%T\n", chars, chars)
	fmt.Printf("byte 初始切片值和类型: %c,%T\n", bytes, bytes)
	fmt.Printf("b 初始数组值和类型: %f,%T\n", b, b)
	fmt.Printf("newSlice1 以数组 b 为基础的切片值和类型: %.2f,%T\n", newSlice1, newSlice1)
	fmt.Printf("newSlice2 以数组 b 为基础的切片值和类型: %.2f,%T\n", newSlice2, newSlice2)
	fmt.Printf("newSlice3 以数组 b 为基础的切片值和类型: %.2f,%T\n", newSlice3, newSlice3)
	fmt.Printf("newSlice4 以切片 newSlice3 为基础的切片值和类型: %.2f,%T\n", newSlice4, newSlice4)
	fmt.Println("****************************************************")

	// 切片由地址指针&、长度len()、容量cap()组成
	// 切片 newSlice1 newSlice2 newSlice3 newSlice4 都有同一个初始数组 b
	fmt.Printf("b 数组值: %.2f\n\t起始地址: %v\n\t结束地址: %v\n\t长度: %d\n\t容量：%d\n", b, &b[0], &b[len(b)-1], len(b), cap(b))
	// 假设初始数组 b 地址范围是 0x????20~0x????6c 长度是 20 容量是 20
	// 假设切片 newSlice1 是从 b 索引 [0,20) 得到的切片
	// 得到 newSlice1 地址范围是 0x????20~0x????6c 长度是 20-0=>20 容量是 20-0=>20
	// 进而可得到 newSlice1 的地址与 b 地址一致，即 newSlice1 只是对 b 的地址引用
	fmt.Printf("newSlice1 切片值: %.2f\n\t起始地址: %v\n\t结束地址: %v\n\t长度: %d\n\t容量：%d\n", newSlice1, &newSlice1[0], &newSlice1[len(newSlice1)-1], len(newSlice1), cap(newSlice1))
	// 假设切片 newSlice2 是从 b 索引 [0,10) 得到的切片
	// 得到 newSlice2 地址范围是 0x????20~0x????44 长度是 10-0=>10 容量是 20-0=>20
	fmt.Printf("newSlice2 切片值: %.2f\n\t起始地址: %v\n\t结束地址: %v\n\t长度: %d\n\t容量：%d\n", newSlice2, &newSlice2[0], &newSlice2[len(newSlice2)-1], len(newSlice2), cap(newSlice2))
	// 假设切片 newSlice3 是从 b 索引 [9,20) 得到的切片
	// 得到 newSlice3 地址范围是 0x????44~0x????6c 长度是 20-9=>11 容量是 20-9=>11
	// 进而可得到 newSlice3 的容量受到 newSlice3 起始索引和 b 末尾索引范围影响
	fmt.Printf("newSlice3 切片值: %.2f\n\t起始地址: %v\n\t结束地址: %v\n\t长度: %d\n\t容量：%d\n", newSlice3, &newSlice3[0], &newSlice3[len(newSlice3)-1], len(newSlice3), cap(newSlice3))
	// 假设切片 newSlice3 是从 b 索引 [9,20) 得到的切片
	// 假设切片 newSlice4 是从切片 newSlice3 索引 [5,11) 得到的切片
	// 从 b 创建的切片 newSlice3 再次创建切片 newSlice4
	// 得到 newSlice4 切片地址范围是 0x????58~0x????6c 长度是 11-5=>6 容量是 11-5=>6
	// 则可知
	// newSlice4 的容量受到 newSlice4 起始索引和 newSlice3 的末尾索引影响，与 b 没关系
	// 且 newSlice4 引用仍是 b 的地址
	// 切片长度则是自身的起止索引的固有长度
	fmt.Printf("newSlice4 切片值: %.2f\n\t起始地址: %v\n\t结束地址: %v\n\t长度: %d\n\t容量：%d\n", newSlice4, &newSlice4[0], &newSlice4[len(newSlice4)-1], len(newSlice4), cap(newSlice4))

	fmt.Println("***********************************************")

	// 修改切片 newSlice4 元素会影响 newSlice3 和 b
	// 同理拥有共同初始数组的切片，修改元素会互相影响，也进一步证明切片是对数组的地址引用
	newSlice4[2] = 12.54
	fmt.Printf("b 切片值: %.2f\n\t起始地址: %v\n\t结束地址: %v\n\t长度: %d\n\t容量：%d\n", b, &b[0], &b[len(b)-1], len(b), cap(b))
	fmt.Printf("newSlice3 切片值: %.2f\n\t起始地址: %v\n\t结束地址: %v\n\t长度: %d\n\t容量：%d\n", newSlice3, &newSlice3[0], &newSlice3[len(newSlice3)-1], len(newSlice3), cap(newSlice3))
	fmt.Printf("newSlice4 切片值: %.2f\n\t起始地址: %v\n\t结束地址: %v\n\t长度: %d\n\t容量：%d\n", newSlice4, &newSlice4[0], &newSlice4[len(newSlice4)-1], len(newSlice4), cap(newSlice4))
	// append() 追加元素函数, 切片超过自身容量会怎样？
	// 得到 newSlice4 切片地址范围是 0xc000010450~0xc000010468 长度是 7 容量是 12
	// 则可知切片超过固有长度和容量时
	// 会创建新的数组以及引用新数组的地址，原始数组 b 和切片 newSlice3 将不受影响
	// 切片新增元素超过固有容量后，容量会翻倍，但是具体扩容机制请参考 README.md
	newSlice4 = append(newSlice4, 23.5)
	fmt.Printf("newSlice4 追加切片后值: %.2f\n\t起始地址: %v\n\t结束地址: %v\n\t长度: %d\n\t容量：%d\n", newSlice4, &newSlice4[0], &newSlice4[len(newSlice4)-1], len(newSlice4), cap(newSlice4))
	// copy() 拷贝切片内容
	// 创建一个长度为 3 容量为 5 的 float32 类型切片
	newSlice5 := make([]float32, 3, 5)
	// 将 newSlice4 内容拷贝到 newSlice5
	// 如果 newSlice5 的容量和长度都小于 newSlice4 则从 newSlice4 拷贝的数据将丢失，不会报错
	copy(newSlice5, newSlice4)
	fmt.Printf("newSlice5 切片值: %.2f\n\t起始地址: %v\n\t结束地址: %v\n\t长度: %d\n\t容量：%d\n", newSlice5, &newSlice5[0], &newSlice5[len(newSlice5)-1], len(newSlice5), cap(newSlice5))

	newSlice6 := make([]int, 0, 0)
	// 即使一个切片长度容量都是0，但不能说这个切片为空
	// 判断一个切片是否为空，可以使用len()判断
	fmt.Printf("newSlice6 == nil 判断长度容量都为0的切片是否为空: %t\n", newSlice6 == nil)
	fmt.Printf("len(newSlice6) == 0 判断长度容量都为0的切片是否为空: %t\n", len(newSlice6) == 0)

	// ... 表示拆开变量取值
	newSlice7 := make([]float32, 1, 1)
	fmt.Printf("newSlice7 切片值: %.2f\n\t起始地址: %v\n\t结束地址: %v\n\t长度: %d\n\t容量：%d\n", newSlice7, &newSlice7[0], &newSlice7[len(newSlice7)-1], len(newSlice7), cap(newSlice7))
	newSlice7 = append(newSlice7, newSlice1...)
	fmt.Printf("newSlice7 追加后切片值: %.2f\n\t起始地址: %v\n\t结束地址: %v\n\t长度: %d\n\t容量：%d\n", newSlice7, &newSlice7[0], &newSlice7[len(newSlice7)-1], len(newSlice7), cap(newSlice7))
	fmt.Println("***********************************************")

	// 扩容问题
	// newSlice8 int类型长度为 200 容量为 200
	// 正常情况int类型切片达到200 一般来说追加元素后容量翻两倍
	// 算法中的 num = number of elements being added(要追加的元素个数)
	// Existing entries [0, oldLen) are copied over to the new backing store.
	// Added entries [oldLen, newLen) are not initialized by growslice
	// newLen = new length (= oldLen + num)
	// newcap := oldCap
	// doublecap := newcap + newcap
	// const threshold = 256
	// newLen > doublecap 返回 newLen
	// oldCap < threshold 返回 doublecap
	// 都不满足则死循环执行 newcap += (newcap + 3*threshold) >> 2 直到满足条件 uint(newcap) >= uint(newLen) 后 break
	// 由此做一个循环代码
	// 循环 2048 次，每次追加不同数量的元素到切片中，观察切片容量相较于上次的变化
	// 可知长度在 512 范围内变化时，容量扩容后是原始容量的 2 倍
	// 一旦长度超过 512 变化将随着扩容次数变多按照原始容量的约 1.6 倍 1.5 倍 1.4 倍逐级递减
	for i := 0; i < 2048; i++ {
		// 创建一个临时切片，用于存储要追加的元素
		newSlice8 := []int{}

		// 将 [0, 1, 2, ..., i] 范围内的元素添加到 tempSlice 中
		for j := 0; j <= i; j++ {
			newSlice8 = append(newSlice8, j)
		}

		// 打印
		// fmt.Printf("newSlice8 切片值: %d\n\t起止地址范围: %v~%v\n\t长度: %d\n\t容量：%d\n", newSlice8, &newSlice8[0], &newSlice8[len(newSlice8)-1], len(newSlice8), cap(newSlice8))
		fmt.Printf("newSlice8 长度: %d \t容量：%d \t起止地址范围: %v~%v\n", len(newSlice8), cap(newSlice8), &newSlice8[0], &newSlice8[len(newSlice8)-1])

	}

	fmt.Println("*******************************************")
	// 切片赋值问题
	newSlice9 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// %p 格式化输出切片变量地址
	fmt.Printf("newSlice9 \n\t值: %d\n\t长度: %d \t容量：%d \t地址: %p\t起止地址范围: %v~%v\n", newSlice9, len(newSlice9), cap(newSlice9), newSlice9[:], &newSlice9[0], &newSlice9[len(newSlice9)-1])
	for i, v := range newSlice9 {
		fmt.Printf("\t地址 -> 值: %v -> %d\n", &newSlice9[i], v)
	}
	newSlice10 := newSlice9[:]
	fmt.Printf("newSlice10 \n\t值: %d\n\t长度: %d \t容量：%d \t地址: %p\t起止地址范围: %v~%v\n", newSlice10, len(newSlice10), cap(newSlice10), newSlice10, &newSlice10[0], &newSlice10[len(newSlice10)-1])
	for i, v := range newSlice10 {
		fmt.Printf("\t地址 -> 值: %v -> %d\n", &newSlice10[i], v)
	}
	// append(newSlice10[:1], newSlice10[2:]...) 追加后
	// newSlice10 中的第1个索引没有追加，所以 1 这个索引的值将被后续追加的值修改，
	// 影响了之后的每个值都往前一个位置追加赋值，此时也没超过容量，所以底层数组的值被修改了
	// 此时追加 newSlice10 对底层数组 newSlice9 的影响
	// 地址不变，值改变则 newSlice10 -> newSlice9: [9 7 6 5 4 3 2 1 0] -> [9 7 6 5 4 3 2 1 0 0]
	newSlice10 = append(newSlice10[:1], newSlice10[2:]...)
	fmt.Printf("追加后 newSlice10 \n\t值: %d\n\t长度: %d \t容量：%d \t地址: %p\t起止地址范围: %v~%v\n", newSlice10, len(newSlice10), cap(newSlice10), newSlice10, &newSlice10[0], &newSlice10[len(newSlice10)-1])
	for i, v := range newSlice10 {
		fmt.Printf("\t地址 -> 值: %v -> %d\n", &newSlice10[i], v)
	}
	// 此时可以看 newSlice9 效果
	fmt.Printf("追加 newSlice10 对底层数组 newSlice9 的影响\n\t值: %d\n\t长度: %d \t容量：%d \t地址: %p\t起止地址范围: %v~%v\n", newSlice9, len(newSlice9), cap(newSlice9), newSlice9[:], &newSlice9[0], &newSlice9[len(newSlice9)-1])
	for i, v := range newSlice9 {
		fmt.Printf("\t地址 -> 值: %v -> %d\n", &newSlice9[i], v)
	}
	fmt.Println("*******************************************")
	// 排序函数 sort
	// 同理地址不变，值改变则 newSlice10 -> newSlice9: [0 1 2 3 4 5 6 7 9] -> [0 1 2 3 4 5 6 7 9 0]
	sort.Ints(newSlice10)
	// 此时可以看 newSlice9 效果
	fmt.Printf("newSlice10 排序对底层数组 newSlice9 的影响\n\t值: %d\n\t长度: %d \t容量：%d \t地址: %p\t起止地址范围: %v~%v\n", newSlice9, len(newSlice9), cap(newSlice9), newSlice9[:], &newSlice9[0], &newSlice9[len(newSlice9)-1])
	for i, v := range newSlice9 {
		fmt.Printf("\t地址 -> 值: %v -> %d\n", &newSlice9[i], v)
	}
}
