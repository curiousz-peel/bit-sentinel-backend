package models

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	Title       string     `json:"title" gorm:"not null;default:null"`
	Description string     `json:"description" gorm:"not null;default:null"`
	Questions   []Question `json:"questions"`
	CourseID    uint       `json:"courseId" gorm:"not null;default:null"`
	Course      Course     `json:"course"`
	LessonID    uint       `json:"lessonId" gorm:"not null;default:null"`
	Lesson      Lesson     `json:"lesson"`
}

type Question struct {
	gorm.Model
	Text    string   `json:"text" gorm:"not null;default:null"`
	Options []Option `json:"options"`
	QuizID  uint     `json:"quizId" gorm:"not null;default:null"`
	Quiz    Quiz     `json:"quiz"`
}

type Option struct {
	gorm.Model
	Text       string   `json:"text" gorm:"not null;default:null"`
	IsCorrect  bool     `json:"isCorrect" gorm:"not null;default:null"`
	QuestionID uint     `json:"questionId" gorm:"not null;default:null"`
	Question   Question `json:"question"`
}
