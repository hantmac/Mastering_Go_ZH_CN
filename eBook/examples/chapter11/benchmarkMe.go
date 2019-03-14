package main

import (
	"fmt"
)

func fibo1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo1(n-1) + fibo1(n-2)
	}
}

func fibo2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibo2(n-1) + fibo2(n-2)
}

func fibo3(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

func main() {
	fmt.Println(fibo1(40))
	fmt.Println(fibo2(40))
	fmt.Println(fibo3(40))
}
