package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID   uuid.UUID `json:"userId" gorm:"not null;default:null"`
	User     User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CourseID uint      `json:"courseId" gorm:"not nul;default:null;index:type:btree"`
	Course   Course    `gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Text     string    `json:"text" gorm:"not null;default:null"`
}

type UpdateComment struct {
	UserID   uuid.UUID `json:"userID"`
	CourseID int       `json:"courseID"`
	Text     string    `json:"text"`
}

type CommentDTO struct {
	ID       uint      `json:"id"`
	UserID   uuid.UUID `json:"userId"`
	CourseID uint
	Text     string
}

func ToCommentDTO(comment Comment) CommentDTO {
	return CommentDTO{
		ID:       comment.ID,
		UserID:   comment.User.ID,
		CourseID: comment.CourseID,
		Text:     comment.Text,
	}
}
