package web

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/xid"

	"github.com/lukasdietrich/plaincooking/internal/service"
)

const (
	MIMEMarkdown = "text/markdown"
)

type ApiError echo.HTTPError // @name ApiError

// @title     PlainCooking API
// @version   0.1
// @host      http://localhost:8080
// @basePath  /api

type RecipeController struct {
	recipes *service.RecipeService
}

func NewRecipeController(recipes *service.RecipeService) *RecipeController {
	return &RecipeController{
		recipes: recipes,
	}
}

type RecipeMetadataDto struct {
	ID    xid.ID `json:"id"`
	Title string `json:"title"`
} // @name RecipeMetadata

// @summary  List recipes
// @id       listRecipes
// @tags     recipes
// @router   /recipes  [get]
// @produce  application/json
// @success  200  {array}  RecipeMetadataDto
func (c *RecipeController) List(ctx echo.Context) error {
	recipeSlice, err := c.recipes.List(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, mapSlice(recipeSlice, mapRecipeMetadataDto))
}

type CreateRecipeResponse struct {
	ID xid.ID `json:"id"`
} // @name CreateRecipeResponse

// @summary  Create a new recipe
// @id       createRecipe
// @tags     recipes
// @router   /recipes  [post]
// @accept   text/markdown
// @produce  application/json
// @param    content  body  string  true  "Recipe content"
// @success  201  {object}  CreateRecipeResponse
// @failure  400  {object}  ApiError
// @failure  422  {object}  ApiError
func (c *RecipeController) Create(ctx echo.Context) error {
	content, err := c.readMarkdownRequest(ctx)
	if err != nil {
		return err
	}

	id, err := c.recipes.Create(ctx.Request().Context(), content)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, CreateRecipeResponse{ID: id})
}

type ReadRecipeRequest struct {
	ID xid.ID `json:"-" param:"recipeId"`
}

// @summary  Read a recipe
// @id       readRecipe
// @tags     recipes
// @router   /recipes/{recipeId}  [get]
// @produce  text/markdown
// @param    recipeId path  string  true  "Recipe ID"
// @success  200  {string}  string
// @failure  400  {object}  ApiError
// @failure  404  {object}  ApiError
func (c *RecipeController) Read(ctx echo.Context) error {
	var req ReadRecipeRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	content, err := c.recipes.Read(ctx.Request().Context(), req.ID)
	if err != nil {
		return err
	}

	return ctx.Blob(http.StatusOK, MIMEMarkdown, content)
}

// @summary  Read metadata of a recipe
// @id       readRecipeMetadata
// @tags     recipes
// @router   /recipes/{recipeId}/metadata  [get]
// @produce  application/json
// @param    recipeId path  string  true  "Recipe ID"
// @success  200  {object}  RecipeMetadataDto
// @failure  400  {object}  ApiError
// @failure  404  {object}  ApiError
func (c *RecipeController) ReadMetadata(ctx echo.Context) error {
	var req ReadRecipeRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	metadata, err := c.recipes.ReadMetadata(ctx.Request().Context(), req.ID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, mapRecipeMetadataDto(*metadata))
}

type UpdateRecipeRequest struct {
	ID xid.ID `json:"-" param:"recipeId"`
}

// @summary  Update a recipe
// @id       updateRecipe
// @tags     recipes
// @router   /recipes/{recipeId}  [put]
// @accept   text/markdown
// @param    recipeId path  string  true  "Recipe ID"
// @param    content  body  string  true  "Recipe content"
// @success  204
// @failure  400  {object}  ApiError
// @failure  404  {object}  ApiError
// @failure  409  {object}  ApiError
// @failure  422  {object}  ApiError
func (c *RecipeController) Update(ctx echo.Context) error {
	var req UpdateRecipeRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	content, err := c.readMarkdownRequest(ctx)
	if err != nil {
		return err
	}

	if err := c.recipes.Update(ctx.Request().Context(), req.ID, content); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

type DeleteRecipeRequest struct {
	ID xid.ID `param:"recipeId"`
}

// @summary  Delete a recipe
// @id       deleteRecipe
// @tags     recipes
// @router   /recipes/{recipeId}  [delete]
// @param    recipeId path  string  true  "Recipe ID"
// @success  204
// @failure  400  {object}  ApiError
// @failure  404  {object}  ApiError
func (c *RecipeController) Delete(ctx echo.Context) error {
	var req DeleteRecipeRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := c.recipes.Delete(ctx.Request().Context(), req.ID); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *RecipeController) readMarkdownRequest(ctx echo.Context) ([]byte, error) {
	r := ctx.Request().Body
	defer r.Close()

	return io.ReadAll(r)
}
