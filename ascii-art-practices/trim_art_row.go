package main

import (
	"strings"
)

func trimArtRow(rows []string) []string {
	result := make([]string, len(rows))

	for i, row := range rows {
		result[i] = strings.TrimRight(row, " ")
	}
	return result
}
