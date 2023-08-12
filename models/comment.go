package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID   uuid.UUID `json:"userId" gorm:"not null;default:null"`
	User     User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;type:text"`
	CourseID uint      `json:"courseId"`
	Course   Course    `gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;type:text"`
	Text     string    `json:"text" gorm:"not null;default:null"`
}
