package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Give me a natural number!")
		os.Exit(1)
	}
	numberGR, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup
	var i int

	k := make(map[int]int)
	k[1] = 12
	for i = 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			k[i] = i
		}()
	}

	k[2] = 10
	waitGroup.Wait()
	fmt.Println("k = %v\n", k)
}
