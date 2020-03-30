package main

import "fmt"

// 闭包

func main() {
	f := f1()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
}

// 可以把这个函数理解成一个类
// n 是类的一个属性
// 返回的匿名函数是类的一个方法
func f1() func(x int) int {
	var n = 0
	return func(x int) int {
		n = n + x
		return n
	}

}
