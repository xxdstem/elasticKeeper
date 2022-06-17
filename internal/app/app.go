package app

import (
	"fmt"
	"keeper/config"
	"log"

	redisHandler "keeper/internal/controller/redis"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/meilisearch/meilisearch-go"
	"gopkg.in/redis.v5"
)

func Run(conf *config.Config) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   fmt.Sprintf("http://%s", conf.Host),
		APIKey: conf.APIKey,
	})
	db, err := sqlx.Open("mysql", conf.DSN)
	if err != nil {
		log.Fatalf("couldn't start MySQL connection: %v.", err)
		return
	}
	defer db.Close()
	r := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPassword,
	})
	redisHandler.NewRouter(r)

	// An index is where the documents are stored.
	index := client.Index("beatmaps_full")

	_, err = index.Search("Alumetri", &meilisearch.SearchRequest{
		Limit: 10,
		Sort: []string{
			"ranking_data:desc",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

}
