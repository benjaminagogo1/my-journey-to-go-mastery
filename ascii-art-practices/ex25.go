package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "standard.txt"

	input := os.Args[1]
	reas := os.Args[2]

	re, err := strconv.Atoi(reas)
	if err != nil {
		fmt.Println("It is here!")
	}

	r := strings.Repeat(input, re)

	font := loadfile25(fileName)

	fmt.Println(render25(r + "\n", font))

}

func loadfile25(s string) [][]string {
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

func render25(s string, font [][]string) string {
	var result strings.Builder
	for row := 0; row < 8; row++ {
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
