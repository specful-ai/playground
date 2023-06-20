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
        fmt.Println("Please provide a number as command line argument")
        return
    }
    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Invalid input. Please provide a valid number as command line argument")
        return
    }
    fmt.Println(factorial(n))
}

func TestFactorial(t *testing.T) {
    result := factorial(5)
    if result != 120 {
        t.Errorf("factorial(5) = %d; want 120", result)
    }
}
