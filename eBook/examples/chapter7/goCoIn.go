package main

import (
	"fmt"
)

type first struct{}

func (a first) F() {
	a.shared()
}

func (a first) shared() {
	fmt.Println("This is shared() from first!")
}

type second struct {
	first
}

func (a second) shared() {
	fmt.Println("This is shared() from second!")
}

func main() {
	first{}.F()
	second{}.shared()
	i := second{}
	j := i.first
	j.F()
}
