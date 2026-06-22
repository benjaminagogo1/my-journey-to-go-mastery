package main

import (
	"fmt"
	"net/http"
	"time"
)

var startTime time.Time

func Exercise2() {
	startTime = time.Now()
	http.HandleFunc("/health", HealthHandler)
	fmt.Println("Server is Live....")
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	uptime := time.Since(startTime).Seconds()
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, `{"status": "healthy","uptime": "%.0fs"}`+"\n", uptime)

}
