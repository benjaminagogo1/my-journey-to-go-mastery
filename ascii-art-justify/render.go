package main

import (
	"os"
	"strings"
	"golang.org/x/term"
)



func render(str string, font map[rune][]string)  string {
	if str == "" {
		return str
	}

	str = strings.ReplaceAll(str, `\n`, "\n")

	terminalWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		terminalWidth = 80
	}
	var result strings.Builder
	artwidth := bannerWidth(str, font)
	spaces := (terminalWidth - artwidth)/2

	for rows := 0; rows < 8; rows++{
		if spaces > 0 {
			result.WriteString(strings.Repeat(" ", spaces))
		}
		for _, ch := range str {

			result.WriteString(font[ch][rows])
		}
		result.WriteString("\n")
	}
	return result.String()
}