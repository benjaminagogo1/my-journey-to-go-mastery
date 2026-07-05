package main

import (
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

// func main() {
// 	// Test cases
// 	examples := []string{"racecar", "hello", "A man a plan a canal Panama"}

// 	for _, word := range examples {
// 		result := isPalindrome(word)
// 		fmt.Printf("'%s' = %t\n", word, result)
// 	}
// }
