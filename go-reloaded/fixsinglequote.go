package main

import (
	"regexp"
)

func fixSingleQuote(s string) string {
	res := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return res.ReplaceAllString(s, "'$1'")
}
