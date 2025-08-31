package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/common/request"
	"github.com/guothion/xuanyuan/pkg/api/common/response"
	"github.com/guothion/xuanyuan/pkg/service/recipes"
)

type RecipesController struct{}

func (r *RecipesController) CreateRecipes(c *gin.Context) {
	var form request.CreateRecipeRequest
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, recipe := recipes.RecipesService.CreateRecipes(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, recipe)
	}
}

func (r *RecipesController) GetRecipesList(c *gin.Context) {
	var query request.GetRecipeListRequest = request.GetRecipeListRequest{
		Page:     1,
		PageSize: 10,
	}
	if err := c.ShouldBind(&query); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(query, err))
		return
	}
	if err, recipes := recipes.RecipesService.GetRecipesList(query); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, recipes)
	}
}
