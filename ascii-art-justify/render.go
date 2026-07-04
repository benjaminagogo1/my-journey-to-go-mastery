package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func renderArtToLeft(str string, font map[rune][]string) string {
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
	spaces := (terminalWidth - artwidth) / 800

	for rows := 0; rows < 8; rows++ {
		fmt.Println(rows)
		if spaces > 0 {
			result.WriteString(strings.Repeat(" ", spaces))
		}
		for _, ch := range str {
			fmt.Println(string(ch))
			result.WriteString(strings.Repeat(" ", spaces))
			result.WriteString(font[ch][rows])
		}
		result.WriteString("\n")
		fmt.Println(artwidth)
		fmt.Println(terminalWidth)
	}
	return result.String()
}

func renderArtToRight(str string, font map[rune][]string) string {
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
	spaces := (terminalWidth - artWidth) / 2

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

func justifyAsciiArt(str string, font map[rune][]string) string {
	if str == "" {
		return str
	}

	terminalWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		terminalWidth = 80
	}

	words := strings.Fields(str)

	// If there's only one word, justification isn't possible.
	if len(words) == 1 {
		return renderArtToLeft(words[0], font)
	}

	gaps := len(words) - 1

	// Calculate the width of the words only (no spaces).
	artWidth := 0
	for _, word := range words {
		artWidth += bannerWidth(word, font)
	}

	remaining := terminalWidth - artWidth
	spacesPerGap := remaining / gaps
	extraSpaces := remaining % gaps

	var result strings.Builder

	for row := 0; row < 8; row++ {
		for i, word := range words {
			for _, ch := range word {
				result.WriteString(font[ch][row])
			}

			if i < gaps {
				gapWidth := spacesPerGap
				if i < extraSpaces {
					gapWidth++
				}
				result.WriteString(strings.Repeat(" ", gapWidth))
			}
		}
		result.WriteByte('\n')
	}

	return result.String()
}
