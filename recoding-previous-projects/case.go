package main

import (
	"fmt"
	"strconv"
	"strings"
)

func caseTransformation(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(up,":
			words[i+1] = strings.Trim(words[i+1], ")")
			fmt.Printf("%T\n", words[i+1])
			trimmedNum, err := strconv.Atoi(words[i+1])
			fmt.Printf("%T\n", trimmedNum)
			if err != nil {
				fmt.Println(err)
				return ""
			}
			for x := 1; x <= trimmedNum; x++ {
				words[i-x] = strings.ToUpper(words[i-x])
			}
			words = append(words[:i], words[i+2:]...)
			i--

		case "(cap,":
			words[i+1] = strings.Trim(words[i+1], ")")
			retrievedNum2, err := strconv.Atoi(words[i+1])
			if err != nil{
				fmt.Println(err)
				return ""
			}
			for k := 1; k <= retrievedNum2; k++{
				words[i-k] = strings.Title(strings.ToLower(words[i-k]))
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
		
	}
	return strings.Join(words, " ")
}

