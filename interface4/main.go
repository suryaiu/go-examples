package main

import "fmt"

// 空接口

// 空接口作为参数，表示可以接受任何类型
func show(a interface{}) {
	fmt.Printf("type: %T value:%v\n", a, a)
}
func main() {
	// 空接口类型可以存入任何类型数据
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "Tom"
	m1["age"] = 20
	m1["merried"] = false
	m1["hobby"] = [...]string{"coding", "sporting"}

	fmt.Println(m1)

	show(2)
	show(true)
	show('a')

}
