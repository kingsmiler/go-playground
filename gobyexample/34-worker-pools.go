// 线程池模拟
package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan <- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	// 任务通道和结果缓存通道
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 初始化三条线程，由于目前没有任务，将阻塞等待
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 往任务通道发送9个任务
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	// 关闭任务通道，全部任务完成
	close(jobs)

	// 输出任务完成情况
	for a := 1; a <= 9; a++ {
		fmt.Println(<-results)
	}
}
