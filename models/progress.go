package models

import "gorm.io/gorm"

type Progress struct {
	gorm.Model
	EnrollmentID uint    `json:"enrollmentId" gorm:"unique;not null;default:null"`
	Completed    bool    `json:"completed" gorm:"not null;default:false"`
	Progress     float64 `json:"progress" gorm:"not null;default:0"`
}
