package util

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var (
	mobileReg = `^1[3-9]\d{9}$`
	emailReg  = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

type ValidatFun func(string) bool

func getValidator(regStr string) ValidatFun {
	return func(str string) (ok bool) {
		ok, _ = regexp.MatchString(regStr, str)
		if !ok {
			return false
		}
		return true
	}
}

var IsMobile = getValidator(mobileReg)
var IsEmail = getValidator(emailReg)

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	return IsMobile(mobile)
}

func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	return IsEmail(email)
}
