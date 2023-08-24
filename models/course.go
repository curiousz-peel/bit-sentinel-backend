package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Title                 string         `json:"title" gorm:"unique;not null;default:null"`
	AuthorsIDs            datatypes.JSON `json:"authorIds" gorm:"not null;default:null;type:text[]"`
	Authors               []Author       `json:"authors" gorm:"foreignKey:AuthorsIDs;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;-"`
	Lessons               []Lesson       `gorm:"foreignKey:LessonsIDs;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;-"`
	Tags                  []string       `json:"tags" gorm:"not null;default:null;type:text[]"`
	Visible               bool           `json:"visible"`
	Rating                float64        `json:"rating" gorm:"not null;default:0"`
	Ratings               []Rating       `gorm:"foreignKey:RatingIDs;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;-"`
	Comments              []Comment      `gorm:"foreignKey:CommentIDs;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;-"`
	IncludedSubscriptions []string       `json:"subscriptions" gorm:"default:null;type:text[]"`
}

type UpdateCourse struct {
	Title                 string         `json:"title"`
	AuthorsIDs            datatypes.JSON `json:"authorIds"`
	Tags                  []string       `json:"tags"`
	Visible               bool           `json:"visible"`
	Rating                float64        `json:"rating"`
	IncludedSubscriptions []string       `json:"subscriptions"`
}

type AddAuthorsToCourse struct {
	AuthorsIDs []uuid.UUID `json:"authorsIds"`
}
