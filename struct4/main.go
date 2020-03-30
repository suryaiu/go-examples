package main

import "fmt"

// 使用构造器初始化

type person struct {
	name string
	age  uint8
}

func newPerson(name string, age uint8) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p := newPerson("Tom", 18)
	fmt.Println(*p)
}
