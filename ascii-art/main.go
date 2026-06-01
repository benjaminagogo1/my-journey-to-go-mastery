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






package main

import (
	"fmt"
)

type studentResult struct {
		Name string
		Age int
		Scores map[string]int{
			"key1" : value1
			"key2" : value2
			"key3" : value3
		}
	}
func studentReport(student string) []string {
	
	studentScores := studentResult{
		Name : "Benjamin",
		Age : 22,
		Scores : map[string]int{
			"Maths" : 85
			"English" : 70
			"Chemistry" : 90
		},

	}
}