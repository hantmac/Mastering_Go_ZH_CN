package main

import (
	"fmt"
)

func f1() int {
	fmt.Println("Entering f1()")
	return -10
	fmt.Println("Exiting f1()")
	return -1
}

func f2() int {
	if true {
		return 10
	}
	fmt.Println("Exiting f2()")
	return 0
}

func main() {
	fmt.Println(f1())
	fmt.Println("Exiting program...")
}
