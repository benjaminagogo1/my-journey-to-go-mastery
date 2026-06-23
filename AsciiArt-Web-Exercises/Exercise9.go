package main

import (
	"fmt"
	"net/http"
	"sync"
)

var count int
var mu sync.Mutex

func Exercise3() {
	http.HandleFunc("/counter", counterHandler)
	fmt.Println("Server is Live....")
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func counterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	mu.Lock()
	count++
	current := count
	mu.Unlock()

	if current > 1 {
		fmt.Fprintf(w, "You have visited %v times", current)
	} else {
		fmt.Fprintf(w, "You have visited %v time", current)
	}
}
