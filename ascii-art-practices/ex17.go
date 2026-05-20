package main

import (
	"os"
	"fmt"
	"strings"
)

func contain2(row2 []int, row1 int) bool {
	for _, ch := range row2 {
		if ch == row1 {
			return true 
		}
	}
	return false	
}

func render17(input string, fontfile [][]string)  {
	slectedRows2 := []int{0, 3, 2, 5, 7}
	for row := 0; row < 8; row++ {
		if contain2(slectedRows2, row) {
			for _, ch := range input {
				index := int(ch) -32 
				fmt.Print(fontfile[index][row])
			}
			fmt.Println()
		}
	}
}

func loadfile17(s string) [][]string {
	file, err := os.ReadFile(s)
	if err != nil {
		fmt.Println(err)
	}

	result := [][]string{}
	content := string(file)

	blocks := strings.Split(content, "\n\n")
	for _, block := range blocks {
		rows := strings.Split(block, "\n")
		result = append(result, rows)
	}

	return result
	
}

func exercise17()  {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("error")
	}
	fileName1 := "standard.txt"
	if len(os.Args) == 3 {
		fileName1 = os.Args[2]
	}
	input := os.Args[1]
	font1 := loadfile17(fileName1)
	render17(strings.ToUpper(input), font1)
}