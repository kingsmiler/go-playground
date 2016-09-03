package main

import "fmt"

// Map 是Go语言的内置类型，在其它语言中，它被称为 Hash 或者 Dict 。
func main() {

	// 语法格式为 `make(map[key-type]val-type)`
	m := make(map[string]int)

	// 通过 `name[key] = val` 语法赋值
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// 取值 `name[key]`
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// 通过 len 函数计算个数
	fmt.Println("len:", len(m))

	// 通过 delete 删除某属性
	delete(m, "k2")
	fmt.Println("map:", m)

	// 取值语法返回的第二个参数可表明，指定的key是否在Maps中。
	// 对于需要忽略的参数，可使用空标识符_来占位。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 使用map关键字初始化数据
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	x := map[string]int{}
	fmt.Println("map:", x)
}
