package main

import (
	"fmt"
	"os"
)

func renderEx12(input string) {
	for _, ch := range input {
		index := int(ch) -32
		if index < 0 || index > 94 {
			fmt.Fprintf(os.Stderr, "Warning: unsupported character %c\n", ch)
			continue
		} else {
			fmt.Printf("%c, supported\n", ch)
		}
	}

}
func niam() {
	input := os.Args[1]
	renderEx12(input)
}
