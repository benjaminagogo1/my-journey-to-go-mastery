package main

import (

	"strings"
)


func fixSingleQuote(s string) string {
	words := strings.Split(s, "'")
	for i := 1; i < len(words); i +=2 {
		words[i] = strings.TrimSpace(words[i])
	}
	return strings.Join(words, "'")
}
