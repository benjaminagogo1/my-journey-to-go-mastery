package main

import "strings"

// "strings"

func processedText(s string) string {
	r := strings.Split(s, "\n")
	var new []string

	for _, k := range r {
		k := caseTransFromation(k)
		k = fixPunctuation(k)
		k = fixArticle(k)
		k = fixSingleQuote(k)
		k = conversion(k)
		// k = singleUp(k)

		new = append(new, k)
	}
	return strings.Join(new, "\n")

}
