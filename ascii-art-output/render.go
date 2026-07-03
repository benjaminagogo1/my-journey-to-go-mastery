package main

import (
	"strings"
)

func render(str string, font [][]string) string {
	if str == "" {
		return str
	}
	str = strings.ReplaceAll(str, `\n`, "\n")

	words := strings.Split(str, "\n")
	var result strings.Builder

	for _, char := range words {
		for rows := 0; rows < 8; rows ++ {
			for _, ch := range char {
				index := int(ch) - 32 
				if index < 0 || index > 94 {
					continue
				}
				result.WriteString(font[index][rows])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}