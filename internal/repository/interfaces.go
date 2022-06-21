package repository

import (
	"keeper/internal/entity"
)

type UserRepository interface {
	GetUsers(name string) ([]entity.User, error)
	GetUserByID(id int) (*entity.User, error)
}

type UserMeiliRepository interface {
	UpdateUser(*entity.User) error
}

type BeatmapMeiliRepository interface {
	UpdateBeatmap(*entity.BeatmapSet) error
}
type BeatmapRepository interface {
	GetBeatmapByID(id int) (*entity.BeatmapSet, error)
}
