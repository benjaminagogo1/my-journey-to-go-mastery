package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func convertToDecimal() {
	s := "255"
	num := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		value := 0
		if c != '0' && c != '1' {
			value = int(c - '0')
		}
		if c >= 'a' && c <= 'f' {
			value = int(c-'a') + 10
		} else if c >= 'A' && c <= 'F' {
			value = int(c-'A') + 10
		}
		num = num*10 + value

	}

	fmt.Println(num)

}

func mn() {
	num := 255
	base := 2

	var result string

	for num > 0 {
		bit := num % base
		result = fmt.Sprint(bit) + result
		num = num / base
	}

	// fmt.Println(result)
}

func m() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}

func upperCase(s []string) string {
	var result []string
	for i, words := range s {
		if s[i] == "(cap)" && i > 0 {
			words = result[i-1]
			result[i-1] = strings.ToUpper(string(words[0])) + strings.ToLower(words[1:])
			continue
		}
		result = append(result, words)
	}
	return strings.Join(result, " ")
}

/*
func main() {
	fmt.Println(hew([]string{"hello", "world", "welcome", "(cap)", "golang"}))
}
*/

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])

}

func capitalized(s []string) string {
	result := []string{}

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

/*
func main() {
	fmt.Println(hew([]string{"hello", "world", "welcome", "(cap)", "golang"}))
}
*/

func singleUp(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	r := strings.Join(words, " ")
	r = strings.ReplaceAll(r, " !", "!")
	return r
}

func uppcase(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words)-1; i++ {
		switch words[i] {
		case "(up,":
			words[i+1] = strings.Trim(words[i+1], ")")
			num, err := strconv.Atoi(words[i+1])
			if err != nil {
				fmt.Println(err)
			}
			for j := 1; j <= num; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
	}
	return strings.Join(words, " ")

}

/*
func main() {
	fmt.Println(uppcase("This is so exciting (up, 2)"))
}
*/

func ma(s string) string {
	c := []rune(s)
	var result []rune

	for i := 0; i < len(c); i++ {
		if unicode.IsLetter(c[i]) || unicode.IsSpace(c[i]) {
			result = append(result, c[i])
		}
	}
	r := strings.Join(strings.Fields(string(result)), " ")
	return r
}

func formatIt(s string) string {
	var res []rune
	b := []rune(s)

	for i := 0; i < len(b); i++ {
		if !unicode.IsPunct(b[i]) {
			res = append(res, b[i])
		}
		if i > 0 && unicode.IsPunct(b[i]) && unicode.IsSpace(b[i-1]) {
			res = append(res, b[i])
		}
	}
	return string(res)

}

// func main() {
// 	fmt.Println(formatam(",, hello , ;world ! . : benjamin ., , ."))
// }

func reg(s string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return re.ReplaceAllString(s, "'$1'")
}

/*
func main()  {
		fmt.Println(reg("I am exactly how they describe me: ' awesome '"))
		fmt.Println(reg("As Elton John said: ' I am the most well-known homosexual in the world '"))
	}
*/

func count(s string) map[string]int {
	c := map[string]int{}
	for _, b := range s {
		c[string(b)]++
	}
	return c
}

/*
func main() {
	fmt.Println(count("emmanuelmylevelcoiiimmitteeemmanuelmygod"))
}
*/

func fixQuote(s string) string {
	words := strings.Split(s, "'")
	for i := 1; i < len(words); i++ {
		words[i] = strings.TrimSpace(words[i])
	}
	return strings.Join(words, "'")
}

/*
func main()  {
	fmt.Printf("%q", hee("As Elton John said: ' I am the most well-known homosexual in the world '"))
}
*/

func anagram(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	for _, r := range b {
		if !strings.Contains(string(a), string(r)) {
			return false
		}
	}
	return true
}

func fixArticle(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words)-1; i++ {
		isVowels := strings.ContainsAny("aeiouhAEIOUH", string(words[i+1][0]))
		if words[i] == "A" && isVowels {
			words[i] = "An"
		} else if words[i] == "a" && isVowels {
			words[i] = "an"
		} else if words[i] == "An" && !isVowels {
			words[i] = "A"
		} else if words[i] == "an" && !isVowels {
			words[i] = "a"
		}
	}
	return strings.Join(words, " ")
}

/*
func main()  {
	if len(os.Args) != 3 {
		fmt.Println("Error: Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		fmt.Println("Error: Inputfile  and outputFile can not be thesame",)
		return
	}

	// input, err := os.ReadFile(inputFile)
	if err != nil {
		// fmt.Println("Error: failed to read inputFile", err)

		return
	}

	// result := processed(string(input))

	// err = os.WriteFile(outputFile, []byte(result), 0644)
	// if err != nil {
		// fmt.Println("Error: failed to write outputFile", err)
		return
	}
}
*/
/*
func TestArticle(t *testing.T)  {
	// Input := Conversion("1E (hex) files were added")
	Expected := "30 files were added"
	// if Input != Expected {
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
*/

func duplicate_count(s string) int {
	s = strings.ToLower(s)
	count := make(map[rune]int)

	for _, k := range s {
		if unicode.IsLetter(k) || unicode.IsDigit(k) {
			count[k]++
		}

	}

	duplicate := 0

	for _, r := range count {
		if r > 1 {
			duplicate++
		}
	}
	return duplicate
}

func findShortWord(s string) int {
	words := strings.Fields(s)

	shortest := len(words[0])
	for _, r := range words {
		if len(r) < shortest {
			shortest = len(r)
		}
	}
	return shortest
}

func countLowNum(d []int) []int {
	result := []int{}

	for _, ch := range d {
		count := 0
		for _, r := range d {
			if r < ch {
				count++
			}
		}
		result = append(result, count)
	}
	return result
}

func main() {
	fmt.Println(countLowNum([]int{8, 0, 1, 2, 3, 5}))
}
