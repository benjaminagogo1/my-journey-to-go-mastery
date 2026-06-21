package main


import (
	"fmt"
	"html/template"
	"net/http"
)

func Exercise7() {
	http.HandleFunc("/render", renderHandle)
	fmt.Println("Server running....")
	http.ListenAndServe(":8080", nil)
}

const tmplStr = `
<!DOCTYPE html>
<html>
<head><title>{{.Title}}</title></head>
<body>
	<h1>{{.Title}}</h1>
	<p>{{.Body}}</p>
</body>
`

type pageData struct {
	Title string
	Body  string
}

func renderHandle(w http.ResponseWriter, r *http.Request) {

	tpl := template.Must(template.New("page").Parse(tmplStr))

	title := r.URL.Query().Get("title")
	body := r.URL.Query().Get("body")
	if title == "" || body == "" {
		http.Error(w, "title and body are required", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	 
	err := tpl.Execute(w, pageData{Title: title, Body: body})
	if err != nil {
		http.Error(w, "template execution failsed", http.StatusInternalServerError)
		return
	}
}
