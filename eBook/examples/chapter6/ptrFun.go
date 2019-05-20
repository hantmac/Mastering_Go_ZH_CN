package main

import (
    "fmt"
)

func getPtr(v *float64) float64 {
    return *v * *v
}

func main() {
    x := 12.2
    fmt.Println(getPrt(&x))
    x = 12
    fmt.Println(getPtr(&x))
}