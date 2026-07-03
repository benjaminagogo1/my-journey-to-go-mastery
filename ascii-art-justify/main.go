package main

import (
	"fmt"
	"os"
)

func main()  {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --align=right something standard")
		return
	}

	fileName := "standard.txt"

	alignType := os.Args[1]
}