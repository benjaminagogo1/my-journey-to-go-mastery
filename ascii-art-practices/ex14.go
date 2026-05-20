package main

import (
	"fmt"
	"os"
	"strings"
)

func bannerWidth(word string, font [][]string) int {
	result := 0
	for _, ch := range word {
		index := int(ch) - 32
		result += len(font[index][0])
		fmt.Println("Width of hello : ", len(font[index][0]))
	}
	return result
}

func loadfile(s string) [][]string {

	file, err := os.ReadFile(s)
	if err != nil {
		fmt.Println(err)
	}

	output := [][]string{}

	content := string(file)

	blocks := strings.Split(content, "\n\n")

	for _, block := range blocks {
		rows := strings.Split(block, "\n")

		output = append(output, rows)
	}
	return output
}

func exercise14() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Invalid input")
	}
	fileName := "standard.txt"
	if len(os.Args) == 3 {
		fileName = os.Args[2]
	}

	input := os.Args[1]
	if input == "" {
		return
	}
	font := loadfile(fileName)
	fmt.Println(bannerWidth(input, font))

}
