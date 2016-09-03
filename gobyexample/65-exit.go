
// 使用 os.Exit 方法将会导致立即退出程序，
// 状态码由参数指定。
package main

import "fmt"
import "os"

func main() {

	// 如果存在 os.Exit，则相应的 defer 不会被执行
	defer fmt.Println("!")

	// 退出状态码为3
	os.Exit(3)
}

