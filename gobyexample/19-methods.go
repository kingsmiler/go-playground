package main

import "fmt"

type rect struct {
	width, height int
}

// 面积
func (r *rect) area() int {
	return r.width * r.height
}

// 周长
func (r rect) perim() int {
	return 2 * r.width + 2 * r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// 结构体方法调用
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go可根据结构体方法的参数, 自动地在指针与值之间自动地转换,
	// 指针传递可避免值拷贝, 同时可能改变结构体的结构.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
