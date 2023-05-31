package main

import (
	"fmt"
	"regexp"
)

func isValidRegex(str string) bool {
	_, err := regexp.Compile(str)
	if err != nil {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValidRegex("^[a-zA-Z0-9]+([._]?[a-zA-Z0-9]+)*@[a-zA-Z0-9]+([.-]?[a-zA-Z0-9]+)*\\.([a-zA-Z]{2,6})$"))
	fmt.Println(isValidRegex("[a-z]+"))
	fmt.Println(isValidRegex("(hello"))
}
