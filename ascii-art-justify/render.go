package main

import (
	"fmt"
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





func justifyAsciiArt(str string, font map[rune][]string) string {
	if str == "" {
		return str
	}

	str = strings.ReplaceAll(str, `\n`, "\n")

	terminalWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		terminalWidth = 80
	}
	words := strings.Fields(str)
	gap := len(words)-1
	artWidth := bannerWidth(str, font)
	spaces := terminalWidth - artWidth
	spacesPerGap := spaces/gap

	var result strings.Builder
	for rows := 0; rows < 8; rows ++{
		for i, word := range words {
			for _, ch := range word {
				result.WriteString(font[ch][rows])
			}
			if i != len(words)-1 {
				result.WriteString(strings.Repeat(" ", spacesPerGap))
			}
			
		}
		result.WriteString("\n")
	}
	return result.String()
}



// between them.

// Compute

// totalWidth =
// width(We) +
// width(will) +
// width(explain!)

// Notice that do not count the spaces yet.

// Then

// remainingSpace = terminalWidth - totalWidth

// Now divide the remaining space among the gaps.

// spacesPerGap = remainingSpace / numberOfGaps
// extraSpaces = remainingSpace % numberOfGaps

// If

// remainingSpace = 23
// gaps = 2

// then

// spacesPerGap = 11
// extraSpaces = 1

// So you print

// gap1 = 12 spaces
// gap2 = 11 spaces

// The first gaps receive the extra spaces.

// Example

// Suppose

// terminal = 30

// word widths

// We        = 4
// will      = 8
// explain   = 10

// Total

// 4 + 8 + 10 = 22

// Remaining

// 30 - 22 = 8

// There are

// 2 gaps

// Each gets

// 8 / 2 = 4

// So you print

// We____will____explain

// (where _ means spaces)

// Notice something important:

// If there is only one word, there are zero gaps.

// How can you justify a single word?



// Most implementations simply print it as left aligned.

// I wouldn't start coding this immediately. The first thing I'd implement is a function that, given:

// terminal width,
// widths of the ASCII-art words,

// returns how many spaces should go after each word. Once that logic works, plugging it into your 8-row ASCII-art rendering becomes much easier.