package main

import (
    "fmt"
    "os"
    "strconv"
)

func permute(nums []int, l int, r int) {
    if l == r {
        fmt.Println(nums)
    } else {
        for i := l; i <= r; i++ {
            nums[l], nums[i] = nums[i], nums[l]
            permute(nums, l+1, r)
            nums[l], nums[i] = nums[i], nums[l]
        }
    }
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <n>")
        os.Exit(1)
    }
    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Invalid argument for n")
        os.Exit(1)
    }
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        nums[i] = i + 1
    }
    permute(nums, 0, n-1)
}
