package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}

	fileName := "standard.txt"
	if len(os.Args) == 4 {
		fileName = os.Args[3]
		if !strings.HasPrefix(fileName, ".txt") {
			fileName += ".txt"
		}
	}

	userInput := os.Args[2]
	outputFile := os.Args[1]

	correct, err := splitFlag(outputFile)
	
	if err != nil {
		fmt.Println("Error: usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}
	result := Loadbanner(fileName)

	art := render(userInput, result)


	err = os.WriteFile(correct, []byte(art), 0664)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully written to the file")
}


