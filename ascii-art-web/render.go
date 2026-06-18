package main

import "strings"

func Render(str string, font [][]string) string {
	if str == "" {
		return str
	}
	str = strings.ReplaceAll(str, `\n`, "\n")

	var result strings.Builder

	splitStr := strings.Split(str, "\n")

	for _, word := range splitStr {
		for row := 0; row < 8; row++ {
			for _, ch := range word {
				index := int(ch) - 32
				if index < 0 || index > 94 {
					continue
				}
				if index >= len(font) {
					continue
				}
				result.WriteString(font[index][row])
			}
			result.WriteString("\n")
		}

	}
	return result.String()
}

/*
package main

import (
	"os"
	"strings"
)

func main() {
	r := strings.NewReplacer("<name>", "Alice", "<role>", "Engineer")
	template := "Hello <name>, you are logged in as an <role>.\n"

	// Writes the replaced text directly to standard output without allocating a new string
	r.WriteString(os.Stdout, template)
}


package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read everything sent through the pipe line by line
	for scanner.Scan() {
		line := scanner.Text()
		// Uppercase everything passing through
		fmt.Println(strings.ToUpper(line))
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Wrap os.Stdin in a scanner to easily read full lines of text
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your name: ")

	// Program pauses here, waiting for user to type and hit Enter
	if scanner.Scan() {
		input := scanner.Text()
		fmt.Printf("Hello, %s!\n", input)
	}
}


package main

import (
	"fmt"
	"os"
)

func main() {
	// Normal program output goes to Stdout
	fmt.Fprintln(os.Stdout, "SUCCESS: The operation completed perfectly.")

	// Error or diagnostic info goes to Stderr
	fmt.Fprintln(os.Stderr, "ERROR: Connection failed. Retrying...")
}


package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1. The standard way (uses os.Stdout implicitly under the hood)
	fmt.Println("Hello World 1")

	// 2. Explicitly writing to os.Stdout using the fmt package
	fmt.Fprintf(os.Stdout, "Hello World %d\n", 2)

	// 3. Low-level direct write (requires converting string to a byte slice)
	io.WriteString(os.Stdout, "Hello World 3\n")
}

*/
