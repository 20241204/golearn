package main

import (
	"fmt"
	"sort"
	"strings"
)

var (
	s = "Okay, I understand. If you frequently need to look up keys by their values in an actual project, it might be worth considering redesigning the data structure or adding a reverse mapping. I hope this information is helpful to you. Feel free to ask me if you have any questions! Keep up the good work!" // 目标文本
)

/*
练习1: 统计单词出现次数
"Okay, I understand. If you frequently need to look up keys by their values in an actual project, it might be worth considering redesigning the data structure or adding a reverse mapping. I hope this information is helpful to you. Feel free to ask me if you have any questions! Keep up the good work!"
打印格式 单词:数字
*/
func code1() { // 笨方法
	// 全文小写
	s = strings.ToLower(s)
	// 去除标点符号
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, "!", "", -1)
	//fmt.Printf("%v\n", s)
	// 以空格为分割点分割文本返回单词字符串数组
	a := strings.Split(s, " ")
	// fmt.Printf("%v,%d\n", a, len(a))
	// 创建 map 容量为整个数组长度
	b := make(map[int]string, len(a))
	// 遍历数组存放到 map
	for i, v := range a {
		b[i] = v
	}
	// 将map长度作为循环终止
	endLen := len(b)
	// fmt.Printf("%v,%d\n", b, endLen)
	sMap := make(map[string]int)     // 最终统计
	for k1 := 0; k1 < endLen; k1++ { // 外层循环 map 用于主判断
		// 初始化计数器
		num := 0
		// 判断 map key 是否存在
		value, ok := b[k1]
		if ok {
			// key 存在自增 ++ 就是把自己算作统计中的1个
			num++
			// k2:=k1+1 是希望第二组map不要从第一组的 k1 头开始对比,这样会重复计数导致不准确
			for k2 := k1 + 1; k2 < endLen; k2++ { // 内曾循环 map 作为对照组判断
				// 判断对照组 map key 存不存在
				value1, ok1 := b[k2]
				if ok1 { // 存在就对比
					//fmt.Printf("%s=%s?\t", wordStr, value1)
					if value == value1 { // 两组 map 不同 key 对应的值一不一样?
						num++         // 相同自增1
						delete(b, k2) // 删除第二组map中与第一组相同 value 的元素省的重复计数
					}
				}
			}
			// fmt.Printf("\n%s:%d\n", value, num)
			sMap[value] = num // map 存入最终统计
			// 删除第一组已经对比过的value
			delete(b, k1)
		}

	}
	// 根据 map 最终统计长度初始化 slice 长度和容量
	sSlice := make([]string, len(sMap), len(sMap))
	nSlice := 0
	for key := range sMap {
		sSlice[nSlice] = key // slice 存入 key
		// 自增索引
		nSlice++
	}

	// fmt.Printf("\n%v\n%v\n", sMap, sSlice)
	sort.Strings(sSlice)       // 对 slice key 排序
	for _, k := range sSlice { // 按 key 的顺序输出
		fmt.Printf("%s:%d\n", k, sMap[k])
	}
}

func code2() { // 简化方法

	// 将所有单词转换为小写，并通过空格分割单词
	words := strings.Fields(strings.ToLower(s))

	// 创建一个 map 来存储单词及其出现的次数
	wordCount := make(map[string]int)

	// 遍历单词并统计出现次数
	for _, word := range words {
		// 去除标点符号
		word = strings.Trim(word, ",.!?")
		// 借用零点值0统计，并用++自增
		wordCount[word]++
	}

	// 打印结果
	// for word, count := range wordCount {
	// 	fmt.Printf("%s:%d\n", word, count)
	// }

	// 根据 map 最终统计长度初始化 slice 长度和容量
	sSlice := make([]string, len(wordCount), len(wordCount))
	nSlice := 0
	for key := range wordCount {
		sSlice[nSlice] = key // slice 存入 key
		// 自增索引
		nSlice++
	}

	// fmt.Printf("\n%v\n%v\n", sMap, sSlice)
	sort.Strings(sSlice)       // 对 slice key 排序
	for _, k := range sSlice { // 按 key 的顺序输出
		fmt.Printf("%s:%d\n", k, wordCount[k])
	}

}

func main() {
	code1()
	fmt.Println("**************************************")
	code2()
}
