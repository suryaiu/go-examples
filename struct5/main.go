package main

import "fmt"

// 方法

type dog struct {
	name string
	age uint8
}

// 构造函数
func newDog(name string, age uint8) *dog {
	return &dog{
		name: name,
		age: age,
	}
}

// 方法是作用于特定类型的函数
// d 表示为调用该方法的变量
// 此处使用的是值接收者
func (d dog) wang() {
	d.age = 4
	fmt.Printf("%s: 汪汪汪...\n", d.name)
}

func (d *dog) eat() {
	d.age = 5
	fmt.Printf("%s: 吃狗粮...\n", d.name)
}


func main() {
	d := *newDog("旺财", 3)
	d.wang()
	fmt.Println(d)
	d.eat()
	fmt.Println(d)
}