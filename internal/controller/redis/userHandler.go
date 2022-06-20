package redis

import (
	"encoding/json"
	"keeper/internal/usecase"
	"keeper/pkg/redispubhandler"
	"log"
)

type handler struct {
	t usecase.UserRepository
}

type result struct {
	User   string `json:"user"`
	UserID int    `json:"user_id"`
}

func NewUserHandler(t usecase.UserRepository) *handler {
	return &handler{
		t: t,
	}
}

func (b *handler) Response(r *redispubhandler.Context) {
	if r.Error != nil {
		log.Fatal(r.Error)
	}

	var res = result{}
	if err := json.Unmarshal([]byte(r.Message), &res); err != nil {
		log.Println("ERROR on unmarshal")
		log.Println(err)
		return
	}
	if res.UserID != 0 {
		u, err := b.t.GetUserByID(res.UserID)
		log.Println(u)
		if err != nil {
			log.Println(err)
		}
	} else {
		users, err := b.t.GetUsers(res.User)
		log.Println(users)
		if err != nil {
			log.Println(err)
		}
	}

}
