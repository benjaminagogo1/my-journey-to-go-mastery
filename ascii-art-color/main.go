// package main

// import (
// 	"fmt"
// 	// "os"
// 	"strings"
// )

// func main()  {
// 	re := "--color="
// 	slic := strings.Split(re, "=")
// 	fmt.Println(slic)
// 	fmt.Println(len(slic))
// 	fmt.Println(slic[0])
// 	fmt.Println(slic[1])
// 	fmt.Println(len(re))

// }

package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	banner    = "standadrd.txt"
	text      string
	substr    string
	colorName string
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		return
	}
	fileName := "standard.txt"

	colors := map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"blue":   "\033[34m",
		"yellow": "\033[33m",
	}

	userColor := os.Args[1]

	flag := strings.Split(userColor, "=")

	if len(flag) != 2 {
		fmt.Println("invalid color format")
		return
	}

	if flag[1] == "" {
		fmt.Println("invalid color")
		return
	}

	realColorName := flag[1]


	value, ok := colors[realColorName]

	if !ok {
		fmt.Println("invalid color")
		return
	}

	banner = "standard.txt"
	text = ""
	substr = ""
	colorName = realColorName

	switch {
	case len(os.Args) == 3:
		colorName = realColorName
		text = os.Args[2]

		loadbannerValue := loadbanner(fileName)
		
		fmt.Println(value, ok)
		fmt.Println(text)
		fmt.Println(banner)
		fmt.Println(render(colorName, substr, text, loadbannerValue))

	case len(os.Args) == 4:
		colorName= realColorName
		substr = os.Args[2]
		text = os.Args[3]
		loadbannerValue := loadbanner(fileName)
		fmt.Println(render(colorName, substr, text, loadbannerValue))

	case len(os.Args) == 5:
		colorName = realColorName
		substr = os.Args[2]
		text = os.Args[3]
		fmt.Println(text)
		banner = os.Args[4]
		loadbannerValue := loadbanner(fileName)
		fmt.Println(render(colorName, substr, text, loadbannerValue))

	}

}
