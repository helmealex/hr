package model

type ReportResponse struct {
	TotalVacancies          []VacancyReportResponse   `json:"total_vacancies"`
	TotalApplicants         []ApplicantReportResponse `json:"total_applicants"`
	TotalExtendedOffers     int                       `json:"total_extended_offers"`
	TotalAcceptedOffers     int                       `json:"total_accepted_offers"`
	TotalRejectedOffers     int                       `json:"total_rejected_offers"`
	TimeToHireAverageDays   float64                   `json:"time_to_hire_average_days"`
	TotalWithdrawals        int                       `json:"total_withdrawals"`
	TotalRejectedCandidates int                       `json:"total_rejected_candidates"`
	TimeToFillReq           float64                   `json:"time_to_fill_req"`
}

type VacancyReportResponse struct {
	Level string `json:"level"`
	Count int    `json:"count"`
}

type ApplicantReportResponse struct {
	JobID    uint   `json:"job_id"`
	Count    int    `json:"count"`
	JobTitle string `json:"title"`
}
