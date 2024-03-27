package web

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/xid"

	"github.com/lukasdietrich/plaincooking/internal/service"
)

const (
	MediaTypeMarkdown = "text/markdown"
)

type ApiError struct {
	Status   int    `json:"status"`
	Code     string `json:"code"`
	Internal error  `json:"-"`
} // @name PlaincookingApiError

func (e ApiError) Error() string {
	return fmt.Sprintf("api error status=%d, code=%q: %v", e.Status, e.Code, e.Internal)
}

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
	ID        xid.ID `json:"id"`
	Title     string `json:"title"`
	ImageHref string `json:"imageHref"`
} // @name RecipeMetadata

type AssetMetadataDto struct {
	ID   xid.ID `json:"id"`
	Href string `json:"href"`
} // @name AssetMetadata

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
// @param    recipeId  path  string  true  "Recipe ID"
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

	return ctx.Blob(http.StatusOK, MediaTypeMarkdown, content)
}

type UpdateRecipeRequest struct {
	ID xid.ID `json:"-" param:"recipeId"`
}

// @summary  Update a recipe
// @id       updateRecipe
// @tags     recipes
// @router   /recipes/{recipeId}  [put]
// @accept   text/markdown
// @param    recipeId  path  string  true  "Recipe ID"
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
// @param    recipeId  path  string  true  "Recipe ID"
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

// @summary  List recipe images
// @id       listRecipeImages
// @tags     recipes assets
// @router   /recipes/{recipeId}/images  [get]
// @param    recipeId  path  string  true  "Recipe ID"
// @produce  application/json
// @success  200  {array}  AssetMetadataDto
func (c *RecipeController) ListImages(ctx echo.Context) error {
	var req ReadRecipeRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	assetSlice, err := c.recipes.ListImages(ctx.Request().Context(), req.ID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, mapSlice(assetSlice, mapAssetMetadataDto))
}

type UploadRecipeImageRequest struct {
	ID xid.ID `json:"-" param:"recipeId"`
}

// @summary  Upload a new recipe image
// @id       uploadRecipeImage
// @tags     recipes assets
// @router   /recipes/{recipeId}/images  [post]
// @accept   multipart/form-data
// @produce  application/json
// @param    recipeId  path  string  true  "Recipe ID"
// @param    image  formData  file  true  "Image"
// @success  201  {object}  AssetMetadataDto
// @failure  400  {object}  ApiError
// @failure  404  {object}  ApiError
// @failure  422  {object}  ApiError
func (c *RecipeController) UploadImage(ctx echo.Context) error {
	var req UploadRecipeImageRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	part, err := c.findFormPart(ctx, "image")
	if err != nil {
		return err
	}

	defer part.Close()

	filename := part.FileName()
	mediaTyp := part.Header.Get(echo.HeaderContentType)

	w, err := c.recipes.ImageWriter(ctx.Request().Context(), req.ID, filename, mediaTyp)
	if err != nil {
		return err
	}

	if _, err := io.Copy(w, part); err != nil {
		w.Close() // nolint:errcheck
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, mapAssetMetadataDto(w.Asset))
}

func (c *RecipeController) readMarkdownRequest(ctx echo.Context) ([]byte, error) {
	r := ctx.Request().Body
	defer r.Close()

	return io.ReadAll(r)
}

func (c *RecipeController) findFormPart(ctx echo.Context, name string) (*multipart.Part, error) {
	r, err := ctx.Request().MultipartReader()
	if err != nil {
		return nil, err
	}

	part, err := r.NextPart()
	if err != nil {
		return nil, err
	}

	if part.FormName() != name {
		return nil, echo.NewHTTPError(
			http.StatusUnprocessableEntity,
			fmt.Sprintf("multipart part is not named %q", name),
		)
	}

	return part, nil
}

type AssetController struct {
	assets *service.AssetService
}

func NewAssetController(assets *service.AssetService) *AssetController {
	return &AssetController{
		assets: assets,
	}
}

type DownloadAssetRequest struct {
	ID        xid.ID                `param:"assetId"`
	Thumbnail service.ThumbnailSize `query:"thumbnail"`
}

// @summary  Download an asset
// @id       downloadAsset
// @tags     assets
// @router   /assets/{assetId}  [get]
// @produce  application/octet-stream
// @param    assetId  path  string  true  "Asset ID"
// @param    thumbnail  query  string  false  "Thumbnail version of an image"
// @success  200  {blob}  blob
// @failure  400  {object}  ApiError
// @failure  404  {object}  ApiError
func (c *AssetController) Download(ctx echo.Context) error {
	var req DownloadAssetRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	r, err := c.reader(ctx, req)
	if err != nil {
		return err
	}

	header := ctx.Response().Header()
	header.Add(echo.HeaderContentLength, fmt.Sprintf("%d", r.TotalSize))

	return ctx.Stream(http.StatusOK, r.MediaType, r)
}

func (c *AssetController) reader(ctx echo.Context, req DownloadAssetRequest) (*service.AssetReader, error) {
	if req.Thumbnail != service.ThumbnailUnknown {
		return c.assets.ThumbnailReader(ctx.Request().Context(), req.ID, req.Thumbnail)
	}

	return c.assets.Reader(ctx.Request().Context(), req.ID)
}
