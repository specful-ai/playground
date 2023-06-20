package main

import "fmt"

func factorial(k int) int {
    if k == 0 {
        return 1
    }
    return k * factorial(k-1)
}

func main() {
    fmt.Println(factorial(5))
}
