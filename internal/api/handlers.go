package api

import (
	"encoding/json"
	"fmt"
	"net/http"
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
