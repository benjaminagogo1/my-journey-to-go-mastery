package main

import "strings"

func TrimArtRows(rows []string) []string {
	res := make([]string, len(rows))
	for i, v := range rows {
		res[i] = strings.TrimRight(v, " ")
	}
	return res
}
