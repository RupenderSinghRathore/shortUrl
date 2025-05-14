package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/RupenderSinghRathore/shortUrl/internal/models"
)

func makeRequest(urlToBeShortned string) string {
	urlJsonReader := jsonWrappedReader(urlToBeShortned)
	url := "http://localhost:8080/hash"

	res, err := http.Post(url, "application/json", urlJsonReader)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	urlByte, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(urlByte)
}

func jsonWrappedReader(urlToBeShortned string) io.Reader {
	urlObj := models.UrlStr{Url: urlToBeShortned}

	jsonData, err := json.Marshal(urlObj)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.NewReader(jsonData)
}
