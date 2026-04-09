package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: Usage go run . <input.txt><output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		fmt.Println("Error: inputfile and outputfile can not be the same")
		return
	}

	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error: Failed to read inputFile", err)
		return

	}

	processedText := string(input)

	err = os.WriteFile(outputFile, []byte(processedText), 0644)
	if err != nil {
		fmt.Println("Error: failed to write outputfile.", err)
	}
	fmt.Printf("Succesfully processed! %s -> %s\n", inputFile, outputFile)

}
