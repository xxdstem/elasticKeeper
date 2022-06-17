package app

import (
	"fmt"
	"keeper/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/meilisearch/meilisearch-go"
	"gopkg.in/redis.v5"
)

func Run(conf *config.Config) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   fmt.Sprintf("http://%s:7700", conf.Host),
		APIKey: conf.APIKey,
	})

	db, err := sqlx.Open("mysql", conf.DSN)
	if err != nil {
		log.Fatalf("couldn't start MySQL connection: %v.", err)
		return
	}
	defer db.Close()
	_ = redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPassword,
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
