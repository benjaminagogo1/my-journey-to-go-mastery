package main

import (
	"os"
	"fmt"
	"strings"
)

func contain3(row2 []int, row1 int) bool {
	for _, ch := range row2 {
		if ch == row1 {
			return true 
		}
	}
	return false	
}

func render18(input string, fontfile [][]string)  {
	slectedRows2 := []int{0, 3, 2, 5, 7}
	for row := 0; row < 8; row++ {
		if contain3(slectedRows2, row) {
			for _, ch := range input {
				index := int(ch) -32 
				fmt.Print(fontfile[index][row])
			}
			fmt.Println()
		}
	}
}

func loadfile18(s string) [][]string {
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

func exercise18()  {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("error")
	}
	
	fileName1 := os.Args[2]
// 	var fileName string

// 		switch fontName {
// 	case "standard":
//     fileName = "standard.txt"
// 	case "shadow":
//     fileName = "shadow.txt"
// 	case "thinkertoy":
//     fileName = "thinkertoy.txt"
// 		default:
//     fmt.Println("Unknown font:", fontName)
//     fmt.Println("Available fonts: standard shadow thinkertoy")
//     os.Exit(1)
// }

	if fileName1 == "standard" {
		fileName1 = "standard.txt"
	} else if fileName1 == "shadow" {
		fileName1 = "shadow.txt"
	} else if fileName1 == "thinkertoy" {
		fileName1 = "thinkertoy.txt"
	} else {
		fmt.Println("Unknown font")
	}
	input := os.Args[1]
	font1 := loadfile18(fileName1)
	render18(strings.ToUpper(input), font1)
}