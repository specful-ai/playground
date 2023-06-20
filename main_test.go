package main

import (
	"reflect"
	"testing"
)

func TestGetPermutations(t *testing.T) {
	tests := []struct {
		n        int
		expected [][]int
	}{
		{1, [][]int{{1}}},
		{2, [][]int{{1, 2}, {2, 1}}},
		{3, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}},
	}

	for _, test := range tests {
		result := getPermutations(test.n)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("getPermutations(%d) = %v, expected %v", test.n, result, test.expected)
		}
	}
}

func TestPermute(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1}, [][]int{{1}}},
		{[]int{1, 2}, [][]int{{1, 2}, {2, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}},
	}

	for _, test := range tests {
		result := permute(test.nums)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("permute(%v) = %v, expected %v", test.nums, result, test.expected)
		}
	}
}

func TestSwap(t *testing.T) {
	nums := []int{1, 2, 3}
	swap(nums, 0, 2)
	expected := []int{3, 2, 1}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("swap(%v, 0, 2) = %v, expected %v", nums, nums, expected)
	}
}

func TestToString(t *testing.T) {
	nums := []int{1, 2, 3}
	result := toString(nums)
	expected := "123"
	if result != expected {
		t.Errorf("toString(%v) = %s, expected %s", nums, result, expected)
	}
}
