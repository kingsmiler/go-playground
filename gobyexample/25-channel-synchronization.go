// 使用通道来使 goroutine 线程同步的示例。
package main

import "fmt"
import "time"

// 模拟的将运行在线程中的工作任务，
// 如果通道中有数据表明完成当前工作。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 通过往通道中填充数据，表明已完成当前处理流程
	done <- true
}

func main() {

	// 启动一个工作线程
	done := make(chan bool, 1)
	go worker(done)

	// 一直阻塞直到有数据返回
	<-done
}
