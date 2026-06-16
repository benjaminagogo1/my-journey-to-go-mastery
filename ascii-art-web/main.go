package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ok", okHandle)
	http.HandleFunc("/notfound", notFoundHandle)
	http.HandleFunc("/badrequest", badRequestHandle)
	http.HandleFunc("/error", errrorHandle)
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
}

func okHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ok" {
		http.Error(w, "successful", 200)
	}

}

func notFoundHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/notfound" {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func badRequestHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/badrequest" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}

func errrorHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/error" {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// curl -i http://localhost:8080/ok 
// curl -i http://localhost:8080/notfound
// curl -i http://localhost:8080/badrequest
// curl -i http://localhost:8080/error