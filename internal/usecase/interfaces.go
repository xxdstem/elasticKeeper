package usecase

import "keeper/internal/entity"

type _useCase struct {
	db    UserRepository
	meili UserMeiliRepository
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
