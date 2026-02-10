package services

import (
	"kasir-api/model"
	"kasir-api/repositories"
	"time"
)

type ReportService struct {
	reps *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{reps: repo}
}

func (s *ReportService) GenerateReport(dateFrom, dateTo time.Time) (*model.ReportDetail, error) {
	return s.reps.FetchReport(dateFrom, dateTo)
}
