package main

import "fmt"

func main() {
    str := "This is a sample string"
    targetStr := "sample"

    index := findIndex(str, targetStr)
    fmt.Println(index)
}

func findIndex(str, targetStr string) int {
    for i := 0; i <= len(str)-len(targetStr); i++ {
        if str[i:i+len(targetStr)] == targetStr {
            return i
        }
    }
    return -1
}
