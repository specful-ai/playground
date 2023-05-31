package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nums [4]int
	fmt.Println("Enter 4 integers from 1 to 9:")
	for i := 0; i < 4; i++ {
		fmt.Scan(&nums[i])
	}

	if !isValid(nums) {
		fmt.Println("Invalid input")
		return
	}

	if !hasSolution(nums) {
		fmt.Println("No solution")
		return
	}

	fmt.Println("Solution found")
}

func isValid(nums [4]int) bool {
	for _, num := range nums {
		if num < 1 || num > 9 {
			return false
		}
	}
	return true
}

func hasSolution(nums [4]int) bool {
	ops := []string{"+", "-", "*", "/"}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == i {
				continue
			}
			for k := 0; k < 4; k++ {
				if k == i || k == j {
					continue
				}
				for l := 0; l < 4; l++ {
					if l == i || l == j || l == k {
						continue
					}
					for _, op1 := range ops {
						for _, op2 := range ops {
							for _, op3 := range ops {
								if evaluate(nums[i], nums[j], nums[k], nums[l], op1, op2, op3) == 24 {
									return true
								}
								if evaluate(nums[i], nums[j], nums[l], nums[k], op1, op2, op3) == 24 {
									return true
								}
								if evaluate(nums[i], nums[k], nums[j], nums[l], op1, op2, op3) == 24 {
									return true
								}
								if evaluate(nums[i], nums[k], nums[l], nums[j], op1, op2, op3) == 24 {
									return true
								}
								if evaluate(nums[i], nums[l], nums[j], nums[k], op1, op2, op3) == 24 {
									return true
								}
								if evaluate(nums[i], nums[l], nums[k], nums[j], op1, op2, op3) == 24 {
									return true
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}

func evaluate(a, b, c, d int, op1, op2, op3 string) float64 {
	expr := "(" + strconv.Itoa(a) + op1 + strconv.Itoa(b) + ")" + op2 + "(" + strconv.Itoa(c) + op3 + strconv.Itoa(d) + ")"
	val, _ := strconv.ParseFloat(expr, 64)
	return val
}
