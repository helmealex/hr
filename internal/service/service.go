package service

import (
	apiModel "hr/internal/api/model"
	"hr/internal/persistence"
	"log/slog"
)

type Service struct {
	logger *slog.Logger
	db     *persistence.DB
}

func New(
	logger *slog.Logger,
	db *persistence.DB,
) *Service {
	return &Service{
		logger: logger,
		db:     db,
	}
}

func (s *Service) GetJobs() ([]apiModel.JobResponse, error) {
	jobs, err := s.db.GetJobs(false)
	if err != nil {
		return nil, err
	}

	var jobResponses []apiModel.JobResponse
	for _, job := range jobs {
		jobResponses = append(jobResponses, apiModel.JobResponse{
			Title:              job.Title,
			Level:              job.Level,
			LocalHiringManager: job.LocalHiringManager,
			DEOwner:            job.DEOwner,
			Status:             job.Status,
			HeadCountStatus:    job.HeadCountStatus,
			FocusRecruiter:     job.FocusRecruiter,
			AdditionalComments: job.AdditionalComments,
		})
	}
	return jobResponses, nil
}
