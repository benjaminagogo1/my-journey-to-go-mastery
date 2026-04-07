package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func main()  {
	scanner := bufio.NewScanner(os.Stdin)

	var stack [] int
	
	

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
		

		for _, r := range ret {
			if r == "-" || r == "+" || r == "*" || r == "/" {


				if len(stack) < 2 {
					fmt.Println("Error: Not enough input.")
					break
				} 
				
				right := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				left := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				var result int 

				switch r {
				case "+":
					result = left + right
				
				case "-":
					result = left - right
					

				case "*":
					result = left * right
				
				case "/":
					result = left / right
				
				}
				stack = append(stack, result)

			} else {
				num := 0
				for j := 0; j < len(r); j++ {
					c := r[j]
					if c < '0' || c > '9' {
						fmt.Println("Error: Invalid decimal")
					
					}
					num = num*10 + int(c - '0')
				}
				stack = append(stack, num)

			}

		}

	}

}



