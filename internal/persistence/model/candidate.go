package model

import (
	"time"

	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	Name                     string `gorm:"not null"`
	Email                    string `gorm:"not null;unique"`
	Phone                    string
	JobID                    uint
	ApplicationDate          *time.Time `gorm:"not null"`
	Recruiter                string     `gorm:"not null"`  //TODO: create separaste table for recruiters
	CurrentStatus            string     `gorm:"not null"`  // e.g., "linkedin_im", "pending_linkedin_connection", "to_give_feedback", "feedback_given"
	Source                   string     `gorm:"not null"`  // e.g., "applicant", "referral", "job_board","agency", "cold_call"
	Comments                 string     `gorm:"type:text"` // Additional comments about the candidate
	HRInterviewStatus        string     `gorm:"not null"`
	MinimumSalaryExpectation float64    `gorm:"not null"`
	MaximumSalaryExpectation float64    `gorm:"not null"`
	LocalTechInterviewStatus string     `gorm:"not null"`
	DEInterviewStatus        string     `gorm:"not null"`
	OfferStatus              string     `gorm:"not null"`
	Hire                     *bool
	AdditionalComments       string `gorm:"type:text"`
}
