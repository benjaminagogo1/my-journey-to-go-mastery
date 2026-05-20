package main

import (
	"fmt"
	"os"
	"strings"
)

func contain(rows []int, row int) bool {
	for _, ch := range rows {
		if ch == row {
			return true
		}
	}
	return false
}

func renderEx16(row string, font [][]string)  {
	slectedRows := []int{0, 3, 2, 5, 7}
	for r := 0; r < 8; r++ {
		if contain(slectedRows, r) {
			for _, ch := range row {
				index := int(ch) - 32
				fmt.Print(font[index][r])
			}
			fmt.Println()
		}
	}
}

func loadfile16(s string) [][]string  {
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

func exercise16()  {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("error")
	}
	file := "standard.txt"
	if len(os.Args) == 3 {
		file = os.Args[2]
	}
	input := os.Args[1]
	font := loadfile16(file)
	renderEx16(input, font)
}