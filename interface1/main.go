package main

import "fmt"

// 接口
type animal interface {
	speak() // 只有实现该方法的变量，都是实现了该接口
}

type cat struct{}

type dog struct{}

func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func da(a animal) {
	a.speak()
}

func main() {
	var c cat
	var d dog
	da(c)
	da(d)
}
