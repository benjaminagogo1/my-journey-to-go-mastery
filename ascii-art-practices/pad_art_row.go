package main

import "strings"

func padArtRow(rows []string, width int) []string {
	if width <= 0 {
		return rows
	}

	result := make([]string, len(rows))

	for i, v := range rows {
		padding := width - len(v)
		if padding > 0 {
			result[i] = v + strings.Repeat(" ", padding)
		} else {
			result[i] = v
		}
	}
	return result

}
