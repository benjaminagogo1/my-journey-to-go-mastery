package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Conversion(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" {
			retrievedNum, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(retrievedNum, 10)
				words = append(words[:i], words[i+1:]...)
				i--
			}else {
				fmt.Println(err)
			}
		}
		if words[i] == "(bin)" {
			retrievedNum2, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(retrievedNum2, 10)
				words = append(words[:i], words[i+1:]...)
				i--
			}else {
				fmt.Println(err)
			}
		}
	}
	return strings.Join(words, " ")
}
