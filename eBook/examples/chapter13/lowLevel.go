package main

import (
	"fmt"
	"net"
)

func main() {
	netaddr, err := net.ResolveIPAddr("ip4", "127.0.0.1")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.ListenIP("ip4:icmp", netaddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFrom(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("% X\n", buffer[0:n])
}
