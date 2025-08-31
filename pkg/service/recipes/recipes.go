package recipes

import (
	"github.com/guothion/xuanyuan/pkg/api/common/request"
	"github.com/guothion/xuanyuan/pkg/mapper"
	"github.com/guothion/xuanyuan/pkg/model"
)

type recipesService struct{}

func (r *recipesService) CreateRecipes(crr request.CreateRecipeRequest) (err error, recipes model.Recipe) {
	err, recipes = mapper.Recipes.CreateRow(crr)
	return
}

func (r *recipesService) GetRecipesList(crr request.GetRecipeListRequest) (err error, recipesList model.RecipesList) {
	err, recipesList = mapper.Recipes.GetList(crr)
	return
}
