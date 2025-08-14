package request

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Role     int8   `form:"role" json:"role" binding:"required"`
}

func (rg *Register) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"name.required":     "name is required",
		"email.required":    "email is required",
		"password.required": "password is required",
		"email.isEmail":     "Email is wrong",
	}
}

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"email.required":    "Email is required",
		"email.isEmail":     "Email is wrong",
		"password.required": "Password is required",
	}
}
