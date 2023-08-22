package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	Title       string         `json:"title" gorm:"not null;default:null"`
	Description string         `json:"description" gorm:"not null;default:null"`
	QuestionIDs datatypes.JSON `json:"questionIDs" gorm:"default:null;type:text[]"`
	Questions   []Question     `json:"questions" gorm:"foreignKey:QuestionIDs;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;-"`
	//change courseID and lessonID to be mandatory after all CRUDS are created
	CourseID uint   `json:"courseId" gorm:"default:null"`
	Course   Course `gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LessonID uint   `json:"lessonId" gorm:"default:null"`
	Lesson   Lesson `gorm:"foreignKey:LessonID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UpdateQuiz struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	QuestionIDs []uint `json:"questionIDs"`
	CourseID    uint   `json:"courseId"`
	LessonID    uint   `json:"lessonId"`
}

type AddQuestionsToQuiz struct {
	QuestionIDs []uint `json:"questionIds"`
}

type Question struct {
	gorm.Model
	Text      string         `json:"text" gorm:"not null;default:null"`
	OptionIDs datatypes.JSON `json:"optionIDs" gorm:"default:null;type:text[]"`
	Options   []Option       `json:"options" gorm:"foreignKey:OptionIDs;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;-"`
	QuizID    uint           `json:"quizId" gorm:"not null;default:null"`
	Quiz      Quiz           `gorm:"foreignKey:QuizID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UpdateQuestion struct {
	Text      string `json:"text"`
	OptionIDs []uint `json:"optionIds"`
	QuizID    uint   `json:"quizId"`
}

type AddOptionsToQuestion struct {
	OptionIDs []uint `json:"optionIds"`
}

type Option struct {
	gorm.Model
	Text       string   `json:"text" gorm:"not null;default:null"`
	IsCorrect  bool     `json:"isCorrect" gorm:"not null;default:null"`
	QuestionID uint     `json:"questionId" gorm:"not null;default:null"`
	Question   Question `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UpdateOption struct {
	QuestionID uint   `json:"questionID"`
	IsCorrect  bool   `json:"isCorrect"`
	Text       string `json:"text"`
}

type DisplayOption struct {
	ID         uint   `json:"ID"`
	QuestionID uint   `json:"questionID"`
	IsCorrect  bool   `json:"isCorrect"`
	Text       string `json:"text"`
}
