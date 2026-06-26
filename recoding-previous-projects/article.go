package main

import (
	"fmt"
	// "net/http"
	"strings"
)

// func fixArticle(str string) string {
// 	words := strings.Fields(str)

// 	for i := 0; i < len(words)-1; i ++ {
// 		isVowel := strings.ContainsAny("AEIOUHaeiouh", string(words[i+1][0]))
// 		if words[i] == "A" && isVowel {
// 			words[i] = "An"
// 		}else if words[i] == "An" && !isVowel {
// 			words[i] = "A"
// 		}else if words[i] == "a" && isVowel {
// 			words[i] = "an"
// 		}else if words[i] == "an" && !isVowel {
// 			words[i] = "a"
// 		}
// 	}
// 	return strings.Join(words, " ")
// }

func fixArticle(str string) string {
	words := strings.Fields(str)
	for i := 0; i < len(words); i++ {
		isVowel := strings.ContainsAny("AOIUEHaeoiuh", string(words[i+1][0]))
		switch {
		case isVowel && words[i] == "A":
			words[i] = "An"
		case !isVowel && words[i] == "An":
			words[i] = "A"
		case isVowel && words[i] == "a":
			words[i] = "an"
		case !isVowel && words[i] == "an":
			words[i] = "a"
		}
	}
	return strings.Join(words, " ")
}

func mai() {
	// fmt.Println(fixArticle("a awesome"))
	// fmt.Println(fixQuote(''))
	// fmt.Println(Conversion("10 (bin)"))
	// fmt.Println(Conversion("1E (hex)"))
	fmt.Println(FixPunc("hello  ? .  ,  "))
	// http.HandleFunc("/", HomeHandle)
	// http.HandleFunc("/count", countHandle)
	// http.HandleFunc("/come", comeHandle)
	// fmt.Println("Server is Live....")
	// http.ListenAndServe(":8080", nil)
	// fmt.Println(caseTransformation("This is so exciting (up, 2)"))
	// fmt.Println(caseTransformation("i am jhfine hgnmncare exciting (cap, 5)"))

}
