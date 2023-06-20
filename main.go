package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 3
	permutations := getPermutations(n)
	for _, p := range permutations {
		fmt.Println(p)
	}
}

func getPermutations(n int) [][]int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return permute(nums)
}

func permute(nums []int) [][]int {
	var result [][]int
	backtrack(nums, 0, &result)
	return result
}

func backtrack(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		permutation := make([]int, len(nums))
		copy(permutation, nums)
		*result = append(*result, permutation)
		return
	}

	for i := start; i < len(nums); i++ {
		swap(nums, start, i)
		backtrack(nums, start+1, result)
		swap(nums, start, i)
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func toString(nums []int) string {
	var str string
	for _, num := range nums {
		str += strconv.Itoa(num)
	}
	return str
}
