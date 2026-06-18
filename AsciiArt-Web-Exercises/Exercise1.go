package main

import (
	"net/http"
	"fmt"
)

func Exercise1()  {
	http.HandleFunc("/ping", pingHandle)
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
}


func pingHandle(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "pong")
}