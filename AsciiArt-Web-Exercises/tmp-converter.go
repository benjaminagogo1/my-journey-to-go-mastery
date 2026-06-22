package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Exercise2() {
	http.HandleFunc("/convert", tempConvertHandler)
	fmt.Println("Server is Live....")
	http.ListenAndServe(":8080", nil)
}

func tempConvertHandler(w http.ResponseWriter, r *http.Request) {
	celsius := r.URL.Query().Get("celsius")
	if celsius == "" {
		http.Error(w, "empty input", http.StatusBadRequest)
		return
	}
	celsiusFloat, err := strconv.ParseFloat(celsius, 64)
	if err != nil {
		fmt.Println("Error converting the string")
		return
	}
	Fahrenheit := celsiusFloat*(9.0/5.0) + 32
	fmt.Fprintf(w, "Result: %.2f°F", Fahrenheit)
}
