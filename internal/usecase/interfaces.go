package usecase

import "keeper/internal/entity"

type _userUseCase struct {
	db    UserRepository
	meili UserMeiliRepository
}
type _beatmapUseCase struct {
	db    BeatmapRepository
	meili BeatmapMeiliRepository
}

type BeatmapsUseCase interface {
	UpdateBeatmapSet(id int) error
}
type UserUseCase interface {
	UpdateUser(id int) error
}

type UserRepository interface {
	GetUsers(name string) ([]entity.User, error)
	GetUserByID(id int) (*entity.User, error)
}

type UserMeiliRepository interface {
	UpdateUser(*entity.User) error
}

type BeatmapRepository interface {
	GetBeatmapByID(id int) (*entity.BeatmapSet, error)
}

type BeatmapMeiliRepository interface {
	UpdateBeatmap(*entity.BeatmapSet) error
}
