package main

import (
	"strconv"
	"strings"
)

func conversion(s string) string {

	r := strings.Fields(s)
	for i := 0; i < len(r); i++ {
		switch r[i] {
		case "(hex)":
			n, err := strconv.ParseInt(r[i-1], 16, 64)
			if err == nil {
				r[i-1] = strconv.FormatInt(n, 10)
				r = append(r[:i], r[i+1:]...)
				i--
			}
		case "(bin)":
			num, err := strconv.ParseInt(r[i-1], 2, 64)
			if err == nil {
				r[i-1] = strconv.FormatInt(num, 10)
				r = append(r[:i], r[i+1:]...)
				i--
			}
		}
	}
	return strings.Join(r, " ")
}
