package main

import "fmt"

// 函数名后先是入参，后是出参。该函数返回两个出参。
func vals() (int, int) {
	return 3, 7
}

func main() {

	// 接收返回的多个参数
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 忽略部分返回参数
	_, c := vals()
	fmt.Println(c)
}
