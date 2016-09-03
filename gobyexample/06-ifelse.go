package main

import "fmt"

func main() {
	// if 格式，条件语句没有圆括号。
	if 8 % 4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// ifelse格式，注意，条件语句没有圆括号。
	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 在判断条件前的语句中定义的变量，在后面的所有分支中都可使用
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if count := 4; num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
		fmt.Println(count, "is negative")
	}
}