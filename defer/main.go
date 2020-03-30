package main

import "fmt"

func main() {
	// deferDemo()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

// Go 语言中 return 语句不是原子的
// 而是分为两步执行的
// 第一步：为返回值赋值
// 第二步：执行 return 指令
// 在存在 defer 语句的方法中，defer 语句执行retrun指令前执行
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// output: 6
// 第一步：返回值赋值为 x, 此时 x = 5
// defer 将 x 的值修改为 +1, 所以 x =6
// 最后：将 x 的值返回
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

// output: 5
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值 = y = x = 5
}

// output:
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return x
}

func deferDemo() {
	fmt.Println("start")
	// 最后执行
	// defer 把它后面的语句延迟到函数即将返回的时候再执行
	// 先出现的语句放到栈的底部，按照先进先出的顺序延迟执行
	defer fmt.Println("heyheyhey1")
	defer fmt.Println("heyheyhey2")
	defer fmt.Println("heyheyhey3")
	fmt.Println("end")
}
