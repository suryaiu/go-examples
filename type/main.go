package main

import "fmt"

// 自定义类型
type myInt int

// 类型别名
type yourInt = int

func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	var m yourInt
	m = 120
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	// rune 是 int32 的别名
	var c rune
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	
}
