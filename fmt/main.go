package main

import (
	"fmt"
)

func main() {
	n := 100
	// 查看类型
	fmt.Printf("%T\n", n)
	// 查看变量的值（v万能的）
	fmt.Printf("%#v\n", n)
	// 查看十进制数
	fmt.Printf("%d\n", n)
	// 查看二进数
	fmt.Printf("%b\n", n)
	// 查看八进数
	fmt.Printf("%o\n", n)
	// 查看十六进制数
	fmt.Printf("%x\n", n)

	s := "hello, world"
	// 查看字符串
	fmt.Printf("%s\n", s)
	// v的前面加上#号时，如果是字符串，则输出带引号的字符串
	fmt.Printf("%#v\n", s)

	fmt.Printf("\n")
}
