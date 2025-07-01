package model

import "time"

type CandidateResponse struct {
	ID                       uint       `json:"id"`
	Name                     string     `json:"name"`
	Email                    string     `json:"email"`
	Phone                    string     `json:"phone"`
	ApplicationDate          *time.Time `json:"application_date"`
	Recruiter                string     `json:"recruiter"`
	CurrentStatus            string     `json:"current_status"`
	Source                   string     `json:"source"`
	Comments                 string     `json:"comments"`
	HRInterviewStatus        string     `json:"hr_interview_status"`
	MinimumSalaryExpectation float64    `json:"minimum_salary"`
	MaximumSalaryExpectation float64    `json:"maximum_salary"`
	LocalTechInterviewStatus string     `json:"local_tech_interview_status"`
	DEInterviewStatus        string     `json:"de_interview_status"`
	OfferStatus              string     `json:"offer_status"`
	Hire                     *bool      `json:"hire"`
	AdditionalComments       string     `json:"additional_comments"`
}

type Candidate struct {
	Name                     string     `json:"name"`
	Email                    string     `json:"email"`
	Phone                    string     `json:"phone"`
	JobID                    uint       `json:"job_id"`
	ApplicationDate          *time.Time `json:"application_date"`
	Recruiter                string     `json:"recruiter"`
	CurrentStatus            string     `json:"current_status"`
	Source                   string     `json:"source"`
	Comments                 string     `json:"comments"`
	HRInterviewStatus        string     `json:"hr_interview_status"`
	MinimumSalaryExpectation float64    `json:"minimum_salary"`
	MaximumSalaryExpectation float64    `json:"maximum_salary"`
	LocalTechInterviewStatus string     `json:"local_tech_interview_status"`
	DEInterviewStatus        string     `json:"de_interview_status"`
	OfferStatus              string     `json:"offer_status"`
	Hire                     *bool      `json:"hire"`
	AdditionalComments       string     `json:"additional_comments"`
}
