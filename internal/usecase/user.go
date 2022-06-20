package usecase

import "keeper/internal/entity"

type userUseCase struct {
	db    UserRepository
	meili UserRepository
}

func New(db UserRepository, meili UserRepository) UserRepository {
	return &userUseCase{db: db, meili: meili}
}

func (u *userUseCase) GetUsers(name string) ([]entity.User, error) {
	return u.db.GetUsers(name)
}

func (u *userUseCase) GetUserByID(id int) (*entity.User, error) {
	return u.meili.GetUserByID(id)
}
