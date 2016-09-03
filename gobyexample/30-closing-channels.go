// 关闭通道表明不再使用通道发送数据。
package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		// 无限循环，直到任务完成
		for {
			// 第二个变量more默认为true。
			// 如果more为false，表明通道已被关闭，且数据已全部被接收。
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				// 接收完毕
				fmt.Println("received all jobs")
				done <- true
				// 退出循环
				return
			}
		}
	}()

	// 发送3条任务，然后关闭通道。
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// 同步等待，直到接收到数据。
	<-done
}
