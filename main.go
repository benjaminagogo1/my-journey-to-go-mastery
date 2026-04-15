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
// 		if c >= 'a' && c <= 'f' {
// 			value = int(c-'a') + 10
// 		} else if c >= 'A' && c <= 'F' {
// 			value = int(c-'A') + 10
// 		}
// 		num = num*10 + value

// 	}

// 	fmt.Println(num)

// }

// func main() {
// 	num := 255
// 	base := 2

// 	var result string

// 	for num > 0 {
// 		bit := num % base
// 		result = fmt.Sprint(bit) + result
// 		num = num / base
// 	}

// // 	fmt.Println(result)
// // }

// package main

// import (
// 	"fmt"
// 	// "strings"
// )

// func main() {
// 	var x [5]float64
// 	x[0] = 98
// 	x[1] = 93
// 	x[2] = 77
// 	x[3] = 82
// 	x[4] = 83

// 	var total float64 = 0
// 	for i := 0; i < len(x); i++  {
// 		total += x[i]
// 	}
// 	fmt.Println(total/float64(len(x)))
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
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

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func capitalize(word string) string {
// 	if len(word) == 0 {
// 		return word
// 	}
// 	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])

// }

// func hew(s []string) string {
// 	result := []string{}

// 	for _, token := range s {

// 		switch token {
// 		case "(cap)":
// 			if len(result) > 0 {
// 				result[len(result)-1] = capitalize(result[len(result)-1])
// 			}

// 		default:
// 			result = append(result, token)
// 		}
// 	}

// 	return strings.Join(result, " ")
// }

// func main() {
// 	fmt.Println(hew([]string{"hello", "world", "welcome", "(cap)", "golang"}))
// }

// func singleUp(s string) string {
// 	words := strings.Fields(s)
// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "(up)" && i > 0 {
// 			words[i-1] = strings.ToUpper(words[i-1])
// 			words = append(words[:i], words[i+1:]...)
// 		}
// 	}
// 	r := strings.Join(words, " ")
// 	r = strings.ReplaceAll(r, " !", "!")
// 	return r
// }

// func uppcase(s string) string {
// 	words := strings.Fields(s)

// 	for i := 0; i < len(words)-1; i++ {
// 		switch words[i] {
// 		case "(up,":
// 			words[i+1] = strings.Trim(words[i+1], ")")
// 			num, err := strconv.Atoi(words[i+1])
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

// // }

// func main() {
// 	fmt.Println(uppcase("This is so exciting (up, 2)"))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"unicode"
// 	// "strings"
// 	// "unicode"
// )

// func ma(s string) string {
// 	c := []rune(s)
// 	var result []rune

// 	for i := 0; i < len(c); i++ {
// 		if unicode.IsLetter(c[i]) || unicode.IsSpace(c[i]) {
// 			result = append(result, c[i])
// 		}
// 	}
// 	r := strings.Join(strings.Fields(string(result)), " ")
// 	return r
// // }

// func count(s string) string {
// 	letters := 0
// 	b := 0
// 	for _, r := range s {
// 		if unicode.IsLetter(r) {
// 			letters++
// 		}
// 		if unicode.IsPunct(r) {
// 			b++
// 		}
// 	}
// 	out := len(strings.Fields(s))
// 	h := strings.Count(s, " ")

// 	return fmt.Sprint("Number of letters: ", letters, " \n"+"space: ", out, "\n"  + "words: ", h, "\n", "Punct: ", b)

// }

// func main() {

// 	fmt.Println(count(",, hello , ;world ! . , : benjamin ., ,. "))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	s := "she went to the (up) market"
// 	w := strings.Fields(s)
// 	for i := 0; i < len(w); i++ {
// 		w[i] = strings.Trim(w[i], "()")
// 		if w[i] == "up" {
// 			w[i-1] = strings.ToUpper(w[i-1])
// 			w = append(w[:i], w[i+1])
// 			i--
// 		}
// 	}
// 	fmt.Println(w)

// }

// package main

// import (
// 	"fmt"
// 	// "strings"
// 	"unicode"
// )

// func formatam(s string) string {
// 	var res []rune
// 	b := []rune(s)

// 	for i := 0; i < len(b); i++ {
// 		if !unicode.IsPunct(b[i]) {
// 			res = append(res, b[i])
// 		}
// 		if i > 0 && unicode.IsPunct(b[i]) && unicode.IsSpace(b[i-1]) {
// 			res = append(res, b[i])
// 		}
// 	}
// 	return string(res)

// }

// func main() {
// 	fmt.Println(formatam(",, hello , ;world ! . : benjamin ., , ."))
// }

// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// func reg(s string) string {
// 	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
// 	return re.ReplaceAllString(s, "'$1'")
// }

//	func main()  {
//		fmt.Println(reg("I am exactly how they describe me: ' awesome '"))
//		fmt.Println(reg("As Elton John said: ' I am the most well-known homosexual in the world '"))
//	}
// package main

// import (
// 	"fmt"
// )

// func count(s string) map[string]int {
// 	c := map[string]int{}
// 	//w := strings.Fields(s)
// 	for _, b := range s {
// 		c[string(b)]++
// 	}
// 	return c
// }
// func main() {
// 	fmt.Println(count("emmanuelmylevelcoiiimmitteeemmanuelmygod"))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )
// func hee(s string) string {
// 	words := strings.Split(s, "'")
// 	for i := 1; i < len(words); i ++ {
// 		words[i] = strings.TrimSpace(words[i])
// 	}
// 	return strings.Join(words, "'")
// }

// func main()  {
// 	fmt.Printf("%q", hee("As Elton John said: ' I am the most well-known homosexual in the world '"))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func ana(a, b string) bool {
// 	a = strings.ToLower(a)
// 	b = strings.ToLower(b)

// 	for _, r := range b {
// 		if !strings.Contains(string(a), string(r)) {
// 			return false
// 		}
// 	}
// 	return true
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func feg(s string) string {
// 	words := strings.Fields(s)

// 	for i := 0; i < len(words)-1; i++ {
// 		isVowels := strings.ContainsAny("aeiouhAEIOUH", string(words[i+1][0]))
// 		if words[i] == "A" && isVowels {
// 			words[i] = "An"
// 		} else if words[i] == "a" && isVowels {
// 			words[i] = "an"
// 		} else if words[i] == "An" && !isVowels {
// 			words[i] = "A"
// 		} else if words[i] == "an" && !isVowels {
// 			words[i] = "a"
// 		}
// 	}
// 	return strings.Join(words, " ")
// // }

// package main

// import (
// 	"fmt"
// 	"strings"
// 	// "unicode"
// )

// package main

// import (
// 	"os"
// 	"fmt"
// )

// func main()  {
// 	if len(os.Args) != 3 {
// 		fmt.Println("Error: Usage: go run . input.txt output.txt")
// 		return
// 	}

// 	inputFile := os.Args[1]
// 	outputFile := os.Args[2]

// 	if inputFile == outputFile {
// 		fmt.Println("Error: Inputfile  and outputFile can not be thesame",)
// 		return
// 	}

// 	input, err := os.ReadFile(inputFile)
// 	if err != nil {
// 		fmt.Println("Error: failed to read inputFile", err)
// 		return
// 	}

// 	result := processed(string(input))

// 	err = os.WriteFile(outputFile, []byte(result), 0644)
// 	if err != nil {
// 		fmt.Println("Error: failed to write outputFile", err)
// 		return
// 	}
// }

package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestArticle(t *testing.T)  {
	Input := Conversion("1E (hex) files were added")
	Expected := "30 files were added"
	if Input != Expected {
		t.Errorf("TestArticle failde: got: %s,  expected %s", Input, Expected)
	}
	t.Log("result:", Expected)
}

func conversion(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++{
		switch words[i] {
		case "(hex)":
			n, err  := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(n, 10)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}
	}
	return strings.Join(words, " ")
}