package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	MIN := 0
	MAX := 100
	TOTAL := 100
	SEED := time.Now().Unix()

	arguments := os.Args
	switch len(arguments) {
	case 2:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(arguments[1])
		MAX = MIN + 100
	case 3:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(arguments[1])
		MAX, _ = strconv.Atoi(arguments[2])
	case 4:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(arguments[1])
		MAX, _ = strconv.Atoi(arguments[2])
		TOTAL, _ = strconv.Atoi(arguments[3])
	case 5:
		MIN, _ = strconv.Atoi(arguments[1])
		MAX, _ = strconv.Atoi(arguments[2])
		TOTAL, _ = strconv.Atoi(arguments[3])
		SEED, _ = strconv.ParseInt(arguments[4], 10, 64)
	default:
		fmt.Println("Using default values!")
	}

	rand.Seed(SEED)
	for i := 0; i < TOTAL; i++ {
		myrand := random(MIN, MAX)
		fmt.Print(myrand)
		fmt.Print(" ")
	}
	fmt.Println()
}