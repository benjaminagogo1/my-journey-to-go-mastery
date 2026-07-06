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
	
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var artists []Artists

	err = json.Unmarshal(body, &artists)
	if err != nil {
		return
	}
	fmt.Println(artists)
}
