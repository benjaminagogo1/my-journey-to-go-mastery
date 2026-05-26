package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func Exercise25() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Invalid input")
	}
	fileName1 := "standard.txt"
	input := os.Args[1]
	// colorName := os.Args[3]
	// str := os.Args[2]

	font := loadfile24(fileName1)
	lines := strings.Split(input, "\n\n")
	for _, ch := range lines {
		fmt.Println(normalRender(ch, font))
		fmt.Println(flipRender(ch, font))
	}
	run := []rune(input)
	res1 := ""
	for i := len(run) - 1; i >= 0; i-- {
		res1 += string(run[i])
		f := normalRender(string(run[i]), font)
		fmt.Print(f)
	}
}

func loadfile24(s string) [][]string {
	file, err := os.ReadFile(s)
	if err != nil {
		fmt.Println(err)
	}

	output := [][]string{}

	blocks := strings.Split(string(file), "\n\n")
	for _, block := range blocks {
		rows := strings.Split(block, "\n")
		output = append(output, rows)
	}
	return output
}

func bannerWidth24(s string, font [][]string) int {
	count := 0

	for _, ch := range s {
		index := int(ch) - 32
		if index < 0 || index > 94 {
			continue
		}
		count += len(font[index][0])
	}
	return count
}

// func render24(s string, str string, color string, font [][]string) string {
// 	colors := map[string]string{
// 		"red":    "\033[31m",
// 		"green":  "\033[32m",
// 		"blue":   "\033[34m",
// 		"yellow": "\033[33m",
// 	}
// 	colorName := colors[color]
// 	terminalWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
// 	if err != nil {
// 		terminalWidth = 80
// 	}
// 	artWidth := bannerWidth24(s, font)
// 	padding := (terminalWidth - artWidth) / 2

// 	var normal strings.Builder
// 	var flipped strings.Builder

// 	start := strings.Index(s, str)
// 	end := start + len(str)

// 	for row := 0; row < 8; row++ {
// 		if padding > 0 {
// 			normal.WriteString(strings.Repeat(" ", padding))
// 		}
// 		for i, ch := range s {
// 			index := int(ch) - 32
// 			if index < 0 || index > 94 {
// 				continue

// 			}
// 			if i >= start && i < end {
// 				normal.WriteString(colorName)
// 				normal.WriteString(font[index][row])
// 				normal.WriteString("\033[0m")
// 			} else {
// 				normal.WriteString(font[index][row])
// 			}
// 		}
// 		normal.WriteString("\n")
// 	}
// 	for row := 7; row >= 0; row-- {
// 		if padding > 0 {
// 			flipped.WriteString(strings.Repeat(" ", padding))
// 		}
// 		for i, ch := range s {
// 			index := int(ch) - 32
// 			if index < 0 || index > 94 {
// 				continue
// 			}
// 			if i >= start && i < end {
// 				flipped.WriteString(colorName)
// 				flipped.WriteString(font[index][row])
// 				flipped.WriteString("\033[0m")
// 			} else {
// 				flipped.WriteString(font[index][row])
// 			}
// 		}
// 		flipped.WriteString("\n")
// 	}
// 	return flipped.String() + "\n" + normal.String()

// }

func normalRender(s string, font [][]string) string {
	terminal, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		terminal = 80
	}

	artWidth := bannerWidth24(s, font)
	padding := (terminal - artWidth) / 2

	var result strings.Builder

	for row := 0; row < 8; row++ {
		if padding > 0 {
			result.WriteString(strings.Repeat(" ", padding))
		}
		for _, ch := range s {
			index := int(ch) - 32
			if index < 0 || index > 94 {
				continue
			}
			result.WriteString(font[index][row])
		}
		result.WriteString("\n")
	}
	return result.String()

}

func flipRender(s string, font [][]string) string {
	terminal, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		terminal = 80
	}

	artWidth := bannerWidth24(s, font)
	padding := (terminal - artWidth) / 2

	var result1 strings.Builder

	for row := 7; row >= 0; row-- {
		if padding > 0 {
			result1.WriteString(strings.Repeat(" ", padding))
		}
		for _, char := range s {
			index := int(char) - 32
			if index < 0 || index > 94 {
				continue
			}
			result1.WriteString(font[index][row])
		}
		result1.WriteString("\n")
	}
	return result1.String()
}
