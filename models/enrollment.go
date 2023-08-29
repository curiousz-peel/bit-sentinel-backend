package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enrollment struct {
	gorm.Model
	UserID     uuid.UUID `json:"userId" gorm:"uniqueIndex:idx_user_course;not null;default:null"`
	User       User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CourseID   uint      `json:"courseId" gorm:"uniqueIndex:idx_user_course;not null;default:null"`
	Course     Course    `gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProgressID uint      `json:"progressId" gorm:"default:null"`
	Progress   Progress  `gorm:"foreignKey:ProgressID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UpdateEnrollment struct {
	UserID     uuid.UUID `json:"userID"`
	CourseID   int       `json:"courseID"`
	ProgressID int       `json:"progressID"`
}

type EnrollmentDTO struct {
	ID         uint      `json:"id"`
	UserID     uuid.UUID `json:"userId"`
	CourseID   uint      `json:"courseId"`
	ProgressID uint      `json:"progressId"`
}

func ToEnrollmentDTO(enrollment Enrollment) EnrollmentDTO {
	return EnrollmentDTO{
		ID:         enrollment.ID,
		UserID:     enrollment.UserID,
		CourseID:   enrollment.CourseID,
		ProgressID: enrollment.ProgressID,
	}
}
