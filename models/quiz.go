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
	CourseID    uint           `json:"courseId" gorm:"not null;default:null"`
	Course      Course         `gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LessonID    uint           `json:"lessonId" gorm:"not null;default:null"`
	Lesson      Lesson         `gorm:"foreignKey:LessonID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UpdateQuiz struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	QuestionIDs []uint `json:"questionIDs"`
	CourseID    uint   `json:"courseId"`
	LessonID    uint   `json:"lessonId"`
}

type QuizDTO struct {
	ID          uint           `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	QuestionIDs datatypes.JSON `json:"questionIDs"`
	CourseID    uint           `json:"courseId"`
	LessonID    uint           `json:"lessonId"`
}

func ToQuizDTO(quiz Quiz) QuizDTO {
	return QuizDTO{
		ID:          quiz.ID,
		Title:       quiz.Title,
		Description: quiz.Description,
		QuestionIDs: quiz.QuestionIDs,
		CourseID:    quiz.CourseID,
		LessonID:    quiz.LessonID,
	}
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

type QuestionDTO struct {
	ID        uint           `json:"id"`
	Text      string         `json:"text"`
	OptionIDs datatypes.JSON `json:"optionIds"`
	QuizID    uint           `json:"quizId"`
}

func ToQuestionDTO(question Question) QuestionDTO {
	return QuestionDTO{
		ID:        question.ID,
		Text:      question.Text,
		OptionIDs: question.OptionIDs,
		QuizID:    question.QuizID,
	}
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

type OptionDTO struct {
	ID         uint   `json:"ID"`
	QuestionID uint   `json:"questionID"`
	IsCorrect  bool   `json:"isCorrect"`
	Text       string `json:"text"`
}

func ToOptionDTO(option Option) OptionDTO {
	return OptionDTO{
		ID:         option.ID,
		QuestionID: option.QuestionID,
		IsCorrect:  option.IsCorrect,
		Text:       option.Text,
	}
}
