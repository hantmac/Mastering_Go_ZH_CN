package documentMe

import (
	"fmt"
)

func ExampleS1() {
	fmt.Println(S1("123456789"))
	fmt.Println(S1(""))
	// Output:
	// 9
	// 0
}

func ExampleF1() {
	fmt.Println(F1(10))
	fmt.Println(F1(2))
	// Output:
	// 1
	// 55
}
