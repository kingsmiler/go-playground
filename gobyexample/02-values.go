// Go语言包含有大量的值类型，包括字符串，整数，浮点数，布尔类型，等。
// 下面是一个小小的基本示例。
package main

import "fmt"

func main() {

	// Strings, which can be added together with `+`.
	fmt.Println("go" + "lang")

	// Integers and floats.
	fmt.Println("1+1 =", 1 + 1)
	fmt.Println("7.0/3.0 =", 7.0 / 3.0)

	// Booleans, with boolean operators as you'd expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
