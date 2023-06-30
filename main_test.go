package main

import "testing"

func TestFactorial(t *testing.T) {
    result := factorial(5)
    if result != 120 {
        t.Errorf("Expected factorial(5) to be 120, but got %d", result)
    }
}
