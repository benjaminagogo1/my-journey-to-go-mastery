package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Artists struct {
	ID           int
	Name         string
	Image        string
	Member       []string
	CreationDate int
	FirstAlbum   string
}

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists/1"
	
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		
	}
}
