package main 


func bannerWidth(str string, font map[rune][]string) int {

	count := 0

	for _, ch := range str {
		count += len(font[ch][1])
	}
	return count
}

