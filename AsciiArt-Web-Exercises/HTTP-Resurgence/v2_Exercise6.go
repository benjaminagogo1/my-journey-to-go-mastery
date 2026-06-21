package main


import (
	"fmt"
	"net/http"
)

func Exercise6() {
	mainMux := http.NewServeMux()
	apiMux := http.NewServeMux()

	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	mainMux.HandleFunc("/api/v1/ping", pingHandle)
	mainMux.HandleFunc("/api/v1/greet", greetHandle)
	fmt.Println("Server is running.....")
	http.ListenAndServe(":8080", mainMux)
}

func pingHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func greetHandle(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Greetings, Stranger!")
	} else {
		fmt.Fprintf(w, "Greetings, %v!", name)
	}
}
