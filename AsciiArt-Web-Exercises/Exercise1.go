package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/ping", pingHandle)
	http.HandleFunc("/hello", HelloHandle)
	http.HandleFunc("/count", CountHandle)
	http.HandleFunc("/calculate", CalculateHandle)
	http.HandleFunc("/agent", agentHandle)
	http.HandleFunc("/dashboard", dashboardHandle)
	http.HandleFunc("/legacy", legacyHandle)
	http.HandleFunc("/v2", redirectHandle)
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
}

func pingHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprintf(w, "Hello, Guest!")
		} else {
			fmt.Fprintf(w, "Hello, %v!", name)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return

	}
}

func CountHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Send a POST request with text to count words")
	case "POST":
		sliceOfBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusBadRequest)
			return
		}
		convrtToStr := string(sliceOfBytes)
		fmt.Fprint(w, len(convrtToStr))

	}
}

func CalculateHandle(w http.ResponseWriter, r *http.Request) {
	num1 := r.URL.Query().Get("a")
	var1, err := strconv.Atoi(num1)
	if err != nil {
		http.Error(w, "error", 400)
		return
	}
	num2 := r.URL.Query().Get("b")
	var2, err := strconv.Atoi(num2)
	if err != nil {
		http.Error(w, "error", 400)
		return
	}

	op := r.URL.Query().Get("op")
	if op != "add" && op != "substract" && op != "multiply" {
		http.Error(w, "error", 400)
		return
	}

	switch op {
	case "add":
		fmt.Fprint(w, var1+var2)
		return

	case "substract":
		fmt.Fprint(w, var2-var1)
		return

	case "multiply":
		fmt.Fprint(w, var1*var2)
		return

	}
}


func agentHandle(w http.ResponseWriter, r *http.Request) {
	headerValue := r.Header.Get("User-Agent")
	if headerValue == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	} else {
		fmt.Fprint(w, headerValue)
	}

}



func dashboardHandle(w http.ResponseWriter, r *http.Request) {
	headerValue := r.Header.Get("X-API-Key")
	if headerValue != "secret123" {
		http.Error(w, "Unauthorized", 401)
		return
	}
	fmt.Fprint(w, "Welcome")
}

func legacyHandle(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/v2", 301)
}

func redirectHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to version 2")
}
