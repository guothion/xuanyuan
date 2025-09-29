package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/guothion/xuanyuan/internal/util"
)

// todo: register
type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
	Role     int8   `form:"role" json:"role" binding:"required"`
}

func (Register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":     "name is required",
		"email.required":    "Email is required",
		"email.email":       "Email is wrong",
		"password.required": "password is required",
	}
}

// todo: login
type Login struct {
	Type     string `form:"type" json:"type" binding:"required,oneof=account email,login_type"`
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"type.required":     "type is required",
		"type.oneof":        "type is in account and email",
		"type.login_type":   "account need username and email need email",
		"password.required": "Password is required",
	}
}

func ValidateLoginType(fl validator.FieldLevel) bool {
	typeVal := fl.Parent().FieldByName("Type").String()
	username := fl.Parent().FieldByName("Username").String()
	email := fl.Parent().FieldByName("Email").String()
	switch typeVal {
	case "account":
		return username != ""
	case "email":
		return util.IsEmail(email)
	}
	return false // 其他 type 不合法
}
