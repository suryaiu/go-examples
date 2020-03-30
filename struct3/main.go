package main

import (
	"fmt"
	"unsafe"
)

// 结构内的属性占用的是一块连续的内存地址

type person struct {
	name string
	age  uint8
}

func main() {
	p := person{"a", 19}

	fmt.Printf("string size: %d\n", unsafe.Alignof(string("a")))
	fmt.Printf("uint8 size: %d\n", unsafe.Alignof(uint8(19)))

	fmt.Printf("%p\n", &(p.name))
	fmt.Printf("%p\n", &(p.age))
}
