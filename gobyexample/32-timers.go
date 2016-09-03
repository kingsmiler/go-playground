// 计时器
package main

import "time"
import "fmt"

func main() {
	// 2秒的计时器
	timer1 := time.NewTimer(time.Second * 2)

	// <-timer1.C 阻塞操作，直到C收到数据表示超时
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// 如果仅用于等待，应当使用 time.Sleep。
	// 计时器的另一特性是可以在超时之前被停止的。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
