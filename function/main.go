package main

import "fmt"

// 函数

func main() {
	r := sum(2, 3)
	fmt.Println(r)

	i, s := f4()
	fmt.Println(i, s)
	_, v := f4()
	fmt.Println(v)

	f6(1)
	f6(2, 3)
	f6(3, 4, 5, 6)
}

// 返回值命名，相当于在函数声名一个变量
func sum(x int, y int) (res int) {
	res = x + y
	return
}


// 可变长参数
// 注意：第二个参数 y 是一个切片
// 可变长参数必须是最后一个参数
func f6(x int, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

func f5(x, y int) int {
	return x + y
}

// 多个返回值
func f4() (int, string) {
	return 1, "golang"
}

func f1(x int, y int) {
	fmt.Println(x + y)
}

func f2() {
	fmt.Println("f2")
}

func f3() int {
	return 3
}
