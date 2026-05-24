package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Invalid input")
		return
	}

	fileName := "standard.txt"


	input := os.Args[1]
	colorName := os.Args[3]
	str := os.Args[2]

	input = strings.ReplaceAll(input, `\n`, "\n")
	font := loadfile23(fileName)

	lines := strings.Split(input, "\n\n")
	for _, line := range lines {
		fmt.Println(render23(line, str, colorName, font))
	}
}

func bannerWidth23(s string, font [][]string) int {
	count := 0
	for _, ch := range s {
		index := int(ch) - 32
		if index < 0 || index > 94 {
			continue
		}
		count += len(font[index][1])
	}
	return count
}

func loadfile23(s string) [][]string {
	file, err := os.ReadFile(s)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	output := [][]string{}

	blocks := strings.Split(string(file), "\n\n")

	for _, block := range blocks {
		if block == "" {
			continue
		}
		rows := strings.Split(block, "\n")
		output = append(output, rows)
	}
	return output
}



func render23(input string, str string, color string, font [][]string) string {
	colors := map[string]string{
		"red" : "\033[31m",
		"green" : "\033[32m",
		"blue" :  "\033[34m",
		"yellow" : "\033[33m",
	}
	colorName := colors[color]
	terminalWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		terminalWidth = 80
	}
	artWidth := bannerWidth23(input, font)
	padding := (terminalWidth - artWidth)/2
	
	var result strings.Builder

	start := strings.Index(input, str)
	end := start + len(str)
	
	for row := 0; row < 8; row++{
		if padding > 0 {
			result.WriteString(strings.Repeat(" ", padding))
		}
		for i, ch := range input {
			index := int(ch) - 32 
			if index < 0 || index > 94 {
				continue
			}
			if i >= start && i < end {
				
				result.WriteString(colorName)
				result.WriteString(font[index][row])
				result.WriteString("\033[0m")
				// result.WriteString("\033[47m")
			} else {
				result.WriteString(font[index][row])
			}
			
		}
		result.WriteString("\n")
	}
	return result.String()
	
}

