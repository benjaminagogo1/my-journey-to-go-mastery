package main

import (
    "fmt"
    "net/http"
)



func main()  {
    http.HandleFunc("/", Homehandle)
    http.HandleFunc("/ascii-art", AsciiArthandle)
    http.HandleFunc("/ascii-switch", switchhandle)
    fmt.Println("Live")
    http.ListenAndServe(":8080", nil)
}





// package main

// import (
//  ""
// )

// type Artists struct {
//  ID int
//  Image string
//  Name string
//  Member []string
//  CreationDate int
//  FirstAlbum string

// }

package main

import (
    // "os"
    // "strings"
    // "context"
    "html/template"
    "net/http"
    "os"
    "strings"
)

var tmpl = template.Must(template.ParseGlob("templates/index.html"))

type PageData struct {
    Banner string
    Text   string
    Result string
}

func Homehandle(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "Page Not Found", 404)
        return
    }

    tmpl.Execute(w, nil)
}

func AsciiArthandle(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/ascii-art" {
        http.Error(w, "Page Not Found", 404)
        return
    }
    if r.Method != "POST" {
        http.Error(w, "Method Not Allowed", 405)
        return
    }
    texts := r.FormValue("text")
    if texts == "" {
        http.Error(w, "Bad Request", 400)
        return
    }
    banners := r.FormValue("banner")
    if banners == "" {
        http.Error(w, "Bad Request", 400)
        return
    }
    bannerDir := "banners/" + banners + ".txt"

    loadbannerV := LoadaBanner(bannerDir)
    if loadbannerV == nil {
        http.Error(w, "File Not Found", 404)
        return
    }
    renderV := render(texts, loadbannerV)

    Data := PageData{
        Banner: banners,
        Text:   texts,
        Result: renderV,
    }

    tmpl.ExecuteTemplate(w, "index.html", Data)
}

func switchhandle(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/ascii-switch" {
        http.Error(w, "Page Not Found", 404)
        return
    }
    if r.Method != "GET" {
        http.Error(w, "Method Not Allowed", 405)
        return
    }
    texts := r.URL.Query().Get("text")
    if texts == "" {
        http.Error(w, "Bad Request", 400)
        return
    }
    banners := r.URL.Query().Get("banner")
    if banners == "" {
        http.Error(w, "Bad Request", 400)
        return
    }

    if banners != "standard" && banners != "shadow" && banners != "thinkertoy" {
        http.Error(w, "File Not Found", 404)
        return
    }
    bannerDir := "banners/" + banners + ".txt"

    loadbannerV := LoadaBanner(bannerDir)
    if loadbannerV == nil {
        http.Error(w, "File Not Found", 404)
        return
    }
    renderV := render(texts, loadbannerV)

    Data := PageData{
        Banner: banners,
        Text:   texts,
        Result: renderV,
    }
    tmpl.ExecuteTemplate(w, "index.html", Data)
}

func LoadaBanner(fileName string) [][]string {
    file, err := os.ReadFile(fileName)
    if err != nil {
        return nil
    }

    content := string(file)
    Output := [][]string{}

    blocks := strings.Split(content, "\n\n")

    for _, block := range blocks {
        rows := strings.Split(block, "\n")

        Output = append(Output, rows)
    }
    return Output
}

func render(s string, font [][]string) string {
    var result strings.Builder

    words := strings.Split(s, "\n")

    for _, char := range words {
        for rows := 0; rows < 8; rows++ {
            for _, ch := range char {
                index := int(ch) - 32
                if index < 0 || index > 94 {
                    continue
                }
                result.WriteString(font[index][rows])
            }
            result.WriteString("\n")

        }
    }
    return result.String()
}




joseph benjamin <josephbenjamin046@gmail.com>
	
11:18 PM (9 minutes ago)
	
	
to me
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>AsciiArt-Web</title>
</head>
<body>
    {{if not .Result}}
    <form action="/ascii-art" method="POST">
        <textarea name="text" id="text">{{.Text}}</textarea><br>

        <select name="banner">
            <option value="standard">standard</option>
            <option value="shadow">shadow</option>
            <option value="thinkertoy">thinkertoy</option>
        </select>
        <button type="submit">Generate</button>
    </form>
    {{end}}
    {{if .Result}}
    <pre>{{.Result}}</pre>
    <a href="ascii-switch?text={{.Text}}&banner=standard" class="active">standard</a>
    <a href="ascii-switch?text={{.Text}}&banner=shadow" class="active">shadow</a>
    <a href="ascii-switch?text={{.Text}}&banner=thinkertoy" class="active">thinkertoy</a><br><br>
    <a href="/">Back</a>
    {{end}}
</body>
</html>
 