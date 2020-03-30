package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 文件写操作

func main() {
	// writeFile()
	// wirteFileByBufio()
	writerFileByIoutil()
}

func writeFile() {
	fileObj, err := os.OpenFile("a.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	// 方法一：write
	fileObj.Write([]byte("Hello, world.\n"))
	// 方法二：writeString
	fileObj.WriteString("Hello, goland.\n")

	defer fileObj.Close()
}

func wirteFileByBufio() {
	fileObj, err := os.OpenFile("a.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	writer := bufio.NewWriter((fileObj))
	for i := 0; i < 10; i++ {
		writer.WriteString("Hello, world, by bufio.\n")
	}
	// 将缓存中的内存写入到文件
	writer.Flush()
	defer fileObj.Close()
}

func writerFileByIoutil() {
	str := "hello, golang by ioutil"
	err := ioutil.WriteFile("a.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
