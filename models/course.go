package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Title      string      `json:"title" gorm:"unique;not null;default:null"`
	AuthorsIDs []uuid.UUID `json:"authorIds" gorm:"not null;default:null"`
	Authors    []Author    `json:"authors"`
	LessonsIDs []uint      `json:"lessonIds" gorm:"not null;default:null"`
	Lessons    []Lesson    `json:"lessons"`
	Tags       []string    `json:"tags" gorm:"not null;default:null"`
	Visible    bool        `json:"visible"`
	Rating     float64     `json:"rating" gorm:"not null;default:0"`
	Ratings    []Rating    `json:"ratings"`
	// CommentIDs             []uint      `json:"commentIds"`
	Comments               []Comment `json:"comments"`
	SubscriptionsAvailable []string  `json:"subscriptions"`
}
