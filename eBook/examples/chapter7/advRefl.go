package main

import (
	"fmt"
	"os"
	"reflect"
)

type t1 int
type t2 int

type a struct {
	X    int
	Y    float64
	Text string
}

func (a1 a) compareStruct(a2 a) bool {
	r1 := reflect.ValueOf(&a1).Elem()
	r2 := reflect.ValueOf(&a2).Elem()
	for i := 0; i < r1.NumField(); i++ {
		if r1.Field(i).Interface() != r2.Field(i).Interface() {
			return false
		}
	}
	return true
}

func printMethods(i interface{}) {
	r := reflect.ValueOf(i)
	t := r.Type()
	fmt.Printf("Type to examine: %s\n", t)
	for j := 0; j < r.NumMethod(); j++ {
		m := r.Method(j).Type()
		fmt.Println(t.Method(j).Name, "-->", m)
	}
}

func main() {
	x1 := t1(100)
	x2 := t2(100)
	fmt.Printf("The type of x1 is %s\n", reflect.TypeOf(x1))
	fmt.Printf("The type of x2 is %s\n", reflect.TypeOf(x2))
	var p struct{}
	r := reflect.New(reflect.ValueOf(&p).Type()).Elem()
	fmt.Printf("The type of r is %s\n", reflect.TypeOf(r))
	a1 := a{1, 2.1, "A1"}
	a2 := a{1, -2, "A2"}
	if a1.compareStruct(a1) {
		fmt.Println("Equal!")
	}
	if !a1.compareStruct(a2) {
		fmt.Println("Not Equal!")
	}
	var f *os.File
	printMethods(f)
}
