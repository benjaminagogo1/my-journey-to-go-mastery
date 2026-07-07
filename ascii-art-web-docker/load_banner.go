package main

import (
	"fmt"
	"os"
	"strings"
)

func Loadbanner(filename string) (map[rune][]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file")
	}
	content := string(file)
	if len(content) == 0 {
		return nil, fmt.Errorf("empty file")
	}

	output := map[rune][]string{}

	lines := strings.Split(content, "\n")
	
	for index := 32; index <= 126; index++{
		start := (index- 32) * 9
		output[rune(index)] = lines[start+1: start+9]
	} 
	return output, nil
}
