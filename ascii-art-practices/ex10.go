package main

import (
	"fmt"
	"os"
	"strings"
)

func exercise10() {
	fileName := "standard.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	content := string(file)

	result := [][]string{}

	blocks := strings.Split(content, "\n\n")

	for _, block := range blocks {
		rows := strings.Split(block, "\n")
		result = append(result, rows)

	}

	fmt.Println("Number of blocks: ", len(blocks))
	fmt.Println("First row of block 33: ", strings.Join(result[33], "\n"))

}





