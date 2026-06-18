package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandle)
	http.HandleFunc("/about", AsciiArtHandle)
	// http.HandleFunc("/about-me", AboutMeHandle)
	fmt.Println("Server Live on port: 8080")
	http.ListenAndServe(":8080", nil)
}

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		if r.URL.Path != "/" {
			http.Error(w, "Page Not Found", http.StatusNotFound)
			return
		}
		tpl, err := template.ParseGlob("templates/*.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = tpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func AsciiArtHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	tpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	templateText := r.FormValue("text")
	if templateText == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	templateFile := r.FormValue("banner")
	if templateFile == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	bannerDir := "banners/" + templateFile

	loadebannerValue := LoadBanner(bannerDir)

	if loadebannerValue == nil {
		http.Error(w, "File Not Found", http.StatusNotFound)
		return
	}
	renderValue := Render(templateText, loadebannerValue)

	err = tpl.ExecuteTemplate(w, "index.html", renderValue)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

// func AboutMeHandle(w http.ResponseWriter, r *http.Request) {

// }
