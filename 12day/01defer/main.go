package main

import "fmt"

// defer 是延迟执行代码块或函数的关键字
// defer 修饰的代码块会等待整体代码块执行完再延迟执行
// defer 修饰的代码块会按照 先进后出，后进先出的顺序执行
// 可以用在代码执行结束时理相关资源的释放以及处理
/*
  defer 在 return 中使用要注意，return 是非原子(非原子表明是可分割步骤的)执行
  即 return 会执行两步操作
   1. 赋值变量
   2. return 返回变量
  defer 会在中间阶段执行
   1. 赋值变量
   2. defer 代码块
   3. return 返回变量
*/
func code1() { // defer 会等待整个代码块执行完后减缓执行并按照先进后出，后进先出的顺序执行
	// 未被修饰代码，按照顺序执行 1 3 5
	fmt.Println("执行顺序1")
	// defer 修饰的代码 延迟执行 先进后出，后进先出 4 2
	defer fmt.Println("defer修饰1执行顺序2")
	fmt.Println("执行顺序3")
	defer fmt.Println("defer修饰2执行顺序4")
	fmt.Println("执行顺序5")
}

func code2() (x int) { // defer 和 return的关系很特别 defer 会在 return 两步变量处理之间执行
	a := 10 // a = 10
	defer func() {
		fmt.Printf("return 赋值后 x = %d -> ", x)
		a++ // a = 10 + 1 -> 11 a 参与运算但此时 x 不参与运算
		fmt.Printf("defer 执行后 x = %d -> ", x)
	}()
	/*
	   三步执行
	   1. 变量赋值 x = a
	   2. defer执行 a++
	   3. 返回变量 return x
	*/
	return a
}

func code3() (x int) { // defer 和 return的关系很特别 defer 会在 return 两步变量处理之间执行
	defer func() { // 这个是匿名函数，可以在函数内定义比较特别的函数
		// x = 10
		fmt.Printf("return 赋值后 x = %d -> ", x)
		x++ // x = 10 + 1 -> 11 此时 x 参与运算了
		fmt.Printf("defer 执行后 x = %d -> ", x)
	}()
	/*
	   三步执行
	   1. 变量赋值 x = 10
	   2. defer执行 x++
	   3. 返回变量 return x
	*/
	return 10
}

func code4() (x int) { // defer 和 return的关系很特别 defer 会在 return 两步变量处理之间执行
	defer func(x int) {
		// 匿名函数是在 defer 被解释时捕获参数的。当时的命名返回值 x 仍然是 0，所以匿名函数内的 x 初始化为 0
		fmt.Printf("return 赋值后 x = %d -> ", x)
		x++ // x = 0 + 1 -> 1 此时匿名函数内的 x 参与运算了
		fmt.Printf("defer 执行后 x = %d -> ", x)
	}(x) // x = 10 值传递是 x 的副本
	/*
	   三步执行
	   1. 变量赋值 x = 10
	   2. defer执行 x++
	   3. 返回变量 return x
	*/
	return 10
}

func code5() (x int) { // defer 和 return的关系很特别 defer 会在 return 两步变量处理之间执行
	defer func(x *int) { //这次传*int类型指针
		// 匿名函数是在 defer 被解释时捕获参数的。当时的命名返回地址 x 即匿名函数内的 x 指针变量为地址
		fmt.Printf("return 赋值后地址 x = %v -> ", x)
		(*x)++ // x = 10 + 1 -> 11 此时 x 指针变量的值参与运算了
		fmt.Printf("defer 执行后地址 x = %v -> ", x)
	}(&x) // x = 10 取x地址传递是 x 的地址
	/*
	   三步执行
	   1. 变量赋值 x = 10
	   2. defer执行 x++
	   3. 返回变量 return x
	*/
	return 10
}

func code6(s string, b, c int) int {
	d := b + c
	fmt.Printf("执行(s,b,c)->%s,%d,%d ; return -> %d\n", s, b, c, d)
	return d
}

func main() {
	code1()
	a := code2()
	fmt.Printf("defer 修饰代码块后 return 结果1: %d\n", a)
	a = code3()
	fmt.Printf("defer 修饰代码块后 return 结果2: %d\n", a)
	a = code4()
	fmt.Printf("defer 修饰代码块后 return 结果3: %d\n", a)
	a = code5()
	fmt.Printf("defer 修饰代码块后 return 结果4: %d\n", a)
	/*
		defer 修饰的是外部的函数 code6 会延迟执行,
		内部的函数 code6 会按照正常的顺序调用执行
			b = 1
			c = 2
			内部函数 code6("内一",b,c) -> 内一,1,2 ; return b+c -> 3
			此时外部函数没执行但是已经传值
			外部函数 code6("外一",b,c) -> 外一,1,3
			b = 0
			内部函数 code6("内二",b,c) -> 内二,0,2 ; return b+c -> 2
			此时外部函数没执行但是已经传值
			外部函数 code6("外二",b,c) -> 外二,0,2
			c = 1
		直到全部执行完，defer才会起作用此时
		外部的函数 code6 会按照先进后出，后进先出的顺序执行
			外部函数执行 code6("外二",0,2) -> 外二,0,2 ; return b+c -> 2
			外部函数执行 code6("外一",1,3) -> 外一,1,3 ; return b+c -> 4
	*/
	b := 1
	c := 2
	defer code6("外一", b, code6("内一", b, c))
	b = 0
	defer code6("外二", b, code6("内二", b, c))
	c = 1
}
