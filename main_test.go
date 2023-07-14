package main

import "testing"

func TestFindIndex(t *testing.T) {
    str := "This is a sample string"
    targetStr := "sample"

    index := findIndex(str, targetStr)
    expected := 10

    if index != expected {
        t.Errorf("Expected index %d, but got %d", expected, index)
    }
}

func TestFindIndex_NotFound(t *testing.T) {
    str := "This is a sample string"
    targetStr := "notfound"

    index := findIndex(str, targetStr)
    expected := -1

    if index != expected {
        t.Errorf("Expected index %d, but got %d", expected, index)
    }
}
