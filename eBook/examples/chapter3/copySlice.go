package main

import "fmt"

func main() {
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6)
	fmt.Printf("a4:", a4)

	copy(a6, a4)
	fmt.Println("a6:", a6)
	fmt.Printf("a4:", a4)
	fmt.Println()

	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}
	fmt.Println("b6:", b6)
	fmt.Printf("b4:", b4)

	copy(b4, b6)
	fmt.Println("b6:", b6)
	fmt.Printf("b4:", b4)
	fmt.Println()

	fmt.Println()
	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, -1, 1, -1, -5, 5}

	copy(s6, array4[0:])
	fmt.Println("array4:", array4[0:])
	fmt.Printf("s6:", s6)
	fmt.Println()

	array5 := [5]int{5, -5, 5, -5, 5}
	s7 := []int{7, 7, -7, 7, -7, 7}
	copy(array5[0:], s7)
	fmt.Println("array5:", array5)
	fmt.Printf("s7:", s7)
	fmt.Println()

}
