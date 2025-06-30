package model

import (
	"time"

	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	Name                     string `gorm:"not null" json:"name"`
	Email                    string `gorm:"not null;unique" json:"email"`
	Phone                    string
	JobID                    uint
	ApplicationDate          *time.Time `gorm:"not null" json:"application_date"`
	Recruiter                string     `gorm:"not null" json:"recruiter"`      //TODO: create separaste table for recruiters
	CurrentStatus            string     `gorm:"not null" json:"current_status"` // e.g., "linkedin_im", "pending_linkedin_connection", "to_give_feedback", "feedback_given"
	Source                   string     `gorm:"not null" json:"source"`         // e.g., "applicant", "referral", "job_board","agency", "cold_call"
	Comments                 string     `gorm:"type:text" json:"comments"`      // Additional comments about the candidate
	HRInterviewStatus        string     `gorm:"not null" json:"hr_interview_status"`
	MinimumSalaryExpectation float64    `gorm:"not null" json:"minimum_salary"`
	MaximumSalaryExpectation float64    `gorm:"not null" json:"maximum_salary"`
	LocalTechInterviewStatus string     `gorm:"not null" json:"local_tech_interview_status"`
	DEInterviewStatus        string     `gorm:"not null" json:"de_interview_status"`
	OfferStatus              string     `gorm:"not null" json:"offer_status"`
	Hire                     bool       `gorm:"not null" json:"hire"`
	AdditionalComments       string     `gorm:"type:text" json:"additional_comments"`
}
