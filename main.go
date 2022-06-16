package main

import (
	"log"

	"github.com/meilisearch/meilisearch-go"
)

func main() {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://95.216.193.66:7700",
		APIKey: "GATARI",
	})
	// An index is where the documents are stored.
	index := client.Index("beatmaps_full")

	resp, err := index.Search("Alumetri", &meilisearch.SearchRequest{
		Limit: 10,
		Sort: []string{
			"ranking_data:desc",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
}
