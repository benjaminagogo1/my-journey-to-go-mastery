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







func mai() {
	s := "1E"
	num := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		value := 0
		if c >= '0' && c <= '9' {
			value = int(c - '0')
		}
		if c >= 'A' && c <= 'F' {
			value = int(c-'A') + 10
		} else if c >= 'a' && c <= 'f' {
			value = int(c-'a') + 10
		}
		num = num*16 + value

	}
	fmt.Println(num)

}