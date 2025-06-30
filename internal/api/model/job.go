package model

type JobResponse struct {
	Title              string `json:"title"`
	Level              string `json:"level"`
	LocalHiringManager string `json:"local_hiring_manager"`
	DEOwner            string `json:"de_owner"`
	Status             string `json:"status"`
	HeadCountStatus    string `json:"head_count_status"`
	FocusRecruiter     string `json:"focus_recruiter"`
	AdditionalComments string `json:"additional_comments"`
}
