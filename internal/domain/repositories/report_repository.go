package repositories

import (
	"repo/internal/domain/entities"
)

type ReportRepository interface {
	GetAll() ([]*entities.Report, error)
	GetByID(id int) (*entities.Report, error)
	Create(report *entities.Report) error
	Update(report *entities.Report) error
	Delete(id int) error
}
