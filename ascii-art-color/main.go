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
		fmt.Println("Usage: go run . [OPTION] [STRING]")
        fmt.Println("EX: go run . --color=<color> <substring to be colored> something")
		return
	}

	if flag[1] == "" {
		fmt.Println("invalid color")
		fmt.Println("Usage: go run . [OPTION] [STRING]")
        fmt.Println("EX: go run . --color=<color> <substring to be colored> something")
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
	colorName = value
	
	switch {
	case len(os.Args) == 3:
		colorName = value
		text = os.Args[2]
		substr = text

	case len(os.Args) == 4:
		colorName= value
		substr = os.Args[2]
		text = os.Args[3]
		
	case len(os.Args) == 5:
		colorName = value
		substr = os.Args[2]
		text = os.Args[3]
		banner = os.Args[4]
	}

	if banner != "standard.txt" &&
		banner != "shadow.txt" &&
		banner != "thinkertoy.txt" {
		fmt.Println("invalid banner input")
		return
	}


	loadbannerValue := loadbanner(banner)
	fmt.Println(render(colorName, substr, text, loadbannerValue))


}
