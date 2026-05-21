package main

import (
	"fmt"
	"strings"
)

func PadArtRows(rows []string, width int) []string {
	if width <= 0 {
		return rows
	}
	res := make([]string, len(rows))
	for i, v := range rows {
		padding := width - len(v)
		if padding > 0 {
			res[i] = strings.Repeat(" ", padding) + v
		} else {
			res[i] = v
		}
	}
	return res
}

func main() {
	words := []string{
		"h",
		"there",
		"come",
		"go",
		"hey",
	}
	for _, ch := range PadArtRows(words, 8) {
		fmt.Printf("| %s |\n", ch)
	}
}
