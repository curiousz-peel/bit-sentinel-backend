package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Title                 string         `json:"title" gorm:"unique;not null;default:null"`
	Description           string         `json:"description" gorm:"not null;default:null"`
	AuthorsIDs            datatypes.JSON `json:"authorIds" gorm:"not null;default:null;type:text[]"`
	Authors               []Author       `json:"authors" gorm:"foreignKey:AuthorsIDs;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;-"`
	Lessons               []Lesson       `gorm:"foreignKey:LessonsIDs;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;-"`
	Tags                  datatypes.JSON `json:"tags" gorm:"not null;default:null;type:text[]"`
	Visible               bool           `json:"visible"`
	Rating                float64        `json:"rating" gorm:"not null;default:0"`
	Ratings               []Rating       `gorm:"foreignKey:RatingIDs;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;-"`
	Comments              []Comment      `gorm:"foreignKey:CommentIDs;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;-"`
	IncludedSubscriptions datatypes.JSON `json:"subscriptions" gorm:"default:null;type:text[]"`
}

type UpdateCourse struct {
	Title                 string         `json:"title"`
	Description           string         `json:"description"`
	AuthorsIDs            datatypes.JSON `json:"authorIds"`
	Tags                  datatypes.JSON `json:"tags"`
	Visible               bool           `json:"visible"`
	Rating                float64        `json:"rating"`
	IncludedSubscriptions datatypes.JSON `json:"subscriptions"`
}

type CourseDTO struct {
	ID                    uint           `json:"id"`
	Title                 string         `json:"title"`
	Description           string         `json:"description"`
	Tags                  datatypes.JSON `json:"tags"`
	Visible               bool           `json:"visible"`
	Rating                float64        `json:"rating"`
	RatingNo              int            `json:"ratingNo"`
	IncludedSubscriptions datatypes.JSON `json:"subscriptions"`
	AuthorsIDs            datatypes.JSON `json:"authorIds"`
}

type AddAuthorsToCourse struct {
	AuthorsIDs []uuid.UUID `json:"authorsIds"`
}

func ToCourseDTO(course Course) CourseDTO {
	return CourseDTO{
		ID:                    course.ID,
		Title:                 course.Title,
		Description:           course.Description,
		Tags:                  course.Tags,
		Visible:               course.Visible,
		Rating:                course.Rating,
		RatingNo:              len(course.Ratings),
		IncludedSubscriptions: course.IncludedSubscriptions,
		AuthorsIDs:            course.AuthorsIDs,
	}
}
