package redis

import (
	userUseCase "keeper/internal/usecase/user"
	"keeper/pkg/redispubhandler"
	"log"

	"gopkg.in/redis.v5"
)

func NewRouter(r *redis.Client) {
	err := redispubhandler.Handle(r, "keeper:user_update", userUseCase.New())
	if err != nil {
		log.Fatal(err)
	}
}
