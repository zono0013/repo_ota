package repositories

import (
	"repo/internal/domain/entities"
)

type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id int) (*entities.User, error)
	Update(user *entities.User) error
}
