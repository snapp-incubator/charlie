package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("python", "script.py")
	cmd.Stdin = strings.NewReader("2")

	var out strings.Builder

	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	result := out.String()
	result = strings.Trim(result, "\n")

	fmt.Println(result == "2")
}
