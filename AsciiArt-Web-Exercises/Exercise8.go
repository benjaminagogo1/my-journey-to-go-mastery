package main

import (
	"fmt"
	"log"
	"net/http"
)

func Exercise8() {
	http.Handle("/log", loginMiddleware(http.HandlerFunc(logHandle)))
	fmt.Println("Server Is Live...")
	http.ListenAndServe(":8080", nil)
}

func logHandle(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "saved")
	
}

func loginMiddleware(text http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s, %s", r.Method, r.URL.Path)

		text.ServeHTTP(w, r)
		
	})
}
