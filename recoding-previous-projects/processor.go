package main

import "strings"


func processor(s string) string {
	words := strings.Split(s, "\n")

	result := []string{}

	for _, value := range words {
		value := caseTransformation(value)
		value = fixArticle(value)
		value = fixQuote(value)
		value = Conversion(value)
		value = FixPunc(value)
		
		result = append(result, value)
	}
	return strings.Join(result, "\n")	
}