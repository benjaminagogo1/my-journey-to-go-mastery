package main

import (
	"os"
	"strings"
)

func asciiArt(s string, r string) string {
	file, err := os.ReadFile(r)
	if err != nil {
		return ""
	}

	banner := strings.Split(string(file), "\n")

	// Convert literal "\n" → real newline
	s = strings.ReplaceAll(s, "\\n", "\n")

	segments := strings.Split(s, "\n")

	var out strings.Builder

	for i, segment := range segments {

		// ❗ Skip artificial trailing empty segment
		if segment == "" && i == len(segments)-1 {
			continue
		}

		// Empty segment → ONE blank line
		if segment == "" {
			out.WriteString("\n")
			continue
		}

		// Render ASCII (8 rows)
		for row := 0; row < 8; row++ {
			for _, ch := range segment {
				// safe ASCII range
				if ch < 32 || ch > 126 {
					continue
				}
				index := (int(ch)-32)*9 + row + 1
				out.WriteString(banner[index])
			}
			out.WriteString("\n")
		}
	}

	return out.String()
}














































































































// func asciiArt(s string, r string) string {
// 	inputFile, err := os.ReadFile(r)
// 	if err != nil {
// 		fmt.Println("Error: Failed to read inputfile", err)
// 		return ""
// 	}

// 	inputFileLine := strings.Split(string(inputFile), "\n")
// 	words := strings.Split(s, "\\n")
// 	result := ""

// 	for _, word := range words {
// 		for i := 0; i < 8; i++ {
// 			for _, ch := range word {
// 				result += inputFileLine[i+(int(ch-32)*9)+1]
// 			}
// 			result += "\n"
// 		}
// 	}
// 	return result
// }
