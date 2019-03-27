package ex

import (
	"fmt"
)

func ExampleF1() {
	fmt.Println(F1(10))
	fmt.Println(F1(2))
	// Output:
	// 55
	// 1
}

func ExampleS1() {
	fmt.Println(S1("123456789"))
	fmt.Println(S1(""))
	// Output:
	// 8
	// 0
}
