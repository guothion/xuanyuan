package mapper

import (
	"github.com/guothion/xuanyuan/pkg/api/common/request"
	"github.com/guothion/xuanyuan/pkg/global"
	"github.com/guothion/xuanyuan/pkg/model"
)

type recipesMapper struct{}

func (recipesMapper *recipesMapper) CreateRow(crr request.CreateRecipeRequest) (err error, recipes model.Recipe) {
	recipes = model.Recipe{
		Name:         crr.Name,
		Description:  crr.Description,
		Instructions: crr.Instruction,
		PrepTime:     crr.PrepTime,
		CookTime:     crr.CookTime,
		Servings:     crr.Servings,
		ImageUrl:     crr.ImageUrl,
		Status:       crr.Status,
	}
	err = global.App.DB.Create(&recipes).Error
	return
}

func (recipesMapper *recipesMapper) GetList(listRequest request.GetRecipeListRequest) (err error, recipes model.RecipesList) {
	var total int64
	if err = global.App.DB.Model(&model.Recipe{}).Count(&total).Error; err != nil {
		recipes.Data = nil
		return
	}
	pageSize := listRequest.PageSize
	page := listRequest.Page
	offset := pageSize * (page - 1)
	recipes.PageSize = pageSize
	recipes.Page = page
	recipes.Total = total
	if err = global.App.DB.Model(&model.Recipe{}).Offset(offset).Limit(pageSize).Find(&recipes.Data).Error; err != nil {
		return
	}
	return
}
