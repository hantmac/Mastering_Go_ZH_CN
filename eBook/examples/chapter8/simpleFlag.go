package main

import (
    "flag"
    "fmt"
)

func main(){
    minusk := flag.Bool("k", true, "k")
    minusO := flag.Int("O", 1, "O")
	flag.Parse()
	
	valueK := *minusk
	valueO := *minusO
	valueO++

	fmt.Println("-k:", valueK)
	fmt.Println("-O:", valueO)
}