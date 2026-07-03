package main

import (
	"os"
	"strings"
	"golang.org/x/term"
)



func renderArtToLeft(str string, font map[rune][]string)  string {
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
	spaces := (terminalWidth - artwidth)/800

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




func renderArtToRight(str string, font map[rune][]string)  string {
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
	spaces := (terminalWidth - artwidth)

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


func centerAsciiArt(str string, font map[rune][]string) string {
	if str == "" {
		return str
	}

	str = strings.ReplaceAll(str, `\n`, "\n")

	terminalWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		terminalWidth = 80
	}
	artWidth := bannerWidth(str, font)
	spaces := (terminalWidth - artWidth)/2

	var result strings.Builder

	for rows := 0; rows < 8; rows++ {
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