package main

import (
	"fmt"
	"net/http"
)

func Exercise4() {
	http.HandleFunc("/form", formHandle)
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
}

func formHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	userName := r.FormValue("name")
	if userName == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}
	language := r.FormValue("Go")
	if language == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}
	if userName != "" && language != "" {
		fmt.Fprintf(w, "Hello %s, you are coding in %s", userName, language)
	}

}
