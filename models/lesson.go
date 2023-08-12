package models

import (
	"gorm.io/gorm"
)

type Lesson struct {
	gorm.Model
	Title      string  `json:"title" gorm:"unique;not null;default:null"`
	Order      int     `json:"order" gorm:"not null;default:null"`
	CourseID   uint    `json:"courseId" gorm:"not null;default:null"`
	Course     Course  `json:"course"`
	Summary    string  `json:"summary" gorm:"not null;default:null"`
	ContentIds []uint  `json:"mediaIds" gorm:"not null;default:null;type:text"`
	Content    []Media `json:"content" gorm:"type:text"`
	Quizzes    []Quiz  `json:"quizzes" gorm:"type:text"`
	// VideoURL string? maybe not, just parse medias for the .mp4 formats?
}
