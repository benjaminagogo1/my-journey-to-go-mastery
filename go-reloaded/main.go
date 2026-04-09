package main

import (
	"fmt"
	"strconv"
	"strings"
)

func cap(s string) []string {

	w := strings.Fields(s)

	for i := 0; i < len(w); i++ {
		switch w[i] {

		case "(cap,":
			w[i+1] = strings.Trim(w[i+1], ")")

			n, err := strconv.Atoi(w[i+1])
			if err != nil {
				fmt.Println(err)
			}

			for j := 1; j <= n; j++ {
				w[i-j] = strings.ToUpper(string(w[i-j][0])) + strings.ToLower(w[i-j][1:len(w[i-j])])
			}
			w = append(w[:i], w[i+2:]...)
			i--
		case "(low,":
			w[i+1] = strings.Trim(w[i+1], ")")
			numb, err := strconv.Atoi(w[i+1])
			if err != nil {
				fmt.Println(err)
			}
			for x := 1; x <= numb; x++ {
				w[i-x] = strings.ToLower(w[i-x])
			}
			w = append(w[:i], w[i+1:]...)

		case "(low)":
			b := strings.ToLower(w[i-1])
			fmt.Println(b)
		}

	}
	return w

}

func conversion(s []string) []string {
	b := strings.Join(s, " ")

	r := strings.Fields(b)
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
	return r
}

func main() {
	fmt.Println(cap("at any instance of (cap, 2) i want you to capitalize (cap, 3)"))
	fmt.Println(conversion([]string{"1E", "(hex)", "files", "were", "added", "for", "10", "(bin)", "year", "WORK", "(low)"}))
}
