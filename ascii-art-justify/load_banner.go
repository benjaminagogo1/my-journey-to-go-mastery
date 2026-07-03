package main

import (
	"fmt"
	"os"
	"strings"
)

func Loadbanner(fileName string) map[rune][]string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	content := string(file)

	mapping := map[rune][]string{}

	blocks := strings.Split(content, "\n")
	
	for index := 32; index <= 126; index ++ {
		start := (index - 32) * 9 
		mapping[rune(index)] = blocks[start+1 : start+9]
	}
	return mapping
}