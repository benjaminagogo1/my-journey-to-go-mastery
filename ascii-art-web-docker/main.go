package main

import (
	"net/http"
	"html/template"
	"fmt"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func main()  {

	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", HomeHandle)
	http.HandleFunc("/ascii-art", AsciiArtHandle)
	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}
