package main

import (
	"fmt"
	"unicode"
)

func main() {
	const sL= "\x99\x00ab\x50\x00\x23\x50\x29\x9c"

	for i := 0; i < len(sL); i++ {
		if unicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c\n", sL[i])
		} else {
			fmt.Println("Not printable!")
		}
	}
}