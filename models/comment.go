package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID   uuid.UUID `json:"userId" gorm:"not null;default:null"`
	User     User      `json:"user"`
	CourseID uint      `json:"courseId"`
	Course   Course    `json:"course"`
	Text     string    `json:"text" gorm:"not null;default:null"`
}
