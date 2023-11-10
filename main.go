package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64

	fmt.Print("a= ")
	fmt.Scan(&a)

	fmt.Println("S=", math.Pow(a, 2))

}
