package main

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println("l", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)
	fmt.Println("Read:", <-c)
	time.Sleep(1 * time.Second)
	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}
