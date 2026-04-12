package main

import (
	"fmt"
	"strconv"
	"strings"
)

func caseTransFromation(s string) string {

	w := strings.Fields(s)

	for i := 0; i < len(w); i++ {
		switch w[i] {

		case "(cap,":

			w[i+1] = strings.Trim(w[i+1], ")")

			n, err := strconv.Atoi(w[i+1])
			if err != nil {
				fmt.Println(err)
			}

			for j := 1; j <= n; j++ {
				w[i-j] = strings.ToUpper(string(w[i-j][0])) + strings.ToLower(w[i-j][1:len(w[i-j])])
			}
			w = append(w[:i], w[i+2:]...)
			i--

		case "(up,":

			w[i+1] = strings.Trim(w[i+1], ")")

			n, err := strconv.Atoi(w[i+1])
			if err != nil {
				fmt.Println(err)
			}
			for b := 1; b <= n; b++ {
				w[i-b] = strings.ToUpper(w[i-b])
			}
			w = append(w[:i], w[i+2:]...)
			i--

		case "(low,":
			w[i+1] = strings.Trim(w[i+1], ")")
			numb, err := strconv.Atoi(w[i+1])
			if err != nil {
				fmt.Println(err)
			}
			for x := 1; x <= numb; x++ {
				w[i-x] = strings.ToLower(w[i-x])
			}
			w = append(w[:i], w[i+2:]...)
			i--

		case "(low)":
			w[i-1] = strings.ToLower(w[i-1])
			w = append(w[:i], w[i+1:]...)
			i--

		case "(cap)":
			w[i-1] = strings.ToUpper(string(w[i-1][0])) + strings.ToLower(w[i-1][1:len(w[i-1])])
			w = append(w[:i], w[i+1:]...)
			i--

		case "(up)":
			w[i-1] = strings.ToUpper(w[i-1])
			w = append(w[:i], w[i+1:]...)
			i--

		}

	}
	return strings.Join(w, " ")

}

// func uppcase(s string) string {
// 	words := strings.Fields(s)

// 	for i := 0; i < len(words)-1; i++ {
// 		switch words[i] {
// 		case "(up,":
// 			wor := strings.Trim(words[i+1], ")")
// 			num, err := strconv.Atoi(wor)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			for j := 1; j <= num; j++ {
// 				words[i-j] = strings.ToUpper(words[i-j])
// 			}
// 			words = append(words[:i], words[i+2:]...)
// 			i--
// 		}
// 	}
// 	return strings.Join(words, " ")

// }
