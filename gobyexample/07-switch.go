package main

import "fmt"
import "time"

// 注意：Go语言中的 switch 没有使用break，与Java不同。
func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 在一个 case 语句中，可使用逗号分隔多个条件。
	// 可使用 default 默认分支。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	// 如果 switch 没有条件语句，则类似于一个ifelse的逻辑结构。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
}

