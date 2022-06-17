package app

import (
	"fmt"
	"keeper/config"
	redisHandler "keeper/internal/controller/redis"
	repo "keeper/internal/repository/user"
	"keeper/internal/usecase"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/meilisearch/meilisearch-go"
	"gopkg.in/redis.v5"
)

var wg *sync.WaitGroup

func Run(conf *config.Config) {
	wg = &sync.WaitGroup{}
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

	usersUseCase := usecase.New(
		repo.New(db),
	)

	redisHandler.NewRouter(r, usersUseCase)

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
	wg.Add(1)
	go func() {
		for {
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()

}
