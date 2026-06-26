package main

import (
	// "fmt"
	"strings"
)



func fixQuote(str string) string {

	words := strings.Split(str, "'")

	for i := 0; i < len(words); i ++ {

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



// func main() {
	// str := "Hello   'world'    today"
	// st := "   one   'two'   three   'four'   "
	// b := "sfrb '' ijf"
	// a := "hello '   world   '"
	// c := "'a''b'"
	// r := "' hello"
// 	// k := "' hello'world"
// 	// bt := "hello'"
// 	tee := `'a''b'`
// 	for i, r := range tee {
// 		fmt.Printf("index =%v value= %c\n", i+1, r)
// 	}
// 	// fmt.Println(fixQuote(str))
// 	// fmt.Println(fixQuote(st))
// 	// fmt.Println(fixQuote(b))
// 	// fmt.Println(fixQuote(a))
// 	// fmt.Println(fixQuote(c))
// 	// fmt.Println(fixQuote(r))
// 	// fmt.Println(fixQuote(k))
// 	// fmt.Println(fixQuote(bt))
// 	// fmt.Println(len(tee))
// 	fmt.Println(fixQuote("As Elton John said: ' I am the most well-known homosexual in the world '"))
// }

// hello'world







func trimOutsideQuotes(text string) string {
	if strings.Count(text, "'")%2 != 0 {
		return ""
	}

	segments := strings.Split(text, "'")

	for i := 0; i < len(segments); i += 2 {
		segments[i] = strings.TrimSpace(segments[i])
	}

	return strings.Join(segments, "'")
}





// if i%2 == 0 {
//     parts[i] = strings.TrimSpace(parts[i])
// } else {
//     parts[i] = strings.ReplaceAll(parts[i], "\t", " ")
// }





// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func trimOutside(text, quote string) string {
// 	segments := strings.Split(text, quote)

// 	for i := 0; i < len(segments); i += 2 {
// 		segments[i] = strings.TrimSpace(segments[i])
// 	}

// 	return strings.Join(segments, quote)
// }

// func main() {
// 	text := `As Elton John said:" Hello World "`

// 	result := trimOutside(text, `"`)
// 	fmt.Println(result)
// }