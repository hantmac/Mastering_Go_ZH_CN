package main

import (
    "aPackage"
    "fmt"
)

func main() {
    fmt.Println("Using aPackage!")
    aPackage.A()
    aPackage.B()
    fmt.Println(aPackage.MyConstant)
}