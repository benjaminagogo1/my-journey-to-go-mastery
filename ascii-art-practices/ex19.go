package main

import (
	"fmt"
	"os"
	"strings"
)

func exercise19()  {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Error: Usage: <string><font>")
		return
	}

	input := os.Args[1]

	result := loadfile19("standard.txt")

	str := render19(input, result)

	err := os.WriteFile("output.txt", []byte(str), 0644)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Output written to output.txt")
}


func loadfile19(s string) [][]string {
	file, err := os.ReadFile(s)
	if err != nil {
		fmt.Println(err)
	}
	out := [][]string{}
	blocks := strings.Split(string(file), "\n\n")
	for _, block := range blocks {
		rows := strings.Split(block, "\n")
		out = append(out, rows)
	}
	return out

}


func render19(s string, font [][]string) string {
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

