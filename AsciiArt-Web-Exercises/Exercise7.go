package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/legacy", legacyHandle)
// 	http.HandleFunc("/v2", redirectHandle)
// 	fmt.Println("Server Running on port: 8080")
// 	http.ListenAndServe(":8080", nil)
// }

// func legacyHandle(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "/v2", 301)

// }

// func redirectHandle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Welcome to version 2")
// }
