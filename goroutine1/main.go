package main

import (
	"time"
	"fmt"
)

func hello(i int) {
	fmt.Println("hello,", i)
}
func main() {
	// 开启一个单独的goroutine去执行hello
	for i := 0; i < 100; i++ {
		go hello(i)
	}
	for i := 0; i < 100; i++ {
		fmt.Println("main, ", i)
	}
	time.Sleep((time.Second))
}