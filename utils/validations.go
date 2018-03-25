package utils

import (
	"github.com/asaskevich/govalidator"
)

func IsValidEmail(id string) bool {
	return govalidator.IsEmail(id)
}

func IsValidPassword(pwd string) bool {
	return len(pwd) > 8
}
