package request

type CreateRecipeRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required,min=20"`
	Instruction string `json:"instruction" binding:"required"`
	PrepTime    int    `json:"prep_time" binding:"required"`
	CookTime    int    `json:"cook_time" binding:"required"`
	Servings    int    `json:"servings" binding:"required"`
	ImageUrl    string `json:"image_url"`
	Status      string `json:"status" binding:"required,oneof=draft published archived"`
}

func (crr CreateRecipeRequest) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":        "Name is required",
		"description.required": "Description is required",
		"description.min":      "Description lenth is more than 20 characters",
		"instruction.required": "Instruction is required",
		"prep_time.required":   "PrepTime is required",
		"cook_time.required":   "CookTime is required",
		"servings.required":    "Servings is required",
		"status.required":      "Status is required",
		"status.oneof":         "Status is wrong",
	}
}

type GetRecipeListRequest struct {
	Page     int `json:"page" binding:"required"`
	PageSize int `json:"page_size" binding:"required"`
}

func (grlr GetRecipeListRequest) GetMessages() ValidatorMessages {
	return ValidatorMessages{}
}
