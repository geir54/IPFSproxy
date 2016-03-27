package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Link struct {
	Hash       string
	ChangeDate time.Time
}

func getFromTracker(url string) (Link, bool) {
	res, err := http.Get("http://" + tracker + "?link=" + url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		// fmt.Println("Not found")
		return Link{}, false
	}

	data := &Link{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(data)
	if err != nil {
		log.Fatal(err)
	}

	return *data, true
}
