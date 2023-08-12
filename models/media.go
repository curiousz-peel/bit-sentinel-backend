package models

import (
	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	LessonID     uint      `json:"lessonId" gorm:"not null;default:null"`
	Lesson       Lesson    `gorm:"foreignKey:LessonID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;type:text"`
	FilePath     string    `json:"filePath" gorm:"unique;not null;default:null"`
	FileTypeName string    `json:"fileType" gorm:"not null;default:null"`
	FileType     MediaType `gorm:"foreignKey:FileTypeName;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;type:text"`
}

type MediaType struct {
	gorm.Model
	Type string `json:"type" gorm:"unique;not null;default:null;primarykey"`
}
