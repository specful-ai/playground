package main

import (
	"math"
	"testing"
)

func TestCalculatePi(t *testing.T) {
	precision := 1000
	expectedPi := math.Pi
	calculatedPi := calculatePi(precision)

	if math.Abs(calculatedPi-expectedPi) > 0.0001 {
		t.Errorf("Expected Pi: %f, Calculated Pi: %f", expectedPi, calculatedPi)
	}
}
