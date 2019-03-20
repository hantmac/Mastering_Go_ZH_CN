package main

import (
	"fmt"
	"os"
	"testing"
)

var ERR error

func benchmarkCreate(b *testing.B, buffer, filesize int) {
	var err error
	for i := 0; i < b.N; i++ {
		err = Create("/tmp/random", buffer, filesize)
	}
	ERR = err
	err = os.Remove("/tmp/random")
	if err != nil {
		fmt.Println(err)
	}
}

func Benchmark1Create(b *testing.B) {
	benchmarkCreate(b, 1, 1000000)
}

func Benchmark2Create(b *testing.B) {
	benchmarkCreate(b, 2, 1000000)
}

func Benchmark4Create(b *testing.B) {
	benchmarkCreate(b, 4, 1000000)
}

func Benchmark10Create(b *testing.B) {
	benchmarkCreate(b, 10, 1000000)
}

func Benchmark1000Create(b *testing.B) {
	benchmarkCreate(b, 1000, 1000000)
}
