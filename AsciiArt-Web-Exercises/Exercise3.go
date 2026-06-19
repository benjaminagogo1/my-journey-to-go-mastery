package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func Exercise3() {
// 	http.HandleFunc("/count", CountHandle)
// 	fmt.Println("Live")
// 	http.ListenAndServe(":8080", nil)
// }

// // func CountHandle(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "GET":
// 		fmt.Fprintf(w, "Send a POST request with text to count words")
// 	case "POST":
// 		sliceOfBytes, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			http.Error(w, "failed to read body", http.StatusBadRequest)
// 			return
// 		}
// 		convrtToStr := string(sliceOfBytes)
// 		fmt.Fprint(w, len(convrtToStr))

// 	}
// }
