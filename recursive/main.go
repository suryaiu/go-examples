package main

import "fmt"

// 有n个台阶，一次可以走1步，也可以走2步，问共有多少种走法
func takeSteps(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return takeSteps(n-1) + takeSteps(n-2)
}

func main() {
	fmt.Println(takeSteps(4))
}
