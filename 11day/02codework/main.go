package main

import "fmt"

// 回文判断
// 一段字符串，从右往左读==从左往右读
// 12345654321
// 上海自来水来自水上
// 山西运煤车煤运西山
// 黄山落叶松叶落山黄
/*
	思路：
	先将字符串转换字符类型切片，这样才能获取到正确的长度
	接下来就是对比切片索引对应的值是否一致
	上 sSlice[0] == sSlice[len(s)-1]
	海 sSlice[1] == sSlice[len(s)-2]
	自 sSlice[2] == sSlice[len(s)-3]
	来 sSlice[3] == sSlice[len(s)-4]
  	...
	c  sSlice[i] == sSlice[len(s)-1-i]
	经过对比发现，其实只要对比一般的字符串即可完成回文的确认 即 len(s)/2
*/
// checkString 判断字符串是否为回文
func checkString(s string) string {
	// 打印字符串的值、类型和长度
	fmt.Printf("%s -> %T -> 长度: %d\n", s, s, len(s))

	// 将字符串转换为 rune 类型的切片，处理多字节字符
	sSlice := []rune(s)
	fmt.Printf("%v -> %T -> 长度: %d\n", sSlice, sSlice, len(sSlice))

	// 比较字符串是否为回文
	for i := 0; i < len(sSlice)/2; i++ {
		fmt.Printf("%d == %d?\n", sSlice[i], sSlice[len(sSlice)-i-1])
		if sSlice[i] != sSlice[len(sSlice)-i-1] {
			return "不是回文"
		}
	}
	return "是回文"
}

func main() {
	// 测试字符串
	s0 := "12345654321"
	s1 := "上海自来水来自水上"
	s2 := "山西运煤车煤运西山"
	s3 := "黄山落叶松叶落山黄"
	fmt.Printf("%s %s\n", s0, checkString(s0))
	fmt.Printf("%s %s\n", s1, checkString(s1))
	fmt.Printf("%s %s\n", s2, checkString(s2))
	fmt.Printf("%s %s\n", s3, checkString(s3))
}
