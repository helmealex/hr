package model

type Report struct {
	TotalVacancies          []VacancyReport   `json:"total_vacancies"`
	TotalApplicants         []ApplicantReport `json:"total_applicants"`
	TotalExtendedOffers     int               `json:"total_extended_offers"`
	TotalAcceptedOffers     int               `json:"total_accepted_offers"`
	TotalRejectedOffers     int               `json:"total_rejected_offers"`
	TimeToHireAverageDays   float64           `json:"time_to_hire_average_days"`
	TotalWithdrawals        int               `json:"total_withdrawals"`
	TotalRejectedCandidates int               `json:"total_rejected_candidates"`
	TimeToFillReq           float64           `json:"time_to_fill_req"`
}

type VacancyReport struct {
	Level string `json:"level"`
	Count int    `json:"count"`
}

type ApplicantReport struct {
	JobID uint   `json:"job_id"`
	Count int    `json:"count"`
	Title string `json:"title"`
}
