// 在Go语言中，变量必须显式地申明后才能使用，
// 并且没有被使用的变量是不允许存在的，否则编译将报错。

package main

import "fmt"

func main() {
	// `var` 申明一个变量并赋值，显式指明数据类型。
	var a string = "initial"
	fmt.Println(a)

	// 一次申明多个变量并赋值，显式指明数据类型。
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 也可以根据初始化数据推断数据类型。
	var d = true
	fmt.Println(d)

	// 申明的变量如果没有指定初始值，将使用默认值。
	var e int
	fmt.Println(e)

	// 使用 ':=' 快速申明变量与赋值。
	var f1 string = "short"
	f2 := "short"
	fmt.Println(f1, f2)
}
