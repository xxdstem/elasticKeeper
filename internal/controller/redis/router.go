package redis

import (
	"keeper/internal/usecase"
	"keeper/pkg/redispubhandler"
	"log"

	"gopkg.in/redis.v5"
)

func NewRouter(r *redis.Client, t usecase.UserUseCase) {
	userHandler := NewUserHandler(t)
	err := redispubhandler.Handle(r, "keeper:user_update", userHandler)
	if err != nil {
		log.Fatal(err)
	}
	err = redispubhandler.Handle(r, "peppy:ban", userHandler)
	if err != nil {
		log.Fatal(err)
	}
}
