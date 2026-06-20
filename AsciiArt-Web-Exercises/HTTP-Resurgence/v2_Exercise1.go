package main

import (
	"net/http"
	"fmt"
)

func Exercise1()  {
	http.HandleFunc("/method-inspector", MethodHandle)
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
}

func MethodHandle(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "You made a %v request", r.Method)
}