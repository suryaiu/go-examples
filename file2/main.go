package main

import (
	"fmt"
	"os"
)

// 文件写操作

func main() {
	writeFile()
}

func writeFile() {
	fileObj, err := os.OpenFile("a.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
	}
	// 方法一：write
	fileObj.Write([]byte("Hello, world."))
	// 方法二：writeString
	fileObj.WriteString("Hello, goland.")

	defer fileObj.Close()
}
