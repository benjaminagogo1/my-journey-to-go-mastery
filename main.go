// package main

// import (
// 	"fmt"
// )

// func main() {
// 	s := "255"
// 	num := 0
// 	for i := 0; i < len(s); i++ {
// 		c := s[i]
// 		value := 0
// 		if c != '0' && c != '1' {
// 			value = int(c - '0')
// 		}
// 		// if c >= 'a' && c <= 'f' {
// 		// 	value = int(c-'a') + 10
// 		// } else if c >= 'A' && c <= 'F' {
// 		// 	value = int(c-'A') + 10
// 		// }
// 		num = num*10 + value

// 	}

// 	fmt.Println(num)

// }

// func mai() {
// 	num := 255
// 	base := 2

// 	var result string

// 	for num > 0 {
// 		bit := num % base
// 		result = fmt.Sprint(bit) + result
// 		num = num / base
// 	}

// 	fmt.Println(result)
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func hew(s []string) string {
// 	var result []string
// 	for i, words := range s {
// 		if s[i] == "(cap)" && i > 0 {
// 			words = result[i-1]
// 			result[i-1] = strings.ToUpper(string(words[0])) + strings.ToLower(words[1:])
// 			continue
// 		}
// 		result = append(result, words)
// 	}
// 	return strings.Join(result, " ")
// }

// func main() {
// 	fmt.Println(hew([]string{"hello", "world", "welcome", "(cap)", "golang"}))
// }

package main

import (
	"fmt"
	"strings"
)

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

func hew(s []string) string {
	result := make([]string, 0, len(s))

	for _, token := range s {

		switch token {
		case "(cap)":
			if len(result) > 0 {
				result[len(result)-1] = capitalize(result[len(result)-1])
			}

		default:
			result = append(result, token)
		}
	}

	return strings.Join(result, " ")
}

func main() {
	fmt.Println(hew([]string{"hello", "world", "welcome", "(cap)", "golang"}))
}
