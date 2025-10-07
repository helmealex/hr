package model

type JobResponse struct {
	ID                 uint                `json:"id"`
	Title              string              `json:"title"`
	Level              string              `json:"level"`
	LocalHiringManager string              `json:"local_hiring_manager"`
	DEOwner            string              `json:"de_owner"`
	Status             string              `json:"status"`
	HeadCountStatus    string              `json:"head_count_status"`
	FocusRecruiter     string              `json:"focus_recruiter"`
	AdditionalComments string              `json:"additional_comments"`
	Candidates         []CandidateResponse `json:"candidates,omitempty"` // Optional field to include candidates
}

type Job struct {
	Title              string `json:"title"`
	Level              string `json:"level"`
	LocalHiringManager string `json:"local_hiring_manager"`
	DEOwner            string `json:"de_owner"`
	Status             string `json:"status"`
	HeadCountStatus    string `json:"head_count_status"`
	FocusRecruiter     string `json:"focus_recruiter"`
	AdditionalComments string `json:"additional_comments"`
}

/*
{
  "title": "Software Engineer",
  "level": "Senior",
  "local_hiring_manager": "Jane Doe",
  "de_owner": "John Smith",
  "status": "open",
  "head_count_status": "Approved",
  "focus_recruiter": "Alice Johnson",
  "additional_comments": "Looking for a candidate with strong Go experience."
}
*/
