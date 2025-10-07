package model

import "time"

type CandidateResponse struct {
	ID                       uint       `json:"id"`
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

// Example JSON payload:
// {
//   "id": 1,
//   "name": "John Doe",
//   "email": "john.doe@example.com",
//   "phone": "123-456-7890",
//   "job_id": 101,
//   "application_date": "2023-10-27T10:00:00Z",
//   "recruiter": "Jane Smith",
//   "current_status": "In Progress",
//   "source": "LinkedIn",
//   "comments": "Strong candidate with relevant experience.",
//   "hr_interview_status": "Passed",
//   "minimum_salary": 80000,
//   "maximum_salary": 95000,
//   "local_tech_interview_status": "Scheduled",
//   "de_interview_status": "Pending",
//   "offer_status": "Not Yet Offered",
//   "hire": null,
//   "additional_comments": "Follow up next week."
// }

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
