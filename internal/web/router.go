package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(recipes *RecipeController) http.Handler {
	r := echo.New()
	r.Binder = new(binder)

	r.Use(middleware.Recover())
	r.Use(logError())
	r.Use(handleBusinessError())

	api := r.Group("/api")
	api.GET("/recipes", recipes.List)
	api.POST("/recipes", recipes.Create)
	api.GET("/recipes/:recipeId", recipes.Read)
	api.PUT("/recipes/:recipeId", recipes.Update)
	api.DELETE("/recipes/:recipeId", recipes.Delete)

	return r
}
