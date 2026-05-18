package main

import (
	"fmt"
)

func exercise9() {
	hRows := []string{
		"  _   _  ",
		" | | | | ",
		" | |_| | ",
		" |  _  | ",
		" | | | | ",
		" |_| |_| ",
		"         ",
		"         ",
	}

	iRows := []string{
		"  _  ",
		" (_) ",
		"  _  ",
		" | | ",
		" | | ",
		" |_| ",
		"     ",
		"     ",
	}
	// //  p := []string{}
	// for i := 0; i < len(hRows) && i < len(iRows); i++ {
	// 	 p = append(p, hRows[i], iRows[i], "\n")
	// }
	// fmt.Println(p)

	// // OR //

	// for i := 0; i < len(hRows) && i < len(iRows); i++ {
	// 	combined := hRows[i] + iRows[i]
	// 	fmt.Println(combined)
	// }

	// OR //

	p := []string{}
	for i := 0; i < len(hRows) && i < len(iRows); i++ {
		combined := hRows[i] + iRows[i]
		p = append(p, combined)
	}
	for _, row := range p {
		fmt.Println(row)
	}
}
