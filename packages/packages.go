package main

import (
	"fmt"
)

const (
	statusOk = 200
	notFount = 404
)

// 批量声明常量，n2 n3 和上面的一样
const (
	n1 = 100
	n2
	n3
)

const (
	a1 = iota
	a2
	a3
	_
	a4
)

const (
	b1 = iota
	b2 = 100
	b3
	b4
)

const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4
)

// const 中出现iota时置为0，每新增一行就增加1
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

func main() {
	a := 8
	fmt.Println("My favorite number is", a)

	fmt.Println("ok: ", statusOk)
	fmt.Println("notFound:", notFount)

	fmt.Println("n1: ", n1)
	fmt.Println("n2: ", n2)
	fmt.Println("n3: ", n3)

	fmt.Println("a1: ", a1)
	fmt.Println("a2: ", a2)
	fmt.Println("a3: ", a3)
	fmt.Println("a4: ", a4)

	fmt.Println("b1: ", b1)
	fmt.Println("b2: ", b2)
	fmt.Println("b3: ", b3)
	fmt.Println("b4: ", b4)

	fmt.Println("c1: ", c1)
	fmt.Println("c2: ", c2)
	fmt.Println("c3: ", c3)
	fmt.Println("c4: ", c4)

	fmt.Println("d1: ", d1)
	fmt.Println("d2: ", d2)
	fmt.Println("d3: ", d3)
	fmt.Println("d4: ", d4)
}
