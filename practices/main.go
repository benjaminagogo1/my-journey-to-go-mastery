package main 

import (
	"fmt"
)

func main()  {
	for i := 0; i < 26; i++{
		fmt.Printf("%c = %c\n", 'A'+i, 'a'+i)
	}

	for t := 'a'; t <= 'z'; t++{
		fmt.Printf("%c\n", t)
	}
}

