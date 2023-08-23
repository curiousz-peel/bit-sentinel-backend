package models

import (
	"time"

	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	LessonID     uint      `json:"lessonId" gorm:"not null;default:null"`
	Lesson       Lesson    `gorm:"foreignKey:LessonID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FilePath     string    `json:"filePath" gorm:"unique;not null;default:null"`
	FileTypeName string    `json:"fileType" gorm:"not null;default:null"`
	FileType     MediaType `gorm:"foreignKey:FileTypeName;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type MediaType struct {
	ID        uint           `gorm:"autoIncrement"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Type      string         `json:"type" gorm:"unique;not null;default:null;primarykey"`
}

type UpdateMedia struct {
	LessonID     uint   `json:"lessonID"`
	FilePath     string `json:"filePath"`
	FileTypeName string `json:"fileTypeName"`
}

type UpdateMediaType struct {
	Type string `json:"type"`
}
