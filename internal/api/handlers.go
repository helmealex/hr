package api

import (
	"encoding/json"
	"fmt"
	"hr/internal/api/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) getJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := s.service.GetJobs()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get jobs: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(jobs)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal jobs: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResp)
}

func (s *Server) createJob(w http.ResponseWriter, r *http.Request) {
	var job model.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode job: %v", err), http.StatusBadRequest)
		return
	}

	if err := s.service.CreateJob(job); err != nil {
		http.Error(w, fmt.Sprintf("failed to create job: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) getJobByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "job ID is required", http.StatusBadRequest)
		return
	}

	job, err := s.service.GetJobByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get job by ID: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(job)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal job: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResp)
}

func (s *Server) getCandidates(w http.ResponseWriter, r *http.Request) {
	candidates, err := s.service.GetCandidates()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get candidates: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(candidates)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal candidates: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResp)
}

func (s *Server) createCandidate(w http.ResponseWriter, r *http.Request) {
	var candidate model.Candidate
	if err := json.NewDecoder(r.Body).Decode(&candidate); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode candidate: %v", err), http.StatusBadRequest)
		return
	}

	if err := s.service.CreateCandidate(candidate); err != nil {
		http.Error(w, fmt.Sprintf("failed to create candidate: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) getCandidateByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "candidate ID is required", http.StatusBadRequest)
		return
	}

	candidate, err := s.service.GetCandidateByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get candidate by ID: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(candidate)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal candidate: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResp)
}

func (s *Server) getReport(w http.ResponseWriter, r *http.Request) {
	report, err := s.service.GetReport()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to generate report: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(report)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal report: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(jsonResp)
}
