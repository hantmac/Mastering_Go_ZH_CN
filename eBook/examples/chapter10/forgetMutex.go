package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex

func function() {
	m.Lock()
	fmt.Println("Locked!")
}

func main() {
	var w sync.WaitGroup

	go func() {
		defer w.Done()
		function()
	}()
	w.Add(1)

	go func() {
		defer w.Done()
		function()
	}()
	w.Add(1)

	w.Wait()
}
