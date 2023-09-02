package models

import (
	"errors"
	"time"

	mailer "github.com/curiousz-peel/web-learning-platform-backend/mailer"
	validator "github.com/curiousz-peel/web-learning-platform-backend/validator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey"`
	FirstName   string       `json:"firstName" gorm:"not null;default:null"`
	LastName    string       `json:"lastName" gorm:"not null;default:null"`
	UserName    string       `json:"userName" gorm:"unique;not null;default:null"`
	Email       string       `json:"email" gorm:"unique;not null;default:null"`
	Password    string       `json:"password" gorm:"not null;default:null"`
	Birthday    time.Time    `json:"birthday" gorm:"default:null"`
	IsMod       bool         `json:"isModerator" gorm:"default:false"`
	IsAuthor    bool         `json:"isAuthor" gorm:"default:false"`
	Ratings     []Rating     `json:"ratings" gorm:"default:null;-"`
	Comments    []Comment    `json:"comments" gorm:"default:null;-"`
	Enrollments []Enrollment `json:"enrollments" gorm:"default:null;-"`
}

type UpdateUser struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	UserName  string    `json:"userName" `
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Birthday  time.Time `json:"birthday"`
	IsMod     bool      `json:"isModerator"`
}

type LoginUser struct {
	UserName string `json:"user"`
	Password string `json:"pass"`
}

type UserDTO struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	UserName  string    `json:"userName" `
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Birthday  time.Time `json:"birthday"`
	IsMod     bool      `json:"isModerator"`
	IsAuthor  bool      `json:"isAuthor"`
}

func ToUserDTO(user User) UserDTO {
	return UserDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  user.Password,
		Birthday:  user.Birthday,
		IsMod:     user.IsMod,
		IsAuthor:  user.IsAuthor,
	}
}

func (u *User) isValid() (err error) {
	err = validator.ValidatePassword(u.Password)
	if err != nil {
		return err
	}
	err = validator.ValidateEmail(u.Email)
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

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	m := mailer.Mailer{AddressTo: u.Email}
	go m.SendRegistrationEmail()
	return
}
