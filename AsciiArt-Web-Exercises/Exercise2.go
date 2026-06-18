package main

import (
	"fmt"
	"net/http"
)

func Exercise2() {
	http.HandleFunc("/hello", HelloHandle)
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
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
