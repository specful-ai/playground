package main

import (
	"fmt"
	"math/rand"
	"time"
)

type node struct {
	val  float64
	expr string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	digits := generateDigits()
	fmt.Println("Digits:", digits)
	solutions := solve24(digits)
	if len(solutions) == 0 {
		fmt.Println("No solution")
	} else {
		for _, sol := range solutions {
			fmt.Println(sol.expr, "=", sol.val)
		}
	}
}

func generateDigits() [4]int {
	var digits [4]int
	for i := 0; i < 4; i++ {
		digits[i] = rand.Intn(9) + 1
	}
	return digits
}

func solve24(digits [4]int) []node {
	var solutions []node
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			for k := 0; k < 4; k++ {
				if i == k || j == k {
					continue
				}
				for l := 0; l < 4; l++ {
					if i == l || j == l || k == l {
						continue
					}
					nodes := solve3(node{float64(digits[i]), fmt.Sprintf("%d", digits[i])},
						node{float64(digits[j]), fmt.Sprintf("%d", digits[j])},
						node{float64(digits[k]), fmt.Sprintf("%d", digits[k])},
						node{float64(digits[l]), fmt.Sprintf("%d", digits[l])})
					for _, n := range nodes {
						if n.val == 24 {
							solutions = append(solutions, n)
						}
					}
				}
			}
		}
	}
	return solutions
}

func solve3(a, b, c, d node) []node {
	var nodes []node
	// ((a op b) op c) op d
	for _, op1 := range []func(float64, float64) float64{add, sub, mul, div} {
		for _, op2 := range []func(float64, float64) float64{add, sub, mul, div} {
			for _, op3 := range []func(float64, float64) float64{add, sub, mul, div} {
				val := op3(op2(op1(a.val, b.val), c.val), d.val)
				expr := fmt.Sprintf("((%s %s %s) %s %s) %s %s", a.expr, op1.__name__(), b.expr, op2.__name__(), c.expr, op3.__name__(), d.expr)
				nodes = append(nodes, node{val, expr})
			}
		}
	}
	// (a op (b op c)) op d
	for _, op1 := range []func(float64, float64) float64{add, sub, mul, div} {
		for _, op2 := range []func(float64, float64) float64{add, sub, mul, div} {
			val := op2(op1(a.val, b.val), op1(c.val, d.val))
			expr := fmt.Sprintf("(%s %s %s) %s (%s %s %s)", a.expr, op1.__name__(), b.expr, op2.__name__(), c.expr, op1.__name__(), d.expr)
			nodes = append(nodes, node{val, expr})
		}
	}
	return nodes
}

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	return a / b
}

func (f func(float64, float64) float64) __name__() string {
	switch f {
	case add:
		return "+"
	case sub:
		return "-"
	case mul:
		return "*"
	case div:
		return "/"
	default:
		return ""
	}
}
