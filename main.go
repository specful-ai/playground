package main

import (
    "fmt"
    "os"
    "strconv"
)

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a number as a command line argument.")
        return
    }
    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Invalid number provided.")
        return
    }
    fmt.Println(factorial(n))
}
