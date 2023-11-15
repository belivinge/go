package main

import "fmt"

type User struct {
	name  string
	email string
	phone string
}

type UserValidation struct {
	User
	emailValidator EmailValidator
	phoneValidator PhoneValidator
}

type EmailValidator interface {
	isValid(email string) bool
}

type PhoneValidator interface {
	isValid(phone string) bool
}

func (uv *UserValidation) Validate() error {
	if !uv.emailValidator.isValid(uv.User.email) || !uv.phoneValidator.isValid(uv.User.phone) {
		return fmt.Errorf("invalid")
	}
	return nil
}
