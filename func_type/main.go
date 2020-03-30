package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("hello, golang")
}

func f2() int {
	return 10
}

// 函数作为参数来传递
func f3(f func() int) {
	res := f()
	fmt.Println(res)
}

func f4(x, y int) int {
	return x + y
}

// 函数还可以作为返回值
// f5 的参数类型是 func() int
// f5 的返回值类型是 func(int, int) int
func f5(x func() int) func(int, int) int {
	return ff
}

func ff(a, b int) int {
	return a + b
}

func main() {
	// f1 不带括号不是函数调用，而是获取函数的类型
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)

	f3(f2)

	fff := f5(f2)
	fmt.Printf("%T\n", fff)
	fmt.Println(fff(2, 0))
}
