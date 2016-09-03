package main

import "fmt"

// 返回一个匿名闭包函数，闭包函数能持久变量
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// 接收闭包函数
	nextInt := intSeq()

	// 调用闭包函数
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 使用新的闭包函数后，闭包所包含的值被重新初始化
	newInts := intSeq()
	fmt.Println(newInts())
}