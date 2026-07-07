package main

import (
	"strings"
)

func render(str string, font map[rune][]string) string {
	if str == "" {
		return str
	}

	replace := strings.NewReplacer(
		`\n`, "\n",
		`\t`, "\t",
		"\r\n", "\n",
	)
	str = replace.Replace(str)

	words := strings.Split(str, "\n")
	var result strings.Builder
	for _, char := range words {
		for rows := 0; rows < 8; rows++ {
			for _, ch := range char {
				if ch < 32 || ch > 126 {
					continue
				}
				result.WriteString(font[ch][rows])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}