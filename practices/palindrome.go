package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	examples := []string{"racecar", "hello", "A man a plan a canal Panama"}

	for _, word := range examples {
		result := isPalindrome(word)
		fmt.Printf("'%s' = %t\n", word, result)	
	}
}


docker --name ascii-web ascii-art-web

docker rm
docker run -p 8080:8080 ascii-art-web