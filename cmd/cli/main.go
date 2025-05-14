package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	urlToBeShortened := flag.String("url", "", "Give the url to be shortened")
	flag.Parse()

	if *urlToBeShortened == "" {
		log.Fatal("Empty url not allowed\n\n>>Use '-url https://example.com'")
	}

	urlHash := makeRequest(*urlToBeShortened)

	fmt.Printf(">>Encoded Url Hash = http://localhost:8080/%v\n", urlHash)
}
