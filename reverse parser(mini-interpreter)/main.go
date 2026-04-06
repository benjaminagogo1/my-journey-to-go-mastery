package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


var val int


func main()  {
	scanner := bufio.NewScanner(os.Stdin)

	me := 0
	
	

	for {

		scanner.Scan()
		te := scanner.Text()
		te = strings.TrimSpace(te)

		if te == "exit" {
			fmt.Println("Goodbye........!!")
			break
		}

		if te == "" {
			fmt.Println("Error: Empty input")
			continue
		}

		ret := strings.Fields(te)

		if len(ret) < 2 {
			fmt.Println("Not enough input")
		}
		

		for i := 0; i < len(te); i++ {
			c := te[i]

			if c < '0' || c > '9' {
				fmt.Println("invalid decimal")
				break
			}
			me = me * 10 + int(c - '0')
		}
		fmt.Println(me)

		c := te[i]
		if c != 0 && c != 1 {}
		


	}

}



for i := 0; i < len(te); i++ {
	c := te[i]

	switch {
	case c < 0 || c > 9:
		fmt.Println("invalid decimal")

	
	}
}