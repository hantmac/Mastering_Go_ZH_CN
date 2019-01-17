package main

import "fmt"

func main(){
	type XYZ struct {
		X int
		Y int
		Z int
	}

	var s1 XYZ
	fmt.Println(s1.Y, s1.Z)

	p1 := XYZ{23,12, -2}
	p2 := XYZ{Z:12,Y:13}

	fmt.Println(p1)
	fmt.Println(p2)

	pSlice := [4]XYZ{}
	pSlice[2] = p1
	pSlice[0] = p2

	fmt.Println(pSlice)

	p2 = XYZ{1,2,3}
	fmt.Println(pSlice)
}