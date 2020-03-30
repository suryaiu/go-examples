package main

import "fmt"

// 指针
func main() {
	// & 取地址
	// * 根据地址取值
	n := 18
	p := &n               // p 是指向n的指针
	fmt.Println(p)        // 输出n的内存地址
	fmt.Printf("%T\n", p) // p 的类型是 *int 表示类型为 int 型指针
	fmt.Println(*p)       // *p 输出指针指向内存上存储的值

	var a *int // 指针的空值为 nil
	// *a = 100   // 对 nil 进行取值操作，报错：runtime error: invalid memory address or nil pointer dereference
	fmt.Println(a) // output: <nil>

	newFunc()
}

// new 是一个内置函数，接受一个参数（是一个类型），返回值是一个指针
// func new(Type) *Type
func newFunc() {
	// 使用 new 来申请一块内存地址
	var a = new(int) // a 是指针
	fmt.Println(a)   // a 所指向的内存地址 output: 0xc0000100e8
	fmt.Println(*a)  // a 所指向的内存上存储的值，int 的零值 0
	fmt.Println(&a)  // 指针 a 的内存地址
}
