package web

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/xid"

	"github.com/lukasdietrich/plaincooking/internal/database/models"
)

// @title     PlainCooking API
// @version   0.1
// @host      http://localhost:8080
// @basePath  /api

type RecipeDto struct {
	ID        xid.ID    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
}

// @summary  List recipes
// @tags     recipes
// @router   /recipes  [get]
// @produce  application/json
// @success  200  {object}  []RecipeDto
func listRecipes(ctx echo.Context, querier models.Querier) error {
	recipeSlice, err := querier.ListRecipes(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, mapSlice(recipeSlice, mapRecipeDto))
}

type CreateRecipeRequest struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type CreateRecipeResponse struct {
	ID xid.ID `json:"id"`
}

// @summary  Create a new recipe
// @tags     recipes
// @router   /recipes  [post]
// @accept   application/json
// @produce  application/json
// @param    request body   CreateRecipeRequest  true  "Recipe"
// @success  200  {object}  CreateRecipeResponse
func createRecipe(ctx echo.Context, querier models.Querier) error {
	var req CreateRecipeRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	params := models.CreateRecipeParams{
		ID:        xid.New(),
		CreatedAt: time.Now(),
		Slug:      req.Slug,
		Title:     req.Title,
	}

	recipe, err := querier.CreateRecipe(ctx.Request().Context(), params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, CreateRecipeResponse{
		ID: recipe.ID,
	})
}
