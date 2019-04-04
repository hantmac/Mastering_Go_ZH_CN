package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func gen(min, max int, createNumber chan int, end chan bool) {
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		case <-time.After(4 * time.Second):
			fmt.Println("\ntime.After()!")
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)

	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numbers.\n", n)
	go gen(0, 2*n, createNumber, end)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-createNumber)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Exting...")
	end <- true
}
