package main

import (
	"fmt"
)

func myAtoi(s string) int {
	sign := 1
	result := 0
	i := 0

	// Remove leading whitespaces
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// Check for sign
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Convert string to integer
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		result = result*10 + int(s[i]-'0')
		if result > 1<<31-1 {
			if sign == 1 {
				return 1<<31 - 1
			} else {
				return -1 << 31
			}
		}
		i++
	}

	return result * sign
}

func main() {
	s := "42"
	fmt.Println(myAtoi(s))

	s = "   -42"
	fmt.Println(myAtoi(s))

	s = "4193 with words"
	fmt.Println(myAtoi(s))

	s = "words and 987"
	fmt.Println(myAtoi(s))

	s = "-91283472332"
	fmt.Println(myAtoi(s))
}
