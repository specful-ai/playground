package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    solutions := solveNQueens(n)
    fmt.Println(len(solutions))
}

func solveNQueens(n int) [][]string {
    var res [][]string
    cols := make(map[int]bool)
    diag1 := make(map[int]bool)
    diag2 := make(map[int]bool)
    board := make([]string, n)
    for i := 0; i < n; i++ {
        board[i] = ""
        for j := 0; j < n; j++ {
            board[i] += "."
        }
    }
    queue := [][][]string{{{board}}}
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        last := curr[len(curr)-1]
        if len(last) == n {
            res = append(res, last)
            continue
        }
        row := len(last)
        for col := 0; col < n; col++ {
            if cols[col] || diag1[row+col] || diag2[row-col] {
                continue
            }
            cols[col], diag1[row+col], diag2[row-col] = true, true, true
            newBoard := make([]string, n)
            copy(newBoard, last)
            newBoard[row] = newBoard[row][:col] + "Q" + newBoard[row][col+1:]
            queue = append(queue, append(curr, newBoard))
            cols[col], diag1[row+col], diag2[row-col] = false, false, false
        }
    }
    return res
}
