package main

import "fmt"

// 声明一个数组
var arrayInt [1]int

func main() {
	// 初始化方式1：声明后初始化
	arrayInt = [1]int{1}
	// 初始化方式2：声明并初始化
	var arrayFloat32 [2]float32 = [2]float32{1.99, 2.99}
	// 初始化方式3：简短声明初始化自动推断类型
	arrayComplex64 := [3]complex64{complex(1.99, 2.99), complex(3.99, 4.99), complex(5.99, 6.99)}
	// 初始化方式4：... 动态判断索引数量
	arrayBool := [...]bool{true, false, true, false}
	// 初始化方式5：部分赋值初始化,其余没有赋值则默认值为其类型的零值（比如int默认零值0，bool默认零值false）
	arrayRune := [5]rune{'你', '我'}
	// 初始化方式6：部分索引初始化,根据索引定位赋值
	arrayByte := [6]byte{0: 'A', 5: 'a'}
	fmt.Printf("int数组值和类型打印: %v -> %T\n", arrayInt, arrayInt)
	fmt.Printf("float32数组值和类型打印: %v -> %T\n", arrayFloat32, arrayFloat32)
	// 可以用 %g 或 %v 打印整个复数，或者用 %f 配合 real() 和 imag() 打印复数的实部和虚部
	// %g 是 fmt.Printf 和相关格式化函数中的一种格式化动词，
	// 用于简化打印浮点数和复数。它根据值的大小选择最合适的格式，
	// 通常会打印科学计数法或普通的十进制格式。
	// 浮点数：对于 float32 或 float64 类型的浮点数，
	// %g 会根据值的大小决定是使用科学计数法还是普通的十进制格式，以最简洁的方式显示。
	// 复数：对于 complex64 或 complex128 类型的复数，
	// %g 会打印复数的实部和虚部，通常是以 a+bi 的格式显示。
	fmt.Printf("complex64数组值和类型打印: %g -> %T\n", arrayComplex64, arrayComplex64)
	fmt.Printf("bool数组值和类型打印: %v -> %T\n", arrayBool, arrayBool)
	fmt.Printf("rune数组值和类型打印: %c -> %T\n", arrayRune, arrayRune)
	fmt.Printf("byte数组值和类型打印: %c -> %T\n", arrayByte, arrayByte)
	fmt.Println("***********************************************")
	// 根据索引取值
	fmt.Printf("int类型数组第1个元素: %d\n", arrayInt[0])
	fmt.Printf("float32类型数组第2个元素: %f\n", arrayFloat32[1])
	fmt.Printf("complex64类型数组第3个元素: %v\n", arrayComplex64[2])
	fmt.Printf("bool类型数组第4个元素: %t\n", arrayBool[3])
	fmt.Printf("rune类型数组第5个元素: %c\n", arrayRune[4])
	fmt.Printf("byte类型数组第6个元素: %c\n", arrayByte[5])
	fmt.Println("***********************************************")
	// 根据索引修改值
	arrayInt[0] = 999
	arrayFloat32[1] = 9.99
	arrayComplex64[2] = complex(-9.99, 9.99)
	arrayBool[3] = true
	arrayRune[4] = '它'
	arrayByte[5] = 'J'
	fmt.Printf("为int类型数组第1个元素修改值: %d\n", arrayInt[0])
	fmt.Printf("为float32类型数组第2个元素修改值: %f\n", arrayFloat32[1])
	fmt.Printf("为complex64类型数组第3个元素修改值: %v\n", arrayComplex64[2])
	fmt.Printf("为bool类型数组第4个元素修改值: %t\n", arrayBool[3])
	fmt.Printf("为rune类型数组第5个元素修改值: %c\n", arrayRune[4])
	fmt.Printf("为byte类型数组第6个元素修改值: %c\n", arrayByte[5])
	fmt.Println("***********************************************")
	// 遍历数组
	// 方法一：常规for遍历len()获取数组长度
	fmt.Print("\n遍历int数组元素索引,元素值: ")
	for i := 0; i < len(arrayInt); i++ {
		fmt.Printf("%d,%d\t", i, arrayInt[i])
	}
	fmt.Print("\n遍历float32数组元素索引,元素值: ")
	for i := 0; i < len(arrayFloat32); i++ {
		fmt.Printf("%d,%f\t", i, arrayFloat32[i])
	}
	fmt.Print("\n遍历complex64数组元素索引,元素值: ")
	for i := 0; i < len(arrayComplex64); i++ {
		fmt.Printf("%d,(%.2f+%.2fi)\t", i, real(arrayComplex64[i]), imag(arrayComplex64[i]))
	}
	fmt.Print("\n遍历bool数组元素索引,元素值: ")
	for i := 0; i < len(arrayBool); i++ {
		fmt.Printf("%d,%t\t", i, arrayBool[i])
	}
	fmt.Print("\n遍历rune数组元素索引,元素值: ")
	for i := 0; i < len(arrayRune); i++ {
		fmt.Printf("%d,%c\t", i, arrayRune[i])
	}
	fmt.Print("\n遍历byte数组元素索引,元素值: ")
	for i := 0; i < len(arrayByte); i++ {
		fmt.Printf("%d,%c\t", i, arrayByte[i])
	}
	// 方法2：for range
	fmt.Println("\n***********************************************")
	fmt.Print("\n遍历int数组元素索引,元素值: ")
	for i, v := range arrayInt {
		fmt.Printf("%d,%d\t", i, v)
	}
	fmt.Print("\n遍历float32数组元素索引,元素值: ")
	for i, v := range arrayFloat32 {
		fmt.Printf("%d,%f\t", i, v)
	}
	fmt.Print("\n遍历complex数组元素索引,元素值: ")
	for i, v := range arrayComplex64 {
		fmt.Printf("%d,(%.2f+%.2fi)\t", i, real(v), imag(v))
	}
	fmt.Print("\n遍历bool数组元素索引,元素值: ")
	for i, v := range arrayBool {
		fmt.Printf("%d,%t\t", i, v)
	}
	fmt.Print("\n遍历rune数组元素索引,元素值: ")
	for i, v := range arrayRune {
		fmt.Printf("%d,%c\t", i, v)
	}
	fmt.Print("\n遍历byte数组元素索引,元素值: ")
	for i, v := range arrayByte {
		fmt.Printf("%d,%c\t", i, v)
	}
}
