package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	op_num = iota
	op_add
	op_sub
	op_mul
	op_div
)

type frac struct {
	num, denom int
}

type Expr struct {
	op          int
	left, right *Expr
	value       frac
}

func main() {
	rand.Seed(time.Now().UnixNano())
	digits := generateDigits()
	exprs := generateExpressions(digits)
	solution := findSolution(exprs)
	if solution == nil {
		fmt.Println(digits, ": No solution")
	} else {
		fmt.Println(digits, ":", solution)
	}
}

func generateDigits() [4]int {
	var digits [4]int
	for i := 0; i < 4; i++ {
		digits[i] = rand.Intn(9) + 1
	}
	return digits
}

func generateExpressions(digits [4]int) []*Expr {
	var exprs []*Expr
	for i := 0; i < 4; i++ {
		exprs = append(exprs, &Expr{op: op_num, value: frac{num: digits[i], denom: 1}})
	}
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			for _, op := range []int{op_add, op_sub, op_mul, op_div} {
				exprs = append(exprs, &Expr{op: op, left: exprs[i], right: exprs[j]})
			}
		}
	}
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			for k := j + 1; k < 4; k++ {
				for _, op1 := range []int{op_add, op_sub, op_mul, op_div} {
					for _, op2 := range []int{op_add, op_sub, op_mul, op_div} {
						exprs = append(exprs, &Expr{op: op1, left: exprs[i], right: &Expr{op: op2, left: exprs[j], right: exprs[k]}})
						exprs = append(exprs, &Expr{op: op1, left: &Expr{op: op2, left: exprs[i], right: exprs[j]}, right: exprs[k]})
					}
				}
			}
		}
	}
	return exprs
}

func findSolution(exprs []*Expr) *Expr {
	for _, e := range exprs {
		if e.value.num == 24 && e.value.denom == 1 {
			return e
		}
	}
	return nil
}

func (f frac) String() string {
	if f.denom == 1 {
		return fmt.Sprintf("%d", f.num)
	}
	return fmt.Sprintf("%d/%d", f.num, f.denom)
}

func (e *Expr) String() string {
	if e.op == op_num {
		return e.value.String()
	}
	leftStr := e.left.String()
	rightStr := e.right.String()
	switch e.op {
	case op_add:
		return fmt.Sprintf("(%s + %s)", leftStr, rightStr)
	case op_sub:
		return fmt.Sprintf("(%s - %s)", leftStr, rightStr)
	case op_mul:
		return fmt.Sprintf("(%s * %s)", leftStr, rightStr)
	case op_div:
		return fmt.Sprintf("(%s / %s)", leftStr, rightStr)
	default:
		panic("invalid operator")
	}
}
