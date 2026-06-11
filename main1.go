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


package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// helper: creates a simple next handler that writes a known status.
func newStatusHandler(status int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
	})
}

// TestNewRequestLogger_ReturnsNonNil checks that NewRequestLogger
// returns a non-nil instance.
func TestNewRequestLogger_ReturnsNonNil(t *testing.T) {
	logger := NewRequestLogger()
	if logger == nil {
		t.Error("expected non-nil RequestLogger, got nil")
	}
}

// TestNewRequestLogger_EmptyLogs checks that a new logger has no
// log entries.
func TestNewRequestLogger_EmptyLogs(t *testing.T) {
	logger := NewRequestLogger()
	logs := logger.GetLogs()
	if len(logs) != 0 {
		t.Errorf("expected empty logs, got %d entries", len(logs))
	}
}

// TestMiddleware_LogsMethod checks that the request method is
// captured in the log entry.
func TestMiddleware_LogsMethod(t *testing.T) {
	logger := NewRequestLogger()
	handler := logger.Middleware()(newStatusHandler(http.StatusOK))
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	logs := logger.GetLogs()
	if len(logs) != 1 {
		t.Fatalf("expected 1 log entry, got %d", len(logs))
	}
	if logs[0].Method != "GET" {
		t.Errorf("expected method 'GET', got %q", logs[0].Method)
	}
}

// TestMiddleware_LogsPath checks that the request path is captured
// in the log entry.
func TestMiddleware_LogsPath(t *testing.T) {
	logger := NewRequestLogger()
	handler := logger.Middleware()(newStatusHandler(http.StatusOK))
	req := httptest.NewRequest("GET", "/ascii-art", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	logs := logger.GetLogs()
	if logs[0].Path != "/ascii-art" {
		t.Errorf("expected path '/ascii-art', got %q", logs[0].Path)
	}
}

// TestMiddleware_LogsStatusCode checks that the response status code
// is captured in the log entry.
func TestMiddleware_LogsStatusCode(t *testing.T) {
	logger := NewRequestLogger()
	handler := logger.Middleware()(newStatusHandler(http.StatusNotFound))
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	logs := logger.GetLogs()
	if logs[0].Status != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", logs[0].Status)
	}
}

// TestMiddleware_LogsTimestamp checks that the timestamp is set
// in the log entry.
func TestMiddleware_LogsTimestamp(t *testing.T) {
	logger := NewRequestLogger()
	handler := logger.Middleware()(newStatusHandler(http.StatusOK))
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	logs := logger.GetLogs()
	if logs[0].Time.IsZero() {
		t.Error("expected non-zero timestamp, got zero")
	}
}

// TestMiddleware_CallsNextHandler checks that the next handler is
// called after logging.
func TestMiddleware_CallsNextHandler(t *testing.T) {
	logger := NewRequestLogger()
	called := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})
	handler := logger.Middleware()(next)
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if !called {
		t.Error("expected next handler to be called, but it was not")
	}
}

// TestMiddleware_MultipleRequestsLogged checks that multiple requests
// are all captured in the log.
func TestMiddleware_MultipleRequestsLogged(t *testing.T) {
	logger := NewRequestLogger()
	handler := logger.Middleware()(newStatusHandler(http.StatusOK))
	for i := 0; i < 3; i++ {
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
	}
	logs := logger.GetLogs()
	if len(logs) != 3 {
		t.Errorf("expected 3 log entries, got %d", len(logs))
	}
}

// TestMiddleware_LogsCorrectStatusCode checks that different status
// codes are captured correctly.
func TestMiddleware_LogsCorrectStatusCode(t *testing.T) {
	cases := []int{
		http.StatusOK,
		http.StatusNotFound,
		http.StatusBadRequest,
		http.StatusInternalServerError,
	}
	for _, code := range cases {
		logger := NewRequestLogger()
		handler := logger.Middleware()(newStatusHandler(code))
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
		logs := logger.GetLogs()
		if logs[0].Status != code {
			t.Errorf("expected status %d, got %d", code, logs[0].Status)
		}
	}
}

// TestGetLogs_ReturnsAllEntries checks that GetLogs returns all
// captured entries.
func TestGetLogs_ReturnsAllEntries(t *testing.T) {
	logger := NewRequestLogger()
	handler := logger.Middleware()(newStatusHandler(http.StatusOK))
	paths := []string{"/", "/ascii-art", "/static/style.css"}
	for _, path := range paths {
		req := httptest.NewRequest("GET", path, nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
	}
	logs := logger.GetLogs()
	if len(logs) != 3 {
		t.Errorf("expected 3 log entries, got %d", len(logs))
	}
	for i, path := range paths {
		if logs[i].Path != path {
			t.Errorf("entry %d: expected path %q, got %q", i, path, logs[i].Path)
		}
	}
}



package main

import (
	"fmt"
)

func RetainFirstHalf(s string) string {
	if len(s) == 1  {
		return s
	}
	if s == "" {
		return ""
	}
	fmt.Println(len(s))
	half := len(s)/2
	return s[:half]
}


func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}





package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type pageData struct {
	Result string
	Text   string
}

type application struct {
	AppTemplate *template.Template
}

func main() {
	templateTemplate, err := template.ParseGlob("templates/*.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	appdata := application{
		AppTemplate: templateTemplate,
	}

	http.HandleFunc("/", appdata.HomeHandler)
	http.HandleFunc("/ascii-art", appdata.AsciiArtHandler)

	fmt.Println("Server running on port: 8080")
	http.ListenAndServe(":8080", nil)
}

func (app *application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := app.AppTemplate.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func (app *application) AsciiArtHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	bannerName := r.FormValue("banner")
	textValue := r.FormValue("text")

	if textValue == "" {
		http.Error(w, "Empty Field", http.StatusBadRequest)
		return
	}

	bannerDir := "banners/" + bannerName

	loadedbanner := LoadBanner(bannerDir)
	if loadedbanner == nil {
		http.Error(w, "File Not Found", http.StatusNotFound)
		return
	}
	renderValue := Render(textValue, loadedbanner)

	pageDataValues := pageData{
		Result: renderValue,
		Text:   textValue,
	}
	err := app.AppTemplate.ExecuteTemplate(w, "index.html", pageDataValues)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}









<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Ascii-Art-Web</title>
    <link rel="stylesheet" href="">
    <link rel="icon" type="image/png" href="">
</head>
<body>
    <div class="div1">
        <header>
            <h1>Ascii-Art-Web</h1>
        </header>
    </div>

    <div class="div2">
        <form action="/ascii-art" method="POST">
            <div class="div3">
                <label for="text">Enter the text. </label><br>
                 <textarea name="text" id="text" placeholder="Enter the text here"></textarea>
            </div>
               <div class="div4">
                 <select name="banner" id="banner">
                    <option>Select a banner.</option>
                    <option value="standard.txt">Standard</option>
                    <option value="shadow.txt">Shadow</option>
                    <option value="thinkertoy.txt">Thinkertoy</option>
                </select>
               </div>
               <button type="submit">Generate The Art.</button>
        </form>
    </div>
    <div class="div5">
        <pre>{{ .Result }}</pre>
    </div>
    
</body>
</html>









package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(s string) [][]string {
	fileName, err := os.ReadFile(s)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	output := [][]string{}

	content := string(fileName)
	if content == "" {
		fmt.Println("Empty fontFile")
	}

	blocks := strings.Split(content, "\n\n")

	for _, block := range blocks {
		if block == "" {
			fmt.Println("Missing character")
			return nil
		}
		rows := strings.Split(block, "\n")
		if rows == nil {
			fmt.Println("Bad File Format")
			return nil
		}
		output = append(output, rows)
	}
	return output

}








package main

import (
	"fmt"
	"strings"
)

func Render(s string, font [][]string) string {
	var result strings.Builder

	s = strings.ReplaceAll(s, `\n`, "\n")

	words := strings.Split(s, "\n")

	for _, char := range words {
		for row := 0; row < 8; row++ {
			for _, ch := range char {
				index := int(ch) - 32
				if index < 0 || index > 94 {
					fmt.Println("Unsupported character")
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
