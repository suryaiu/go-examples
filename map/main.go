package main

import "fmt"

func main() {
	defineMap()
}

func defineMap() {
	var m1 map[string]int
	fmt.Println(m1)
	m1 = make(map[string]int, 10)
	m1["tom"] = 18
	m1["amy"] = 19
	fmt.Println(m1)
	fmt.Println(m1["amy"])
	// 获取map中不存在的值时返回零值 0
	fmt.Println(m1["mike"])
	v, ok := m1["mike"]
	if !ok {
		fmt.Println("不存在该值")
	} else {
		fmt.Println(v)
	}

	// 遍历map
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 只遍历 key 值
	for k := range m1 {
		fmt.Println(k)
	}

	// 只遍历 value
	for _, v := range m1 {
		fmt.Println(v)
	}

	// 删除元素
	delete(m1, "tom")
	fmt.Println(m1)
	// 删除一个不存在的值，不会报错
	delete(m1, "tom")
	fmt.Println(m1)
	fmt.Println("--------------------------------------")
	mapInSlice()
}

func mapInSlice() {
	var s1 = make([]map[int]string, 10, 10)
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "Tom"
	fmt.Println(s1)
}
