package usecase

import (
	"keeper/internal/entity"
)

type UserRepository interface {
	GetUsers(name string) ([]entity.User, error)
}
