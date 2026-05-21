package main

func StackTwo(top []string, bottom []string) []string {
	res := make([]string, len(top)+len(bottom))
	copy(res, top)
	copy(res[len(top):], bottom)

	return res
}

func StackAll(blocks [][]string) []string {
	res := []string{}
	for _, v := range blocks {
		res = append(res, v...)
	}
	return res
}

