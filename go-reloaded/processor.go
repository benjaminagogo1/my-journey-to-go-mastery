package main

import (
	"strings"
)

func processedText(s string) string {
	r := strings.Split(s, "\n")
	var new []string

	for _, k := range r {
		k = fixArticle(k)
		k = fixPunctuation(k)
		k = fixSingleQuote(k)
		k = conversion(k)
		k = caseTransFromation(k)
		k = singleUp(k)

		new = append(new, k)
	}
	return strings.Join(new, " ")

}
