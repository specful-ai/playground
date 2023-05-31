package main

import "fmt"

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
    n := 4
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        nums[i] = i + 1
    }
    permute(nums, 0, n-1)
}
