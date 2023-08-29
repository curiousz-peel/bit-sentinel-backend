package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	UserID   uuid.UUID `json:"userId" gorm:"not null;default:null"`
	User     User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CourseID uint      `json:"courseId" gorm:"default:null;index:type:btree"`
	Course   Course    `gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rating   float64   `json:"rating" gorm:"not null;default:null"`
}

type UpdateRating struct {
	UserID   uuid.UUID `json:"userID"`
	CourseID int       `json:"courseID"`
	Rating   float64   `json:"rating"`
}

type RatingDTO struct {
	ID       uint      `json:"id"`
	UserID   uuid.UUID `json:"userID"`
	CourseID uint      `json:"courseID"`
	Rating   float64   `json:"rating"`
}

func ToRatingDTO(rating Rating) RatingDTO {
	return RatingDTO{
		ID:       rating.ID,
		UserID:   rating.UserID,
		CourseID: rating.CourseID,
		Rating:   rating.Rating,
	}
}
