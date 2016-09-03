
// Ticker是一个定时器，周期性地在指定的时间间隔后执行事件。
package main

import "time"
import "fmt"

func main() {

	// 每秒钟触发一次的定时器
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// 3.1秒后停止定时器
	time.Sleep(time.Millisecond * 3100)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
