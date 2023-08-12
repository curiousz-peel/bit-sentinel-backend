package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Title      string      `json:"title" gorm:"unique;not null;default:null"`
	AuthorsIDs []uuid.UUID `json:"authorIds" gorm:"not null;default:null;type:text"`
	Authors    []Author    `json:"authors" gorm:"type:text"`
	LessonsIDs []uint      `json:"lessonIds" gorm:"not null;default:null;type:text"`
	Lessons    []Lesson    `json:"lessons" gorm:"type:text"`
	Tags       []string    `json:"tags" gorm:"not null;default:null;type:text"`
	Visible    bool        `json:"visible"`
	Rating     float64     `json:"rating" gorm:"not null;default:0"`
	Ratings    []Rating    `json:"ratings" gorm:"type:text"`
	// CommentIDs             []uint      `json:"commentIds"`
	Comments               []Comment `json:"comments" gorm:"type:text"`
	SubscriptionsAvailable []string  `json:"subscriptions" gorm:"type:text"`
}
