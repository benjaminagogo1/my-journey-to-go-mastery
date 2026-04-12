package main

import (
	"strings"
)

func fixArticle(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words)-1; i++ {
		isVowels := strings.ContainsAny("aeiouhAEIOUH", string(words[i+1][0]))
		if words[i] == "A" && isVowels {
			words[i] = "An"
		} else if words[i] == "An" && !isVowels {
			words[i] = "A"
		} else if words[i] == "a" && isVowels {
			words[i] = "an"
		} else if words[i] == "an" && !isVowels {
			words[i] = "a"
		}
	}
	return strings.Join(words, " ")

}
