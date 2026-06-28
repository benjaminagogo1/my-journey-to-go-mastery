package main

import (
	"net/http"
)

func FetchArtists() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return
	}
	defer resp.Body.Close()
}