package main

import (
	"github.com/tatsushid/go-fastping"
	"net"
	"os"
	"fmt"
	"time"
)

func main() {
    addr := "192.168.1.8"

	ra, err := net.ResolveIPAddr("ip4:icmp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

    p := fastping.NewPinger()
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
}
