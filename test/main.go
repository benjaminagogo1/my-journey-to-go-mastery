
package main

import (
	"fmt"
	// "os"
	"strings"
)

func main()  {
	re := "--color="
	slic := strings.Split(re, "=")
	fmt.Println(slic)
	fmt.Println(len(slic))
	fmt.Println(slic[0])
	fmt.Println(slic[1])
	fmt.Println(len(re))


	
	colors := map[string]string{
		"red" : "\033[31m",
		"green" : "\033[32m",
		"blue" :  "\033[34m",
		"yellow" : "\033[33m",
	}
	

	userColor := "--color=white"

	flag := strings.SplitN(userColor, "=", 2)

	if len(flag) != 2 {
		fmt.Println("invalid color format")
		return
	}

	if flag[1] == "" {
		fmt.Println("invalid color")
		return
	}

	realColorName := flag[1]

	_, ok := colors[realColorName]
	fmt.Println(ok)
}

