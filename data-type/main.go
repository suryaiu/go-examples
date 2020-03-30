package main

import (
	"fmt"
)

func main() {
	// 十进制数
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) // 转换成二进制数输出
	fmt.Printf("%o\n", i1) // 转换成八进制数输出
	fmt.Printf("%x\n", i1) // 转换成十六进制数输出
	// 八进制数
	var i2 = 077
	fmt.Printf("%d\n", i2) // 转换成十进制数输出
	// 十六进制数
	var i3 = 0x1ef
	fmt.Printf("%d\n", i3)

	// 查看变量的类型
	fmt.Printf("%T\n", i3)

	var i4 int8 = 9
	fmt.Printf("%d\n", i4)
	i5 := int16(10)
	fmt.Printf("%d\n", i5)
}
