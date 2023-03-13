package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2    // initialize z to x/2
	e := 0.000001 // define the desired precision

	for math.Abs(z*z-x) > e {
		z -= (z*z - x) / (2 * z) //newton's algorithm
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
