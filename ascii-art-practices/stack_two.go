package main

func stackTwo(top []string, bottom []string) []string {
	result := make([]string, len(top)+len(bottom))
	copy(result, top)
	copy(result[len(top):], bottom)
	return result

}

func stackAll(block [][]string) []string {
	result := []string{}

	for _, v := range block {
		result = append(result, v...)
	}
	return result

}
