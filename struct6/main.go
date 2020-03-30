package main

import (
	"fmt"
)

// 结构体匿名字段
// 不能出现类型相同的字段
// 可以通过.类型来调用属性

type dog struct {
	string
	uint8
}

// 结构体嵌套
type person struct {
	name    string
	age     uint8
	addr    address // 结构体嵌套
	address         // 匿名嵌套结构体
}

type address struct {
	province string
	city     string
}

func main() {
	d := dog{
		"Tom",
		19,
	}
	fmt.Println(d)
	fmt.Println(d.string)
	fmt.Println(d.uint8)
	fmt.Println("---------------")
	p := person{
		name: "Tom",
		age:  19,
		addr: address{
			province: "浙江",
			city:     "杭州",
		},
		address: address{ // 匿名嵌套结构体赋值
			province: "江苏",
			city:     "苏州",
		},
	}
	fmt.Println(p)
	fmt.Println(p.addr.city)
	// 匿名嵌套结构体	现在自己的结构体中找city字段，
	// 找不到再去匿名嵌套结构体中找city字段
	fmt.Println(p.province + ", " + p.city)
}
