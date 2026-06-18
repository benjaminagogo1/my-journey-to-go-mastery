package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func LoadBanner(str string) [][]string {
	file, err := os.ReadFile(str)
	if err != nil {
		log.Fatal(err)
	}
	strContent := string(file)

	output := [][]string{}

	if len(strContent) == 0 {
		fmt.Println("Bad File Format")
		return nil
	}

	blocks := strings.Split(strContent, "\n\n")
	for _, block := range blocks {
		if block == "" {
			fmt.Println("Incomplete FontFile")
			return nil
		}
		rows := strings.Split(block, "\n")
		output = append(output, rows)
	}
	return output
}
