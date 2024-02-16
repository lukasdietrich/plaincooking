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

type ReadRecipeRequest struct {
	ID xid.ID `json:"-" param:"recipeId"`
}

// @summary  Read recipe
// @tags     recipes
// @router   /recipes/{recipeId}  [get]
// @produce  application/json
// @param    recipeId path  string     true  "Recipe ID"
// @success  200  {object}  RecipeDto
func readRecipe(ctx echo.Context, querier models.Querier) error {
	var req ReadRecipeRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	params := models.ReadRecipeParams{
		ID: req.ID,
	}

	recipe, err := querier.ReadRecipe(ctx.Request().Context(), params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, mapRecipeDto(recipe))
}

type UpdateRecipeRequest struct {
	ID    xid.ID `json:"-" param:"recipeId"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type UpdateRecipeResponse struct {
	ID xid.ID `json:"id"`
}

// @summary  Update a recipe
// @tags     recipes
// @router   /recipes/{recipeId}  [put]
// @accept   application/json
// @produce  application/json
// @param    recipeId path  string                true  "Recipe ID"
// @param    request body   UpdateRecipeRequest   true  "Recipe"
// @success  200  {object}  UpdateRecipeResponse
func updateRecipe(ctx echo.Context, querier models.Querier) error {
	var req UpdateRecipeRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	params := models.UpdateRecipeParams{
		ID:        req.ID,
		UpdatedAt: time.Now(),
		Slug:      req.Slug,
		Title:     req.Title,
	}

	recipe, err := querier.UpdateRecipe(ctx.Request().Context(), params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, UpdateRecipeResponse{
		ID: recipe.ID,
	})
}

type DeleteRecipeRequest struct {
	ID xid.ID `param:"recipeId"`
}

type DeleteRecipeResponse struct {
	ID xid.ID `json:"id"`
}

// @summary  Delete a recipe
// @tags     recipes
// @router   /recipes/{recipeId}  [delete]
// @accept   application/json
// @produce  application/json
// @param    recipeId path  string                true  "Recipe ID"
// @success  200  {object}  DeleteRecipeResponse
func deleteRecipe(ctx echo.Context, querier models.Querier) error {
	var req DeleteRecipeRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	params := models.DeleteRecipeParams{
		ID: req.ID,
	}

	recipe, err := querier.DeleteRecipe(ctx.Request().Context(), params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, DeleteRecipeResponse{
		ID: recipe.ID,
	})
}
