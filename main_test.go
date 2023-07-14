package main

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
	}

	for _, test := range tests {
		result := myAtoi(test.s)
		if result != test.expected {
			t.Errorf("myAtoi(%s) = %d; expected %d", test.s, result, test.expected)
		}
	}
}
