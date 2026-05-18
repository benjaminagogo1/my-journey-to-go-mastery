package main

import (
	"fmt"
	"os"
	"strings"
)

func exercise11() {
	input := os.Args[1]

	segments := strings.Split(input, "\\n")
	for i, segment := range segments {
		fmt.Printf("segment %d: %s\n", i+1, segment)
	}
}
