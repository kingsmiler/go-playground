// 有方向（读取或者接收）的通道示例，如果没有 <- 则表示双向
package main

import "fmt"

// 在参数中，如果 <- 在 chan 的右边，表示通道只能接收数据，如pings
func ping(pings chan <- string, msg string) {
	fmt.Println(msg)
	pings <- msg
}

// 在参数中，如果 <- 在 chan 的左边，表示通道只能读取数据，如pings
func pong(pings <-chan string, pongs chan <- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
