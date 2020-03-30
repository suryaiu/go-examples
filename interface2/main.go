package main

import "fmt"

// 接口
type car interface {
	run()
	didi()
}

type benchi struct {
	brand string
}

type oooo struct {
	brand string
}

// benchi 实现接口中的 run 方法
func (b benchi) run() {
	fmt.Printf("%s 嗖的一下跑了...\n", b.brand)
}

// benchi 实现接口中的 didi 方法
func (b benchi) didi() {
	fmt.Printf("%s 嘀嘀嘀...\n", b.brand)
}

// oooo 实现接口中的 run 方法
func (o oooo) run() {
	fmt.Printf("%s 也嗖的一下跑了...\n", o.brand)
}

// oooo 实现接口中的 didi 方法
func (o oooo) didi() {
	fmt.Printf("%s 也嘀嘀嘀...\n", o.brand)
}

func drive(c car) {
	c.run()
	c.didi()
}

func main() {
	b := benchi{"奔驰"}
	o := oooo{"奥迪"}

	drive(b)
	drive(o)
}
