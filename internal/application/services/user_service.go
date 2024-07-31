package services

import (
	"repo/internal/application/dtos"
	"repo/internal/domain/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(userDTO *dtos.UserDTO) (*dtos.UserDTO, error) {
	user := userDTO.ToEntity()
	err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return dtos.UserToDTO(user), nil
}

func (s *UserService) GetUser(id int) (*dtos.UserDTO, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return dtos.UserToDTO(user), nil
}

func (s *UserService) UpdateUser(userDTO *dtos.UserDTO) (*dtos.UserDTO, error) {
	user := userDTO.ToEntity()
	err := s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}
	return dtos.UserToDTO(user), nil
}
