package main

import "strings"

func render(colorName, substr, userInput string, font [][]string) string {
	if userInput == "" {
		return userInput
	}
	colors := map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"blue":   "\033[34m",
		"yellow": "\033[33m",
	}
	colorNam := colors[colorName]

	userInput = strings.ReplaceAll(userInput, `\n`, "\n")

	var result strings.Builder

	words := strings.Split(userInput, "\n")

	start := strings.Index(userInput, substr)
	end := start + len(substr)

	for _, char := range words {
		for rows := 0; rows < 8; rows++ {
			for i, ch := range char {
				index := int(ch) - 32
				if index < 0 || index > 94 {
					continue
				}
				if i >= start && i < end {
					result.WriteString(colorNam)
					result.WriteString(font[index][rows])
				}
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
