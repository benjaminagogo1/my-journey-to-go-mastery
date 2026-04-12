package main

import (
	"regexp"
)

func fixPunctuation(s string) string {
	res := regexp.MustCompile(`\s*([.,?:;!]+)\s*`)
	output := res.ReplaceAllString(s, "$1 ")
	// output = fmt.Sprintf("%q\n", output)
	return output
}
