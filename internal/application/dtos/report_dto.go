package dtos

import (
	"repo/internal/domain/entities"
)

type ReportDTO struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id" validate:"required"`
	Title        string `json:"title" validate:"required,min=5,max=200"`
	WordCount    int    `json:"word_count" validate:"required,min=1"`
	WritingStyle string `json:"writing_style" validate:"required,oneof=formal informal"`
	Language     string `json:"language" validate:"required,oneof=japanese english"`
}

// エンティティからDTOへの変換
func ReportToDTO(report *entities.Report) *ReportDTO {
	return &ReportDTO{
		ID:           report.ID,
		UserID:       report.UserID,
		Title:        report.Title,
		WordCount:    report.WordCount,
		WritingStyle: report.WritingStyle,
		Language:     report.Language,
	}
}

// DTOからエンティティへの変換
func (dto *ReportDTO) ToEntity() *entities.Report {
	return &entities.Report{
		ID:           dto.ID,
		UserID:       dto.UserID,
		Title:        dto.Title,
		WordCount:    dto.WordCount,
		WritingStyle: dto.WritingStyle,
		Language:     dto.Language,
	}
}
