package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID      `json:"userId" gorm:"not null;unique"`
	User        User           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Profession  string         `json:"profession" gorm:"not null;default:null;size:255;min:3"`
	Description string         `json:"description" gorm:"not null;default:null"`
	Topics      datatypes.JSON `json:"topics" gorm:"default:null;type:text[]"`
}

type UpdateAuthor struct {
	UserID      uuid.UUID      `json:"userId"`
	Profession  string         `json:"profession"`
	Description string         `json:"description"`
	Topics      datatypes.JSON `json:"topics"`
}

type AuthorDTO struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID      `json:"userId"`
	Profession  string         `json:"profession"`
	Description string         `json:"description"`
	Topics      datatypes.JSON `json:"topics"`
	FirstName   string         `json:"firstName"`
	LastName    string         `json:"lastName"`
}

func (a *Author) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}

func ToAuthorDTO(author Author) AuthorDTO {
	return AuthorDTO{
		ID:          author.ID,
		UserID:      author.UserID,
		Profession:  author.Profession,
		Description: author.Description,
		Topics:      author.Topics,
		FirstName:   author.User.FirstName,
		LastName:    author.User.LastName,
	}
}
