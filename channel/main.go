package main

import "fmt"

var b chan int // 需要指定管道中元素的类型

func main() {
	fmt.Println(b)         // nil
	b = make(chan int)     // 没有缓冲区的初始化
	b = make(chan int, 16) // 通道必须使用make初始化，才能使用 带缓冲区的初始化
	fmt.Println(b)
	// 把数据发送到通道中
	b <- 3
	fmt.Println(<-b) // 从通道中取值
}
