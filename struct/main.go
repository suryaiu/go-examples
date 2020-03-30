package main

import (
	"fmt"
)

// 结构体
type student struct {
	age   uint16
	name  string
	hobby []string
}

func main() {
	var tom student
	tom.age = 10
	tom.name = "Tom"
	tom.hobby = []string{"篮球", "读书"}

	fmt.Println(tom)

	// 匿名结构体
	var s struct {
		name string
		age  int
	}

	s.name = "Amy"
	s.age = 19

	fmt.Printf("%T %v\n", s, s)
}
