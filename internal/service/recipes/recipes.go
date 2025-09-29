package recipes

import (
	"github.com/guothion/xuanyuan/internal/api/common/request"
	"github.com/guothion/xuanyuan/internal/mapper"
	"github.com/guothion/xuanyuan/internal/model"
)

type recipesService struct{}

func (r *recipesService) CreateRecipes(crr request.CreateRecipeRequest) (err error, recipes model.Recipe) {
	recipes, err = mapper.Recipes.CreateRow(crr)
	return
}

func (r *recipesService) GetRecipesList(crr request.GetRecipeListRequest) (err error, recipesList model.RecipesList) {
	recipesList, err = mapper.Recipes.GetList(crr)
	return
}
