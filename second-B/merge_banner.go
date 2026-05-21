package main

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	res := make(map[rune][]string)
	for i, v := range base {
		res[i] = v
	}
	for i, v := range priority {
		res[i] = v
	}
	return res
}
