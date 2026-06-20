package main

import (
	"fmt"
	"io"
	"net/http"
)

func Exercise2() {
	http.HandleFunc("/echo", echoHandle)
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
}

func echoHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusNoContent)
	}
	defer r.Body.Close()
	if body == nil {
		http.Error(w, "body cannot be empty", http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, string(body))
}
