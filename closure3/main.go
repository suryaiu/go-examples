package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		// 如果 name 没有指定的后缀，则加上，否则就返回
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	// 返回闭包
	f := makeSuffix(".png")
	fmt.Println("处理后的文件命：", f("summer"))
	fmt.Println("处理后的文件命：", f("winter.png"))
}
