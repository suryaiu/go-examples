package main

import "fmt"

func main() {
	// 数组的声名
	var a [3]int
	var b [4]int
	// a = b // 错误，a 和 b 是不同的类型，因为数组的长度是数组的类型组成的一部分
	// 数组的初始化
	// 默认元素都是 零值(bool: false, 整形和浮点数都是0，字符串: "")
	fmt.Println(a, b)
	// 输出类型 a: [3]int, b: [4]int
	fmt.Printf("a: %T, b: %T\n", a, b)
	// 1. 初始化方式1
	a = [3]int{1, 2, 3}
	// 2. 初始化方式	根据初始值自动推断数组的长度
	a2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a2)
	// 3. 初始化方式 	根据索引来初始化
	a3 := [5]int{1: 1, 2: 5}
	fmt.Println(a3)

	// 1. 使用 普通for方式遍历数组
	citys := [...]string{"杭州", "北京", "上海", "香港"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	// 2 使用 for range 遍历数组
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 多维数组
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	// 多维数组的遍历
	for _, v := range a11 {
		fmt.Println(v)
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Print(b1, b2)
}
