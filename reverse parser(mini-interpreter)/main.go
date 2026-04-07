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
		

		for i, r := range ret {
			if r == "-" || r == "+" || r == "*" || r == "/" {


				if len(stack) < 2 {
					fmt.Println("Error: Not enough input.")
					continue
				} else {
					b := r[i]
					num := 0
					for i := 0; i < len(r); i++ {
						if b != '0' && b != '1' {
							fmt.Println("Error: invalid binary input")
							break
						}
						num = num*2 + int(b - '0')
					}
					fmt.Println(num)
				}
				right := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				left := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				var result int 
				switch r {
				case "+":
					result := left + right
					fmt.Println(result)
				case "-":
					result := left - right
					fmt.Println(result)

				case "*":
					result := left * right
					fmt.Println(result)
				case "/":
					result := left / right
					fmt.Println(result)
				}
				stack = append(stack, result)
			}
		}

	}

}



