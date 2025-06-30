package model

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	FirstName string `gorm:"not null" json:"first_name"`
}
