package main

import "fmt"

// 结构体是值类型

type person struct {
	name, gender string
	age          uint8
}

// Go 语言中函数传参永远都是值类型的，即创建副本
// 结构体是值类型的，所以此处修改的是p的一个副本
func f1(p person) {
	p.age = 20
}

// 使用指针作为参数，可以解决参数拷贝的问题
func f2(p *person) {
	p.age = 25
	// (*p).age = 25 和上面等价的，上面会自动去找对应的变量
}

func main() {
	var p person
	p.name = "老王"
	p.gender = "男"
	p.age = 30

	f1(p)
	fmt.Println(p)

	f2(&p)
	fmt.Println(p)

	// 使用 new 函数创建 p，new函数返回的是一个地址
	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)  // p2 中保存的值的内存地址
	fmt.Printf("%p\n", &p2) // p2 指针的内存地址

	p2.age = 22
	p2.name = "Mike"

	fmt.Println(p2)

	fmt.Println("------------------")
	// 直接使用列表初始化
	var p3 = person{"Michael", "male", 33}
	fmt.Println(p3)
	fmt.Println("------------------")
	// 使用 key-value 方式初始化
	var p4 = person{
		name:   "Linda",
		gender: "famle",
	}
	fmt.Println(p4)

}
