package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enrollment struct {
	gorm.Model
	UserID     uuid.UUID `json:"userId" gorm:"uniqueIndex:idx_user_course;not null;default:null"`
	User       User      `json:"user"`
	CourseID   uint      `json:"courseId" gorm:"uniqueIndex:idx_user_course;not null;default:null"`
	Course     Course    `json:"course"`
	ProgressID uint      `json:"progressId"`
	Progress   Progress  `json:"progress"`
}
