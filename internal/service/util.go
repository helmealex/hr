package service

import (
	apiModel "hr/internal/api/model"
	persistenceModel "hr/internal/persistence/model"
)

func jobApiToPersistence(job apiModel.Job) *persistenceModel.Job {
	return &persistenceModel.Job{
		Title:              job.Title,
		Level:              job.Level,
		LocalHiringManager: job.LocalHiringManager,
		DEOwner:            job.DEOwner,
		Status:             job.Status,
		HeadCountStatus:    job.HeadCountStatus,
		FocusRecruiter:     job.FocusRecruiter,
		AdditionalComments: job.AdditionalComments,
	}
}

func jobPersistenceToApiResponse(job *persistenceModel.Job) *apiModel.JobResponse {
	var candidates []apiModel.CandidateResponse
	for _, candidate := range job.Candidates {
		candidates = append(
			candidates,
			candidatePersistenceToApiResponse(&candidate),
		)
	}

	return &apiModel.JobResponse{
		ID:                 job.ID,
		Title:              job.Title,
		Level:              job.Level,
		LocalHiringManager: job.LocalHiringManager,
		DEOwner:            job.DEOwner,
		Status:             job.Status,
		HeadCountStatus:    job.HeadCountStatus,
		FocusRecruiter:     job.FocusRecruiter,
		AdditionalComments: job.AdditionalComments,
		Candidates:         candidates,
	}
}

func candidatePersistenceToApiResponse(candidate *persistenceModel.Candidate) apiModel.CandidateResponse {
	return apiModel.CandidateResponse{
		ID:                       candidate.ID,
		Name:                     candidate.Name,
		Email:                    candidate.Email,
		JobID:                    candidate.JobID,
		Phone:                    candidate.Phone,
		ApplicationDate:          candidate.ApplicationDate,
		Recruiter:                candidate.Recruiter,
		CurrentStatus:            candidate.CurrentStatus,
		Source:                   candidate.Source,
		Comments:                 candidate.Comments,
		HRInterviewStatus:        candidate.HRInterviewStatus,
		MinimumSalaryExpectation: candidate.MinimumSalaryExpectation,
		MaximumSalaryExpectation: candidate.MaximumSalaryExpectation,
		LocalTechInterviewStatus: candidate.LocalTechInterviewStatus,
		DEInterviewStatus:        candidate.DEInterviewStatus,
		OfferStatus:              candidate.OfferStatus,
		Hire:                     candidate.Hire,
		AdditionalComments:       candidate.AdditionalComments,
	}
}

func candidateApiToPersistence(candidate apiModel.Candidate) *persistenceModel.Candidate {
	return &persistenceModel.Candidate{
		Name:                     candidate.Name,
		Email:                    candidate.Email,
		Phone:                    candidate.Phone,
		JobID:                    candidate.JobID,
		ApplicationDate:          candidate.ApplicationDate,
		Recruiter:                candidate.Recruiter,
		CurrentStatus:            candidate.CurrentStatus,
		Source:                   candidate.Source,
		Comments:                 candidate.Comments,
		HRInterviewStatus:        candidate.HRInterviewStatus,
		MinimumSalaryExpectation: candidate.MinimumSalaryExpectation,
		MaximumSalaryExpectation: candidate.MaximumSalaryExpectation,
		LocalTechInterviewStatus: candidate.LocalTechInterviewStatus,
		DEInterviewStatus:        candidate.DEInterviewStatus,
		OfferStatus:              candidate.OfferStatus,
		Hire:                     candidate.Hire,
		AdditionalComments:       candidate.AdditionalComments,
	}
}
