package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Lesson struct {
	gorm.Model
	Title      string         `json:"title" gorm:"unique;not null;default:null"`
	Order      int            `json:"order" gorm:"not null;default:null;uniqueIndex:idx_course_order"`
	CourseID   uint           `json:"courseId" gorm:"not null;default:null;uniqueIndex:idx_course_order;index:type:btree"`
	Course     Course         `gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Summary    string         `json:"summary" gorm:"not null;default:null"`
	ContentIds datatypes.JSON `json:"mediaIds" gorm:"default:null;type:text[]"`
	Content    []Media        `json:"content" gorm:"foreignKey:ContentIds;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;-"`
	Quizzes    []Quiz         `json:"quizzes" gorm:"default:null;-"`
}

type UpdateLesson struct {
	Title      string         `json:"title"`
	Order      int            `json:"order"`
	CourseID   uint           `json:"courseId"`
	Summary    string         `json:"summary"`
	ContentIds datatypes.JSON `json:"mediaIds"`
}

type LessonDTO struct {
	ID         uint           `json:"id"`
	Title      string         `json:"title"`
	Order      int            `json:"order"`
	CourseID   uint           `json:"courseId"`
	Summary    string         `json:"summary"`
	ContentIds datatypes.JSON `json:"mediaIds"`
}

type AddContentsToLesson struct {
	ContentIDs []uint `json:"contentIds"`
}

func ToLessonDTO(lesson Lesson) LessonDTO {
	return LessonDTO{
		ID:         lesson.ID,
		Title:      lesson.Title,
		Order:      lesson.Order,
		CourseID:   lesson.CourseID,
		Summary:    lesson.Summary,
		ContentIds: lesson.ContentIds,
	}
}
