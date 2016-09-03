package main

import "fmt"

func main() {
    // 通过指定第二个参数，来配置通道的缓冲数据条数
    messages := make(chan string, 2)

    // 填充数据
    messages <- "buffered"
    messages <- "channel"

    // 输出通道中的数据
    fmt.Println(<-messages)
    fmt.Println(<-messages)

    // 当数据输出后，又可以往通道中填充数据
    messages <- "buffered"
    messages <- "channel"

    // 再次输出通道中的数据
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
