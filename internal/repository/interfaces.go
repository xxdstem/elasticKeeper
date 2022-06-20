package repository

import (
	"keeper/internal/entity"
)

type UserRepository interface {
	GetUsers(name string) ([]entity.User, error)
	GetUserByID(id int) (*entity.User, error)
}

type BeatmapRepository interface {
	GetBeatmapByID(id int) (*entity.BeatmapSet, error)
}
