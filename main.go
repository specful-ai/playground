package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Loop through 9 times to produce 9 AOE damages
	for i := 1; i <= 9; i++ {
		// Generate a random damage value between 1 and 100
		damage := rand.Intn(100) + 1
		// Print the AOE damage with a gradient-enhanced label
		fmt.Printf("AOE damage %d: %d (gradient-enhanced)\n", i, damage)
	}
}
