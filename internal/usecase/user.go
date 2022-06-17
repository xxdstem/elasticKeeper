package usecase

import (
	"keeper/pkg/redispubhandler"
	"log"
)

type test struct {
	_interface struct {
		Message string `json:"message"`
	}
}

func New() *test {
	return &test{}
}

func (b *test) Response(r *redispubhandler.Context) {
	log.Println(r.Error)
	log.Println(r.Message)

}
func (b *test) GetInterface() interface{} {
	return &b._interface

}
