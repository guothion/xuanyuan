package request

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	GetMessages() ValidatorMessages
}

type ValidatorMessages map[string]string

// 这个方法只能是Validator 类才可以调用
func GetErrorMsg(request interface{}, err error) string {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		_, isValidator := request.(Validator)

		for _, v := range err.(validator.ValidationErrors) {
			if isValidator {
				if message, exist := request.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exist {
					return message
				}
			}
			return v.Error()
		}
	}

	return err.Error()
}

// 自定义校验
type CustomValidator struct {
	Validator *validator.Validate
}

// 这里我们可以写一些业务相关的自定义注册器
func RegisterCustomValidators(v *validator.Validate) {
	// register：根据 type 判断登录
	v.RegisterValidation("login_type", ValidateLoginType)
}
