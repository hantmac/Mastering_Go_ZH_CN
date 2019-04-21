package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	delay := random(0, 15)
	time.Sleep(time.Duration(delay) * time.Second)

	fmt.Fprintf(w, "Servring: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Dealy: %d\n", delay)
	fmt.Printf("Served: %s\n", r.Host)
}

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	PORT := ":8001"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Using default port nubmer: ", PORT)
	} else {
		PORT = ":" + arguments[1]
	}

	http.HandleFunc("/", myHandler)
	err := http.ListenAndServer(PORT, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(10)
	}
}
