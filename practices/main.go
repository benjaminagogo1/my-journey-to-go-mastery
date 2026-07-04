package main

import (
	"fmt"
	"os"
)

func main()  {
	for i := 0; i < 26; i++{
		fmt.Printf("%c = %c\n", 'A'+i, 'a'+i)
	}

	for t := 'a'; t <= 'z'; t++{
		fmt.Printf("%c\n", t)
	}

	fmt.Println((os.Stdout.Fd()))
	fmt.Println(os.Stderr.Fd())
	fmt.Println(os.Stdin.Fd())
}

