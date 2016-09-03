
// panic 表示出现了未知异常，通常将它用于快速失败，
// 正常操作中不应当出现错误的地方。

package main

import "os"

func main() {

	// 出现异常，直接退出程序。
	// panic("a problem")

	// 检测创建文件是否错误，如果错误则退出
	_, err := os.Create("gopanic.txt")
	if err != nil {
		panic(err)
	}
}
