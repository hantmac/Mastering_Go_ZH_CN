package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("You are using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("with Go version", runtime.Version())
}
