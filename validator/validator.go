package validator

import (
	"errors"
	"net/mail"
	"unicode"
)

func ValidatePassword(pass string) (err error) {
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

func ValidateEmail(email string) (err error) {
	_, err = mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email")
	}
	return
}
