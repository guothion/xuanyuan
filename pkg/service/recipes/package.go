package recipes

import "fmt"

func init() {
	fmt.Println("xuanyuan recipesSerivce is initialized")
}

var (
	RecipesService = &recipesService{}
)
