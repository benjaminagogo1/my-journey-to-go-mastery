package main

import (
	"fmt"
	"strings"
)

func fixQuote(str string) string {

	words := strings.Split(str, "'")
	fmt.Println(words)

	for i := 0; i < len(words); i += 2 {
		fmt.Print(i)

		words[i] = strings.TrimSpace(words[i])
	}
	return strings.Join(words, "'")
}

func doubleQuote(st string) string {
	words := strings.Split(st, `"`)

	for i := 0; i < len(words); i += 2 {
		words[i] = strings.TrimSpace(words[i])
	}
	return strings.Join(words, `"`)
}

func main() {
	str := "Hello   'world'    today"
	st := "   one   'two'   three   'four'   "
	b := "sfrb '' ijf"
	a := "hello '   world   '"
	c := "'a''b'"
	r := "' hello"
	k := "' hello'world"
	bt := "hello'"
	fmt.Println(fixQuote(str))
	fmt.Println(fixQuote(st))
	fmt.Println(fixQuote(b))
	fmt.Println(fixQuote(a))
	fmt.Println(fixQuote(c))
	fmt.Println(fixQuote(r))
	fmt.Println(fixQuote(k))
	fmt.Println(fixQuote(bt))
}

// hello'world
