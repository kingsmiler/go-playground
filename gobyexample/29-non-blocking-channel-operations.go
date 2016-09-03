// 使用 select+default 来实现无阻塞操作。
package main

import (
    "log"
)

func main() {
    log.Println("start 1")

    messages := make(chan string)
    signals := make(chan bool)

    // 如果 messages 通道中有数据，则输出数据； 无数据则 default 退出。
    select {
    case msg := <-messages:
        log.Println("received message", msg)
    default:
        log.Println("no message received")
    }

    log.Println("start 2")
    // A non-blocking send works similarly.
    msg := "hi"
    select {
    case messages <- msg:
        log.Println("sent message", msg)
    default:
        log.Println("no message sent")
    }

    log.Println("start 3")
    // 多 case 的使用方式一致。
    select {
    case msg := <-messages:
        log.Println("received message", msg)
    case sig := <-signals:
        log.Println("received signal", sig)
    default:
        log.Println("no activity")
    }

    log.Println("All Done!")
}
