package main

import (
	"fmt"
	"os"
	"strings"
)

func exercise20()  {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Error: expected <string><font>")
	}

	input := os.Args[1]

	render := loadfile20("standard.txt")
	
	result := countTheChar(input)
	for i, Number := range result {
		fmt.Printf("%s: %d\n", string(i), Number)
	}

	str := render20(input, render)
	// fmt.Print(render20(input, render))
	fmt.Print(str)

	err := os.WriteFile("output.txt", []byte(str), 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successully written to the output.txt")
}


func countTheChar(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, ch := range s {
		counts[ch]++
	}
	return counts
}

func render20(s string, font [][]string) string {
	var result strings.Builder
	for row := 0; row < 8; row++{
		for _, ch := range s {
			index := int(ch) - 32
			result.WriteString(font[index][row])
		}
		result.WriteString("\n")
	}
	return result.String()

}

func loadfile20(s string) [][]string {
	file, err := os.ReadFile(s)
	if err != nil {
		fmt.Println(err)
	}
	out1 := [][]string{}
	blocks := strings.Split(string(file), "\n\n")

	for _, block := range blocks {
		rows := strings.Split(block, "\n")
		out1 = append(out1, rows)
	}
	return out1
}