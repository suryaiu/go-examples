package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 文件操作

func main() {
	// readFromFile()
	readFileByBufio()
}

func readFileByBufio() {
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Printf("open file error: %v\n", err.Error())
		return
	}
	// 关闭文件
	defer file.Close()
	// 创建一个从文件中读取的对象
	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err == io.EOF {
		fmt.Println("读取完毕")
	}
	if err != nil {
		fmt.Printf("read file error: %v\n", err.Error())
		return
	}

	fmt.Println(line)

}

func readFromFile() {
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Printf("open file error: %v\n", err.Error())
		return
	}
	// 关闭文件
	defer file.Close()
	// 共读取的数据长度
	l := 0
	// 读文件
	var temp = make([]byte, 128) // 指定读取的长度
	for {
		n, err := file.Read(temp[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read file error: %v\n", err.Error())
			return
		}
		fmt.Println(string(temp[:n]))
		l += n
	}
	fmt.Printf("读取完毕，攻读去了 %d 个字符", l)
}
