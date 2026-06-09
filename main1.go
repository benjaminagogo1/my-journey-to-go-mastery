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



package main

import (
	"fmt"
)

type accountDetail struct {
	acctName    string
	acctNumber  int
	acctBalance float64
}

func bankTransc(dep float64, acct accountDetail) {

	if dep <= 0 {
		fmt.Println("Deposit amount must be more than zero.")
		return
	}
	newbal := dep + acct.acctBalance
	fmt.Printf("Current Balance %.2f\n", newbal)

}

func Exercise2() {
	fmt.Println("Do you want deposit or withdrawal?\n1: = widthrawal\n2: = Deposit.")

	var choice int
	fmt.Scan(&choice)

	acct := accountDetail{
		acctName:    "Benjamin Agogo",
		acctNumber:  9060561633,
		acctBalance: 8000.00,
	}

	if choice == 1 {
		fmt.Println("widthrawalAmount")

		var widthrawalAmount float64
		fmt.Scan(&widthrawalAmount)

		if widthrawalAmount > acct.acctBalance {
			fmt.Println("Insufficient balance")
			return
		}

		withdrawal(widthrawalAmount, acct)

	} else if choice == 2 {
		fmt.Println("Enter the deposit amount: ")

		var amount float64

		fmt.Scan(&amount)
		realValue := amount

		bankTransc(realValue, acct)
	} else {
		fmt.Println("Invalid input. Please, choose between 1 and 2")
	}
}

func withdrawal(wid float64, acct accountDetail) {
	if wid <= 0 {
		fmt.Println("Widthrawal amount must be more than zero")
		return
	}

	balanceAfterDrawal := acct.acctBalance - wid

	fmt.Printf("Your balance: %.2f\n", balanceAfterDrawal)
}















package main

import (
	"fmt"
	"net/http"
	"strings"
)
func ParseFormInput(r *http.Request) (text string, banner string, err error) {
    err = r.ParseForm()
    if err != nil {
    err = fmt.Errorf("failed to parse form %v", err)
    return
    }
    bannerValue := r.FormValue("banner")
    textValue := r.FormValue("text")
    if bannerValue == "" {
    bannerValue = "standard"
    }
    textValue = strings.TrimSpace(textValue)
    if textValue == "" {
    err = fmt.Errorf("invalid input %s", textValue)
    return
        }
    if bannerValue != "standard" && bannerValue != "shadow" && bannerValue != "thinkertoy" {
    err = fmt.Errorf("invalid %s", bannerValue)
    return
    }
    return textValue, bannerValue, nil
    }






package main

import (
	"fmt"
	"net/http"
)

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	if message == "" {
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(statusCode)
	fmt.Fprint(w, message)

}








package main

import (
	"net/http"
)

func ServeStaticFiles(mux *http.ServeMux, dir string)  {
	if dir == "" {
		return
	}
	fs := http.FileServer(http.Dir(dir))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}






package main

import (
	"fmt"
	"html"
)

func BuildResponse(artOutput string) string {
	if artOutput == "" {
		return "<html><body></body></html>"
	}
	artOutput = html.EscapeString(artOutput)
	return fmt.Sprintf("<html><body><pre>%s</pre></body></html>", artOutput)
}





package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type TemplateCache struct {
	templ *template.Template
}

func NewTemplateCache(dir string) (*TemplateCache, error) {
	tem, err := template.ParseGlob(dir + "/*.html")
	if err != nil {
		return nil, err
	}

	templatVal := TemplateCache{
		templ: tem,
	}
	return &templatVal, nil
}

func (t *TemplateCache) Render(w http.ResponseWriter, name string, data any) error {
	if t.templ.Lookup(name) == nil {
		return fmt.Errorf("template %s is not found", name)
	}
	err := t.templ.ExecuteTemplate(w, name, data)
	if err != nil {
		return err
	}
	return nil
}


package main

import "net/http"



func SetResponseHeaders(w http.ResponseWriter, contentType string, statusCode int)  {
	if contentType == "" {
		contentType = "text/plain"
	}
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)

}



















package main

import (
	"fmt"
	"net/http"
)

func BannerSelectHandler(w http.ResponseWriter, r *http.Request) {
	BannerStyle := r.Header.Get("X-Banner-Style")

	if BannerStyle == "" {
		BannerStyle = "standard"
	}

	if BannerStyle != "standard" && BannerStyle != "shadow" && BannerStyle != "thinkertoy" {
		http.Error(w, "invalid banner", http.StatusBadRequest)
		return
	}

	font := Loadbanner("banners/" + BannerStyle + ".txt")
	if font == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Ok")
}



package main

import (
	"fmt"
	"net/http"
	"strings"
)

func ValidateContentType(r *http.Request, expected string) error {
	headerValue := r.Header.Get("Content-Type")
	if headerValue == "" {
		return fmt.Errorf("empty header")
	}
	if !strings.EqualFold(headerValue, expected) {
		return fmt.Errorf("mismatched value")
	}
	return nil

}














package main

import (
	"html/template"
	"log"
	"net/http"
)

type application struct {
	temple *template.Template
}


func RunApp(port string, templateDir string, staticDir string) error {

	router := http.NewServeMux()

	router.HandleFunc("/", HomeHandle)
	router.HandleFunc("/ascii-art", AsciiArtHandle)


	te, err := NewTemplateCache(templateDir)
	if err != nil {
		log.Fatal(err)
		return err
	}


	app := application {
		temple: te,
	}
	

	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	return http.ListenAndServe(port, LogRequest(SetDefaultHeaders(router)))
}



func HomeHandle(w http.ResponseWriter, r *http.Request)  {
	
}

func AsciiArtHandle(w http.ResponseWriter, r *http.Request)  {
	
}

package main

import (
	"net/http"
)

func StartServer(port string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandle)
	mux.HandleFunc("/ascii-art", handleAscii)
	return http.ListenAndServe(port, mux)
}

func indexHandle(w http.ResponseWriter, r *http.Request)  {

	
}

func handleAscii(w http.ResponseWriter, r *http.Request)  {

}