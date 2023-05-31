package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type LispFunc func(args []int) int

var env = map[string]LispFunc{
    "+": func(args []int) int {
        sum := 0
        for _, arg := range args {
            sum += arg
        }
        return sum
    },
    "-": func(args []int) int {
        if len(args) == 0 {
            return 0
        }
        diff := args[0]
        for _, arg := range args[1:] {
            diff -= arg
        }
        return diff
    },
    "*": func(args []int) int {
        prod := 1
        for _, arg := range args {
            prod *= arg
        }
        return prod
    },
    "/": func(args []int) int {
        if len(args) == 0 {
            return 0
        }
        quotient := args[0]
        for _, arg := range args[1:] {
            quotient /= arg
        }
        return quotient
    },
}

func eval(expr string) int {
    tokens := strings.Fields(expr)
    if len(tokens) == 0 {
        return 0
    }
    if num, err := strconv.Atoi(tokens[0]); err == nil {
        return num
    }
    fn := env[tokens[0]]
    args := make([]int, 0)
    for _, token := range tokens[1:] {
        args = append(args, eval(token))
    }
    return fn(args)
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a lisp program: ")
    expr, _ := reader.ReadString('\n')
    fmt.Println(eval(expr))
}
