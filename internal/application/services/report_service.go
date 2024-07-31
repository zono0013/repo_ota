package services

import (
	"repo/internal/application/dtos"
	"repo/internal/domain/repositories"
)

type ReportService struct {
	reportRepo repositories.ReportRepository
}

func NewReportService(reportRepo repositories.ReportRepository) *ReportService {
	return &ReportService{reportRepo: reportRepo}
}

func (s *ReportService) GetAllReports() ([]*dtos.ReportDTO, error) {
	reports, err := s.reportRepo.GetAll()
	if err != nil {
		return nil, err
	}

	reportDTOs := make([]*dtos.ReportDTO, len(reports))
	for i, report := range reports {
		reportDTOs[i] = dtos.ReportToDTO(report)
	}
	return reportDTOs, nil
}

func (s *ReportService) GetReportByID(id int) (*dtos.ReportDTO, error) {
	report, err := s.reportRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return dtos.ReportToDTO(report), nil
}

func (s *ReportService) CreateReport(reportDTO *dtos.ReportDTO) (*dtos.ReportDTO, error) {
	report := reportDTO.ToEntity()
	err := s.reportRepo.Create(report)
	if err != nil {
		return nil, err
	}
	return dtos.ReportToDTO(report), nil
}

func (s *ReportService) UpdateReport(reportDTO *dtos.ReportDTO) (*dtos.ReportDTO, error) {
	report := reportDTO.ToEntity()
	err := s.reportRepo.Update(report)
	if err != nil {
		return nil, err
	}
	return dtos.ReportToDTO(report), nil
}

func (s *ReportService) DeleteReport(id int) error {
	return s.reportRepo.Delete(id)
}
