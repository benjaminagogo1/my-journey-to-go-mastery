package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Error: Usage go run . <string> [Font]")
		os.Exit(1)
	}

	r := "standard.txt"
	if len(os.Args) == 3 {
		r = os.Args[2]
		if !strings.HasSuffix(r, ".txt") {
			r += ".txt"
		}
	}
	input := os.Args[1]
	fmt.Print(asciiArt(input, r))
}
