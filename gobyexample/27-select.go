// 通过 select 来等待处理不同通道上不同处理时间的线程。
// 在 Go 语言中， goroutine 和 select 是一个强力的组合。
package main

import "time"
import "fmt"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// 模拟在不同线程中，各自的处理时间不同的场景
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "two"
	}()

	// 通过使用 `select` 来同步全部数据到达，
	// 如果还有通道的数据未返回，则会一直阻塞。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	fmt.Println("Done!")
}
