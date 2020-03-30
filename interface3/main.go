package main

import "fmt"

// 接口, 使用指针接收者实现接口中的方法
type car interface {
	didi()
}

type benchi struct {
	brand string
}

type oooo struct {
	brand string
}

// benchi 实现接口中的 didi 方法, 使用指针接收者
func (b *benchi) didi() {
	fmt.Printf("%s 嘀嘀嘀...\n", b.brand)
}

// oooo 实现接口中的 didi 方法, 使用指针接收者
func (o *oooo) didi() {
	fmt.Printf("%s 也嘀嘀嘀...\n", o.brand)
}

func drive(c car) {
	c.didi()
}

func main() {
	// 获取地址
	b := &benchi{"奔驰"}
	o := &oooo{"奥迪"}

	drive(b)
	drive(o)
}
