package dtos

import "repo/internal/domain/entities"

type UserDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required,min=2,max=100"`
}

// エンティティからDTOへの変換
func UserToDTO(user *entities.User) *UserDTO {
	return &UserDTO{
		ID:   user.ID,
		Name: user.Name,
	}
}

// DTOからエンティティへの変換
func (dto *UserDTO) ToEntity() *entities.User {
	return &entities.User{
		ID:   dto.ID,
		Name: dto.Name,
	}
}
