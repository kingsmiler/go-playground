package main

import (
    "time"
    "log"
)

func main() {
    // 通道的创建语法 `make(chan val-type)`
    messages := make(chan string)

    log.Println("start1")

    // 向通道发送数据的语法为： `channel <-`
    go func() {
        time.Sleep(time.Second * 5)
        messages <- "ping"
    }()

    select {
    case msg := <-messages:
        log.Println("received message", msg)
    default:
        log.Println("no message received")
    }

    log.Println("start2")

    // 从通道中读取数据的语法为： `<-channel`， 这是同步阻塞的方式。
    msg := <-messages
    log.Println(msg)

    log.Println("Done!")
}
