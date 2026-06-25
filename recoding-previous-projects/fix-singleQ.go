package main

import "strings"


func fixQuote(str string)  string {
	words := strings.Split(str, "'")
	
	for i := 0; i < len(words); i+=2 {
		words[i] = strings.TrimSpace(words[i])
	}
	return strings.Join(words, "'")
}