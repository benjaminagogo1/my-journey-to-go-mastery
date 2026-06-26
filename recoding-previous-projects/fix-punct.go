package main

import (
	"strings"
	"fmt"
)

func FixPunc(s string) string {
	words := strings.Fields(s)

	result := []string{}

	for _, char := range words {
		for len(char) > 0 && strings.ContainsAny(char[:1], ".,!:?;") {
			fmt.Println(char)
			if len(result) > 0 {
				result[len(result)-1] += char[:1]
				fmt.Println(result)
			}
			char = char[1:]
		}
		if char != "" {
			result = append(result, char)
		}
		fmt.Println(char)
		fmt.Println(result)
	}
	return strings.Join(result, " ")
}

func main()  {
	fmt.Println(FixPunc("hello  ? .  ,  "))
}