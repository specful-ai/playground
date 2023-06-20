1:package main
2:
3:import "testing"
4:
5:func TestFactorial(t *testing.T) {
6:    result := factorial(4)
7:    if result != 24 {
8:        t.Errorf("Expected factorial(4) to be 24, but got %d", result)
9:    }
10:}
