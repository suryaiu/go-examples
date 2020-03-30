package main

import "fmt"

var x = 100

func f1() {
	x := 10
	// 函数中查找变量的顺序
	// 1. 先在函数中查找
	// 2. 找不到就到函数外面找，一直找到全局
	fmt.Println(x)
}

func main() {
	f1()
}
