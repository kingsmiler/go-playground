// In a [previous](range) example we saw how `for` and
// `range` provide iteration over basic data structures.
// We can also use this syntax to iterate over
// values received from a channel.

package main

import "fmt"

func main() {

	// 长度为2的缓冲通道。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 循环输出通道中的数据；如果通道未关闭，则阻塞等待新数据。
	for elem := range queue {
		fmt.Println(elem)
	}
}
