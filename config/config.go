package config

import (
	"github.com/fatih/color"
	"github.com/thehowl/conf"
)

type Config struct {
	DSN           string
	Host          string
	APIKey        string
	RedisAddr     string
	RedisPassword string
	Workers       int
}

var confFile = "./config/config.conf"

func NewConfig() (*Config, error) {
	c := Config{
		DSN:       "root@/ripple",
		Host:      "127.0.0.1",
		APIKey:    "masterKey",
		RedisAddr: "localhost:6379",
		Workers:   4,
	}
	err := conf.Load(&c, confFile)
	switch {
	case err == conf.ErrNoFile:
		color.Yellow("No config.conf was found. Creating it")
		err := conf.Export(&c, confFile)
		if err != nil {
			color.Red("Couldn't create config.conf: %v.", err)
			return nil, err
		} else {
			color.Green("config.conf has been created!")
		}
	case err != nil:
		color.Red("config.conf couldn't be loaded: %v.", err)
		return nil, err
	}
	return &c, nil
}
