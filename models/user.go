package models

import (
	"errors"
	"net/mail"
	"time"
	"unicode"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	FirstName string    `json:"firstName" gorm:"not null;default:null"`
	LastName  string    `json:"lastName" gorm:"not null;default:null"`
	UserName  string    `json:"userName" gorm:"unique;not null;default:null"`
	Email     string    `json:"email" gorm:"unique;not null;default:null"`
	Password  string    `json:"password" gorm:"not null;default:null"`
	Birthday  time.Time `json:"birthday" gorm:"not null;default:null"`
	IsMod     bool      `json:"isModerator" gorm:"default:false"`
}

func validatePassword(pass string) (err error) {
	var upper, lower, special, number bool
	for _, char := range pass {
		switch {
		case unicode.IsNumber(char):
			number = true
		case unicode.IsLower(char):
			lower = true
		case unicode.IsUpper(char):
			upper = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			special = true
		default:
		}
	}

	if !(lower && upper && number && special && (len(pass) >= 6)) {
		return errors.New("password must be of length >= 6 and have at least 1 of :digit, upper case letter, lower case letter and special character")
	}
	return nil
}

func validateEmail(email string) (err error) {
	_, err = mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email")
	}
	return
}

func (u *User) isValid() (err error) {
	err = validatePassword(u.Password)
	if err != nil {
		return err
	}
	err = validateEmail(u.Email)
	if err != nil {
		return err
	}
	if !(len(u.FirstName) >= 2) || !(len(u.LastName) >= 2) || !(len(u.UserName) >= 4) {
		err = errors.New("first, last or user name is too short")
	}
	return
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	err = u.isValid()
	return
}
