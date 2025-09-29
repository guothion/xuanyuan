package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/internal/api/controller"
	"github.com/guothion/xuanyuan/internal/middleware"
	"github.com/guothion/xuanyuan/internal/service/account"
)

func init() {
	c := &RecipeRoute{}
	routes[c.BasePath()] = c
}

type RecipeRoute struct{}

func (r *RecipeRoute) BasePath() string {
	return "/recipes"
}

func (r *RecipeRoute) RegisterRouter(router *gin.RouterGroup) {
	tecipeController := &controller.RecipesController{}
	tecipeRouter := router.Group(r.BasePath()).Use(middleware.JWTAuth(account.AppGuardName))
	{
		// 创建
		tecipeRouter.POST("/create", tecipeController.CreateRecipes)
		// 获取列表
		tecipeRouter.GET("", tecipeController.GetRecipesList)
	}
}
