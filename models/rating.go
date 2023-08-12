package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	UserID   uuid.UUID `json:"userId" gorm:"not null;default:null"`
	User     User      `json:"user"`
	CourseID uint      `json:"courseId" gorm:"not null;default:null"`
	Course   Course    `json:"course"`
	Rating   float64   `json:"rating" gorm:"not null;default:null"`
}
