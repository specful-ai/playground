package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("git", "log", "-1", "--merges", "--pretty=format:%P")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	parents := strings.Split(strings.TrimSpace(string(output)), " ")
	for _, parent := range parents {
		fmt.Println(parent)
	}
}
