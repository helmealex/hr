package model

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Title              string `gorm:"not null" json:"title"`
	Level              string `gorm:"not null" json:"level"`
	LocalHiringManager string `gorm:"not null" json:"local_hiring_manager"` // TODO: create separate table for hiring managers
	DEOwner            string `gorm:"not null" json:"de_owner"`             // TODO: create separate table for DE owners
	Status             string `gorm:"not null" json:"status"`
	HeadCountStatus    string `gorm:"not null" json:"head_count_status"` // e.g., "new", "replacement"
	FocusRecruiter     string `gorm:"not null" json:"focus_recruiter"`   // TODO: create separate table for recruiters
	Candidates         []Candidate
	AdditionalComments string `gorm:"type:text" json:"additional_comments"`
}
