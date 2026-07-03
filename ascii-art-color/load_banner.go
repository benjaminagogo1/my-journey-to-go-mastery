package main

import (
	"fmt"
	"os"
	"strings"
)

func loadbanner(fileName string) map[rune][]string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	content := string(file)

	output := map[rune][]string{}

	blocks := strings.Split(content, "\n")

	for x := ' '; x <= '~'; x++ {
		index := int(x-' ') * 9
		output[x] = blocks[index+1 : index+9]
	}

	return output
}
