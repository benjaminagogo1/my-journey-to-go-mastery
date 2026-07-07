package main

import (
	"fmt"
	"net/http"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	err := templ.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func AsciiArtHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	templateText := r.FormValue("text")
	if templateText == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	templateFile := r.FormValue("banner")
	if templateFile == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	bannerDir := "banners/" + templateFile
	fmt.Println(bannerDir)

	bannerValue, err := Loadbanner(bannerDir)
	fmt.Println(bannerValue)
	if err != nil {
		http.Error(w, "File Not Found", http.StatusNotFound)
		return
	}
	renderValue := render(templateText, bannerValue)

	err = templ.ExecuteTemplate(w, "result.html", renderValue)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
