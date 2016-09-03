package main

import "fmt"

func main() {
	// 定义一个长度为5的整形数组变量a，
	// 数组元素的类型与数组长度都是数组定义的组成部分。
	var a [5]int
	fmt.Println("emp:", a)

	// 通过下标序号为元素赋值
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 通过内置函数len获取数组的长度
	fmt.Println("len:", len(a))

	// 定义数组的同时进行初始化
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
