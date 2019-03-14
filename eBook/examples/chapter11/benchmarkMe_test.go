package main

import (
	"testing"
)

var result int

func benchmarkfibo1(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo1(n)
	}
	result = r
}

func benchmarkfibo2(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo2(n)
	}
	result = r
}

func benchmarkfibo3(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo3(n)
	}
	result = r
}

func Benchmark30fibo1(b *testing.B) {
	benchmarkfibo1(b, 30)
}

func Benchmark30fibo2(b *testing.B) {
	benchmarkfibo2(b, 30)
}

func Benchmark30fibo3(b *testing.B) {
	benchmarkfibo3(b, 30)
}

func Benchmark50fibo1(b *testing.B) {
	benchmarkfibo1(b, 50)
}

func Benchmark50fibo2(b *testing.B) {
	benchmarkfibo2(b, 50)
}

func Benchmark50fibo3(b *testing.B) {
	benchmarkfibo3(b, 50)
}
