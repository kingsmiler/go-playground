package main

import "fmt"

// 定义一个名为 person 的结构体
type person struct {
	name string
	age  int
}

func changeName(p *person)  {
	p.name = "abc"
}

func main() {
	// 初始化对象，按照结构体成员变量的顺序，依次传入参数。
	fmt.Println(person{"Bob", 20})

	// 初始化时，通过指定成员变量的方式来初始化。
	fmt.Println(person{name: "Alice", age: 30})

	// 初始化时，未指定的成员变量，将采用默认值。
	fmt.Println(person{name: "Fred"})

	// 通过 `&` 前缀将结构体实例标识为引用类型。
	fmt.Println(&person{name: "Ann", age: 40})

	// 通过点号来访问结构体属性字段。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 对引用类型，也是点号操作。
	sp := &s
	fmt.Println(sp.age)

	// 结构体数据是可变的。
	sp.age = 51
	fmt.Println(sp.age)

	changeName(sp)
	fmt.Println(sp)
	fmt.Println(s)
}
