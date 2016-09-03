package main

import "fmt"

// 不定长参数个数，参数是同一类型集合
func sum(nums ... int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// slice集合的使用方式为 `func(slice...)`
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
