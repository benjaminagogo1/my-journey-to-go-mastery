package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type bookTracker struct {
	Title       string
	Author      string
	Year        int
	IsAvailable bool
}

func borrowBook(book *bookTracker) {
	book.IsAvailable = false

}

func returnBook(book *bookTracker) {
	book.IsAvailable = true
}

func displAll(book []bookTracker) {
	// view := &book
	// view2 := *view

	fmt.Println(book)
}

func addBook(addbook *[]bookTracker) {
	addToLib := []bookTracker{}

	newBook := bookTracker{
		Title:       "I Won",
		Author:      "Hope",
		Year:        2024,
		IsAvailable: true,
	}
	*addbook = append(addToLib, newBook)

}

func displayBook(book *bookTracker) {
	fmt.Printf("The Title %s: ", book.Title)
	fmt.Printf("The Author %s: ", book.Author)
	fmt.Printf("Published Year %d: ", book.Year)
	fmt.Printf("The Author %v: ", book.IsAvailable)

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	libraryData := []bookTracker{
		{
			Title:       "see at the top",
			Author:      "Zigglar",
			Year:        2002,
			IsAvailable: false,
		},

		{
			Title:       "born to win",
			Author:      "Benjamin",
			Year:        2020,
			IsAvailable: true,
		},

		{
			Title:       "succes is predictable",
			Author:      "Joseph",
			Year:        2018,
			IsAvailable: true,
		},

		{
			Title:       "small habit actions win",
			Author:      "Benjamin",
			Year:        2022,
			IsAvailable: true,
		},
	}

	fmt.Println("\n1: = Looking For A Particular Book?\n2: = Or Just Want To View Our Shelf?")
	var option1 int
	fmt.Scan(&option1)
	if option1 == 0 {
		fmt.Println("Choose at least one option")
		return
	}

	if option1 == 1 {
		fmt.Println("Enter the name of the book")
		bookName, _ := reader.ReadString('\n')
		bookName = strings.TrimSpace(bookName)
		seen := false
		for _, r := range libraryData {
			if strings.ToLower(bookName) == r.Title {
				seen = true
				fmt.Printf("The book is called %q\n", strings.ToUpper(r.Title))
				return
			}
		}
		if !seen {
			fmt.Println("The book is not available!")
			return
		}

	}
	if option1 == 2 {
		displAll(libraryData)
		return
	}
	fmt.Println("\n1: = Available books\n2: = Borrowed Books\n3: = Returned Books")

}
