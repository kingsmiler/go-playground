package main

import "fmt"

func main() {
	// 通过内置函数make来定义slice数组，元素类型与初始长度是必须的参数
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// slice数组元素的赋值方式与普通数组相同
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// 通过len获取实际长度
	fmt.Println("len:", len(s))

	// 通过append来增加元素
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 拷贝数组
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 通过slice[low:high)来获取数组子集
	// 本例得到第3、第4、第5三个元素组成的新数组
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 本例得到前5个元素组成的数组
	l = s[:5]
	fmt.Println("sl2:", l)

	// 本例得到从第3个元素到结尾全部元素组成的数组
	l = s[2:]
	fmt.Println("sl3:", l)

	// 定义的同时初始化
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 多维数组
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	aa := []string{}
	fmt.Println(aa)
}
