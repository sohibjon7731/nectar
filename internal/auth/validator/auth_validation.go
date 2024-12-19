package validator

import (
	"errors"
	"regexp"
)

func EmailValidation(email string) error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(regex, email)
	if !matched {
		return errors.New("it is not email. try again")
	}
	return nil
}

func PasswordValidation(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}
