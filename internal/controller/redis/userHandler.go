package redis

import (
	"keeper/internal/usecase"
	"keeper/pkg/redispubhandler"
	"log"
	"strconv"
)

type handler struct {
	t usecase.UserUseCase
}

// type result struct {
// 	UserID int `json:"user_id"`
// }

func NewUserHandler(t usecase.UserUseCase) *handler {
	return &handler{
		t: t,
	}
}

func (b *handler) Response(r *redispubhandler.Context) {
	if r.Error != nil {
		log.Fatal(r.Error)
	}
	userID, err := strconv.Atoi(r.Message)
	if err == nil && userID != 0 {
		err := b.t.UpdateUser(userID)
		if err != nil {
			log.Println(err)
		}
	}

}
