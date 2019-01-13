package main

import (
	"fmt"
	"math"
	"myInterface"
)

type square struct {
	X float64
}

type circle struct {
	R float64
}

func (s square) Area() float64 {
	return s.X * s.X
}

func (s square) Perimeter() float64 {
	return 4 * s.X
}

func (s circle) Area() float64 {
	return s.R * s.R * math.Pi
}

func (s circle) Perimeter() float64 {
	return 2 * s.R * math.Pi
}

func Calculate(x myInterface.Shape) {
	_, ok := x.(circle)
	if ok {
		fmt.Println("Is a circle!")
	}
	v, ok := x.(square)
	if ok {
		fmt.Println("Is a square:", v)
	}
	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())
}

func main() {
	x := square{X: 10}
	fmt.Println("Perimeter:", x.Perimeter())
	Calculate(x)
	y := circle{R: 5}
	Calculate(y)
}
