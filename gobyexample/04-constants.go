// Go 支持字符、字符串、布尔和数字常量。
package main

import (
	"fmt"
	"math"
)

// 常量使用`const`来申明。
const s string = "constant"

func main() {
	fmt.Println(s)

	// 所有使用 'var' 的地方，都可以出现 'const'。
	const n = 500000000

	// 常量表达式。
	const d = 3e20 / n
	fmt.Println(d)

	// 数字常量默认是没有类型的，不过可对它进行显式地类型转换。
	fmt.Println(int64(d))

	// 在上下文环境中使用常量时，常量自动拥有了类型（或者说被隐式地转换）。
	// 比如说，函数 math.Sin 要求的是float64的类型。
	fmt.Println(math.Sin(n))
}
