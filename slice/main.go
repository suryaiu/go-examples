package main

import "fmt"

//切片 slice

func main() {
	define()
	fmt.Println("----------------------------------------------------------------------")
	initSlice()
	fmt.Println("----------------------------------------------------------------------")
	createSliceByArray()
	fmt.Println("----------------------------------------------------------------------")
	createSliceByMake()
	fmt.Println("----------------------------------------------------------------------")
	appendSlice()
	fmt.Println("----------------------------------------------------------------------")
	copySlice()
	fmt.Println("----------------------------------------------------------------------")
	deleteEle()
	fmt.Println("----------------------------------------------------------------------")
}

// 定义切片
func define() {
	var s1 []int    // 定义一个存放 int 类型的切片
	var s2 []string // 定义一个存放 string 类型的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // true 空，表示还没有开辟内存空间
}

// 初始化切片
func initSlice() {
	var s1 []int
	var s2 []string
	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"杭州", "上海"}
	fmt.Println(s1, s2)
	// 长度和容量
	fmt.Printf("len(s1): %d cap(s1): %d\n", len(s1), cap(s1))
	fmt.Printf("len(s2): %d cap(s2): %d\n", len(s2), cap(s2))
}

// 由数组创建切片
func createSliceByArray() {
	// 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // 基于一个数组切割，左闭右开
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	s5 := a1[:4]
	s6 := a1[3:]
	s7 := a1[:]
	fmt.Println(s5, s6, s7)
	fmt.Printf("len(s5):%d, cap(s5):%d\n", len(s5), cap(s5))
	fmt.Printf("len(s6):%d, cap(s6):%d\n", len(s6), cap(s6))
	fmt.Printf("len(s7):%d, cap(s7):%d\n", len(s7), cap(s7))
}

// 使用 make 函数创建切片
func createSliceByMake() {
	// 使用 make 函数创建切片
	sl1 := make([]int, 5, 10)
	fmt.Printf("sl1=%v len(sl1)=%v cap(sl1)=%v\n", sl1, len(sl1), cap(sl1))
}

// 切片追加元素
// 注意：切片在追加元素时，只有在容量满了后创建新的内存
func appendSlice() {
	s1 := []string{"beijing", "shanghai", "hangzhou"}
	s1[2] = "杭州"
	fmt.Println(s1)
	fmt.Printf("len(s1)=%d cap(s1)=%d\n", len(s1), cap(s1))
	// 调用append函数必须用原来的切片变量接收返回值
	// append 追加元素，原来的
	s1 = append(s1, "shenzhen")
	fmt.Println(s1)
	fmt.Printf("len(s1)=%d cap(s1)=%d\n", len(s1), cap(s1))
	s1 = append(s1, "guangzhou", "chengdu")
	fmt.Println(s1)
	fmt.Printf("len(s1)=%d cap(s1)=%d\n", len(s1), cap(s1))

	ss := []string{"suzhou", "chongqin", "wuhan"}
	// 将切片ss中的元素追加到切片s1中
	// ss... 三个点表示拆开的意思
	s1 = append(s1, ss...)
	fmt.Println(s1)
	fmt.Printf("len(s1)=%d cap(s1)=%d\n", len(s1), cap(s1))
}

// 切片的copy
func copySlice() {
	s1 := []string{"shanghai", "hangzhou", "beijing"}
	// 赋值，两个变量指向同一块内存地址
	s2 := s1
	var s3 = make([]string, 3, 3)
	// 拷贝
	copy(s3, s1)
	fmt.Println(s1, s2, s3)
	s1[0] = "suzhou"
	fmt.Println(s1, s2, s3)
	// 删除索引为1的元素
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1, s2, s3)
	fmt.Printf("len(s1)=%d cap(s1)=%d\n", len(s1), cap(s1))
}

func deleteEle() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	// 删除下标为2的元素
	// 注意：
	s1 = append(s1[:2], s1[3:]...)
	fmt.Println(s1)
	fmt.Printf("len(s1)=%d, cap(s1)=%d\n", len(s1), cap(s1))

	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a1)
	s2 := a1[:]
	fmt.Println(s2, len(s2), cap(s2))

	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(a1)
	fmt.Println(s2, len(s2), cap(s2))

}
