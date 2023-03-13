package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func main() {

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(add(3, 3))
}
