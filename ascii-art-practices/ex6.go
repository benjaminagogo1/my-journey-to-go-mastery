package main

import (
	"fmt"
	"strings"
)

func exercise6() {
	raw := " _\n| |\n|_|\n \n \n \n \n "
	content := strings.Split(raw, "\n")

	for i, part := range content {
		fmt.Printf("[%d] %s\n", i, part)
	}
}
