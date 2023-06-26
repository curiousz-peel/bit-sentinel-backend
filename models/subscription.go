package models

import (
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	Type     string  `json:"type" gorm:"unique"`
	Duration int     `json:"duration"`
	Price    float32 `json:"price"`
}
