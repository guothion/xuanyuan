package model

type Recipe struct {
	ID
	Name         string `json:"name"gorm:"not null;comment:菜名"`
	Description  string `json:"description"gorm:"comment:简要描述"`
	Instructions string `json:"instructions"gorm:"comment:做法步骤（支持 Markdown 或 HTML）"`
	PrepTime     int    `json:"prep_time"gorm:"comment:准备时间（分钟）"`
	CookTime     int    `json:"cook_time"gorm:"comment:烹饪时间（分钟）"`
	Servings     int    `json:"servings"gorm:"comment:分量（几人份）"`
	ImageUrl     string `json:"image_url"gorm:"comment:封面图 URL"`
	Timestamp
	Status string `json:"status" gorm:"not null;comment:状态;default:published"`
}

type RecipesList struct {
	Data     []Recipe `json:"data"`
	Total    int64    `json:"total"`
	PageSize int      `json:"page_size"`
	Page     int      `json:"page"`
}
