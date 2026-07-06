# Why not call it Decode?



json.NewDecoder(file).Decode(&artists)

- uses Decode.

While:

json.Unmarshal(body, &artists)

- uses Unmarshal.

The difference isn't that one converts JSON and the other doesn't—they both do.

The difference is where the JSON comes from.

Unmarshal works with a slice of bytes ([]byte), like the body you got from io.ReadAll.
Decode works directly from a stream (an io.Reader), such as a file or an HTTP response body.



// {"artists"
// :"",
// "locations":"https://groupietrackers.herokuapp.com/api/locations",
// "dates":"https://groupietrackers.herokuapp.com/api/dates",
// "relation":"https://groupietrackers.herokuapp.com/api/relation"}


artists := FetchArtists()

PrintArtists(artists)


artists := FetchArtists()

SearchArtists(artists)


artists := FetchArtists()

RenderTemplate(w, artists)


artists := FetchArtists()
fmt.Println(len(artists))