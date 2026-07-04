package main

import (
	"fmt"
	"os"
	"strings"

)

func main()  {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --align=<type> something standard")
		return
	}

	fileName := "standard.txt"

	alignType := os.Args[1]
	userInput := os.Args[2]

	if len(os.Args) == 4 {
		fileName = os.Args[3]
		if !strings.HasSuffix(fileName, ".txt") {
			fileName += ".txt"
		}
	}

	CorrectAlignType, err := splitFlag(alignType)
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --align=<type> something standard")
		return
	}
	loadebannerValue := Loadbanner(fileName)
	switch CorrectAlignType {
	case "right":
		fmt.Println(renderArtToRight(userInput, loadebannerValue))
	case "left":
		fmt.Println(renderArtToLeft(userInput, loadebannerValue))

	case "center":
		fmt.Println(centerAsciiArt(userInput, loadebannerValue))
	case "justify":
		fmt.Print(justifyAsciiArt(userInput, loadebannerValue))
	default:
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --align=<type> something standard")
		fmt.Println("Replace <type> with any of these; `right`, `center`, `left`")
	}

}

