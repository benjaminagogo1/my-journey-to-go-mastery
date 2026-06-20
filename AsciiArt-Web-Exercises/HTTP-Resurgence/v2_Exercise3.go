package main

import (
	"net/http"
	"fmt"
)

func Exercise3()  {
	http.HandleFunc("/headers", HeadersHandle)
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
}

func HeadersHandle(w http.ResponseWriter, r *http.Request)  {
	HeaderToken := r.Header.Get("X-Custom-Token")
	if HeaderToken == "" {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
		return
	}
	if HeaderToken != "" {
		fmt.Fprint(w, "Token received: abc123")
	}
}