package main

import "fmt"

// 匿名函数

func main() {
	// 使用变量接受匿名函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)
	// 如果只需调用一次，可以简写成立即执行
	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)
}
