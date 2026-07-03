package main

import (
	// "fmt"
	"strings"
)

func render(colorName, substr, userInput string, font [][]string) string {
	if userInput == "" {
		return userInput
	}
	userInput = strings.ReplaceAll(userInput, `\n`, "\n")

	var result strings.Builder

	words := strings.Split(userInput, "\n")

	lengthOfSubstr := 0
	for _, char := range words {
		for rows := 0; rows < 8; rows++{
			for i, ch := range char {
				index := int(ch) - 32
				if strings.HasPrefix(char[i:], substr) {
					lengthOfSubstr = len(substr)
				}
				if lengthOfSubstr > 0 {
					result.WriteString(colorName)
					result.WriteString(font[index][rows])
					result.WriteString("\033[0m")
					lengthOfSubstr --
				}else {
					result.WriteString(font[index][rows])
				}
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
