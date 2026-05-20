package main

import (
	"fmt"
	"os"
	"strings"
)

func renderEx15(input string, font [][]string) {
	for row := 0; row < 8; row++ {
		fmt.Printf("%2d: ", row)
		for _, ch := range input {
			index := int(ch) - 32
			if index < 0 || index > 94 {
				continue
			}
			fmt.Print(font[index][row])
			
		}
		fmt.Println()
	}
}

func loadfile15(s string) [][]string {
	file, err := os.ReadFile(s)
	if err != nil {
		fmt.Println(err)
	}

	out := [][]string{}

	content := string(file)

	blocks := strings.Split(content, "\n\n")

	for _, block := range blocks {
		rows := strings.Split(block, "\n")
		out = append(out, rows)
	}
	return out
}

func exercise15() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("invalid Input")
	}

	fileName := "standard.txt"
	if len(os.Args) == 3 {
		fileName = os.Args[2]
	}

	input := os.Args[1]
	font := loadfile15(fileName)
	renderEx15(input, font)
}
