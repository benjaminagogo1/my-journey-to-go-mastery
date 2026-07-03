package main

import (
	"strings"
)

func render(colorName, substr, userInput string, font map[rune][]string) string {
	if userInput == "" {
		return userInput
	}
	userInput = strings.ReplaceAll(userInput, `\n`, "\n")

	var result strings.Builder

	slice := strings.Split(userInput, "\n")

	lengthOfSubstr := 0
	for _, word := range slice {
		for rows := 0; rows < 8; rows++ {
			for i, ch := range word {

				if strings.HasPrefix(word[i:], substr) {

					lengthOfSubstr = len(substr)

				}
				if lengthOfSubstr > 0 {

					result.WriteString(colorName)
					result.WriteString(font[ch][rows])
					result.WriteString("\033[0m")
					lengthOfSubstr--
				} else {
					result.WriteString(font[ch][rows])
				}
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
