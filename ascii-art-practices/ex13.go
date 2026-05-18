package main

import (
	"fmt"
	"os"
	// "strings"
)

func exercise13() {
	fileName := os.Args[1]
	render13(fileName)
}

func render13(input string) {
	run := []rune(input)

	for i := len(run) - 1; i >= 0; i-- {
		fmt.Print(string(run[i]))
	}
	fmt.Println()

}
