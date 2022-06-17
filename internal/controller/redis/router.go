package redis

import (
	"keeper/internal/usecase"
	"keeper/pkg/redispubhandler"
	"log"

	"gopkg.in/redis.v5"
)

func NewRouter(r *redis.Client, t usecase.UserRepository) {
	err := redispubhandler.Handle(r, "keeper:user_update", New(t))
	if err != nil {
		log.Fatal(err)
	}
}
