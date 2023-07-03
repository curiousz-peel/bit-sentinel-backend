package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID `json:"userId" gorm:"not null;unique"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Profession  string    `json:"profession"`
	Description string    `json:"description"`
	Topics      string    `json:"topics"`
}

func (a *Author) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
