package main

import (
	"os"
	"strings"
)

func loadbanner(fileName string) [][]string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil
	}

	content := string(file)

	output := [][]string{}

	blocks := strings.Split(content, "\n\n")

	for _, block := range blocks {
		rows := strings.Split(block, "\n")
		output = append(output, rows)
	}
	return output
}
