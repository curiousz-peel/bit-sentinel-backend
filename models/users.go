package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName *string    `json:"firstName"`
	LastName  *string    `json:"lastName"`
	Email     *string    `json:"email"`
	Password  *string    `json:"password"`
	Birthday  *time.Time `json:"birthday"`
	IsMod     bool       `json:"isModerator"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	if err != nil {
		log.Fatal("failed to auto migrate table users")
		return err
	}
	return nil
}
