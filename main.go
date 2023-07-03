package main

import (
	"fmt"
)

func main() {
	precision := 100
	pi := calculatePi(precision)
	fmt.Printf("Pi: %.*f\n", precision, pi)
}

func calculatePi(precision int) float64 {
	pi := 0.0
	sign := 1.0
	for i := 0; i < precision; i++ {
		term := 1.0 / (2.0*float64(i) + 1.0)
		pi += sign * term
		sign *= -1.0
	}
	return 4 * pi
}
