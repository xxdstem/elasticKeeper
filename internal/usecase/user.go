package usecase

import "keeper/internal/entity"

type userUseCase struct {
	repo UserRepository
}

func New(repo UserRepository) UserRepository {
	return &userUseCase{repo: repo}
}

func (u *userUseCase) GetUsers(name string) ([]entity.User, error) {
	return u.repo.GetUsers(name)
}
