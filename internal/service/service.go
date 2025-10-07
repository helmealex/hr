package service

import (
	apiModel "hr/internal/api/model"
	"hr/internal/persistence"
	"hr/internal/persistence/model"
	"log/slog"
	"strconv"
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

// Jobs
func (s *Service) GetJobs() ([]apiModel.JobResponse, error) {
	jobs, err := s.db.GetJobs(false)
	if err != nil {
		return nil, err
	}

	var jobResponses []apiModel.JobResponse
	for _, job := range jobs {
		jobResponses = append(jobResponses, *jobPersistenceToApiResponse(&job))
	}
	return jobResponses, nil
}

func (s *Service) CreateJob(job apiModel.Job) error {
	newJob := jobApiToPersistence(job)

	if err := s.db.CreateJob(newJob); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetJobByID(id string) (*apiModel.JobResponse, error) {
	uintId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	job, err := s.db.GetJobByID(uint(uintId))
	if err != nil {
		return nil, err
	}

	return jobPersistenceToApiResponse(job), nil
}

// Candidates
func (s *Service) GetCandidates() ([]apiModel.CandidateResponse, error) {
	candidates, err := s.db.GetCandidates()
	if err != nil {
		return nil, err
	}

	var candidateResponses []apiModel.CandidateResponse
	for _, candidate := range candidates {
		candidateResponses = append(candidateResponses, candidatePersistenceToApiResponse(&candidate))
	}
	return candidateResponses, nil
}

func (s *Service) CreateCandidate(candidate apiModel.Candidate) error {
	newCandidate := candidateApiToPersistence(candidate)

	if err := s.db.CreateCandidate(newCandidate); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetCandidateByID(id string) (*apiModel.CandidateResponse, error) {
	uintId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	candidate, err := s.db.GetCandidateByID(uint(uintId))
	if err != nil {
		return nil, err
	}

	candidateResponse := candidatePersistenceToApiResponse(candidate)

	return &candidateResponse, nil
}

func (s *Service) GetReport() (*apiModel.ReportResponse, error) {
	vacancies, err := s.db.GetVacancies()
	if err != nil {
		return nil, err
	}

	applicants, err := s.db.GetApplicantsTotal()
	if err != nil {
		return nil, err
	}

	persistenceReport := model.Report{
		TotalVacancies:  vacancies,
		TotalApplicants: applicants,
	}

	reportResponse := reportPersistenceToApiResponse(persistenceReport)

	return reportResponse, nil
}
