package main

import "fmt"

func main() {

	// 在数组或 Slice 集合上使用 range 时，返回的参数有两个；
	// 第一个是元素的序号，第二个是元素的值；
	// 如果不需要序号，可使用空标识符_来代替。
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 输出序号
	for i, num := range nums {
		fmt.Println("index:", i, "; value=", num)
	}

	// Range在map对象上返回key-value对。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 在字符串上使用range时，第一个参数为序号，第二个为unicode值。
	for i, c := range "AB" {
		fmt.Println(i, c)
	}

	// 不管range前没有接收参数,都将循环数组长度的次数
	count := 1
	for range "AB" {
		count++
	}
	fmt.Println(count)
}
