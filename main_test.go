package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	expected := "hello world"
	result := "hello world"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
