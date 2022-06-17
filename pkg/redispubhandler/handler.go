package redispubhandler

import (
	"encoding/json"

	"gopkg.in/redis.v5"
)

type Context struct {
	Error   error
	Message interface{}
}

type Request interface {
	GetInterface() interface{}
	Response(*Context)
}

func Handle(r *redis.Client, sub string, req Request) error {
	subscriber, err := r.Subscribe(sub)
	if err != nil {
		return err
	}
	go func() {
		for {
			msg, err := subscriber.ReceiveMessage()
			if err != nil {
				req.Response(&Context{
					Error: err,
				})
			}
			in := req.GetInterface()
			if err := json.Unmarshal([]byte(msg.Payload), &in); err != nil {
				req.Response(&Context{
					Error: err,
				})
			}
			req.Response(&Context{Message: in})
		}
	}()
	return nil
}
