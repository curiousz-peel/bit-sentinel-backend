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
	UserID      uuid.UUID `json:"userId"`
	Profession  string    `json:"profession"`
	Description string    `json:"description"`
	Topics      []string  `json:"topics"`
}

func (a *Author) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
