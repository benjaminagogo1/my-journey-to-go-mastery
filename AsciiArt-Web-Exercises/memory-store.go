package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var result string
var mux sync.Mutex

func Exercise() {
	http.HandleFunc("/store", storeHandler)
	fmt.Println("Server is Live...")
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func storeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mux.Lock()
		value := result
		mux.Unlock()
		if value == "" {
			fmt.Fprint(w, "Nothing stored yet")
			return
		}
		fmt.Fprint(w, value)

	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusNotFound)
			return
		}
		defer r.Body.Close()

		mux.Lock()
		result = string(body)
		mux.Unlock()

		fmt.Fprint(w, "Saved")
	}
}
