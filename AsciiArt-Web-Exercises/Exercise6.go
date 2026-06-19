package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func Exercis6() {
// 	http.HandleFunc("/dashboard", dashboardHandle)
// 	fmt.Println("Live Server on port: 8080")
// 	http.ListenAndServe(":8080", nil)
// }

// func dashboardHandle(w http.ResponseWriter, r *http.Request) {
// 	headerValue := r.Header.Get("X-API-Key")
// 	if headerValue != "secret123" {
// 		http.Error(w, "Unauthorized", 401)
// 		return
// 	}
// 	fmt.Fprint(w, "Welcome")
// }
