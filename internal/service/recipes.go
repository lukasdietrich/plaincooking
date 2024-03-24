package service

import (
	"context"
	"time"

	"github.com/rs/xid"

	"github.com/lukasdietrich/plaincooking/internal/database/models"
	"github.com/lukasdietrich/plaincooking/internal/parser"
)

type RecipeService struct {
	transactions *TransactionService
	assets       *AssetService
	parser       *parser.RecipeParser
}

func NewRecipeService(
	transactions *TransactionService,
	assets *AssetService,
	parser *parser.RecipeParser,
) *RecipeService {
	return &RecipeService{
		transactions: transactions,
		assets:       assets,
		parser:       parser,
	}
}

func (s *RecipeService) List(ctx context.Context) ([]models.ListRecipeMetadataRow, error) {
	querier := s.transactions.Querier(ctx)
	return querier.ListRecipeMetadata(ctx)
}

func (s *RecipeService) Read(ctx context.Context, id xid.ID) ([]byte, error) {
	querier := s.transactions.Querier(ctx)

	recipe, err := querier.ReadRecipe(ctx, models.ReadRecipeParams{ID: id})
	if err != nil {
		return nil, err
	}

	return recipe.Content, nil
}

func (s *RecipeService) Create(ctx context.Context, content []byte) (xid.ID, error) {
	querier := s.transactions.Querier(ctx)

	meta, err := s.parser.ParseRecipe(content)
	if err != nil {
		return xid.NilID(), err
	}

	createRecipeParams := models.CreateRecipeParams{
		ID:        xid.New(),
		CreatedAt: time.Now(),
		Content:   content,
	}

	createRecipeMetadataParams := models.CreateRecipeMetadataParams{
		RecipeID: createRecipeParams.ID,
		Title:    meta.Title,
	}

	if err := querier.CreateRecipe(ctx, createRecipeParams); err != nil {
		return xid.NilID(), err
	}

	if err := querier.CreateRecipeMetadata(ctx, createRecipeMetadataParams); err != nil {
		return xid.NilID(), err
	}

	return createRecipeParams.ID, nil
}

func (s *RecipeService) Update(ctx context.Context, id xid.ID, content []byte) error {
	querier := s.transactions.Querier(ctx)

	meta, err := s.parser.ParseRecipe(content)
	if err != nil {
		return err
	}

	updateRecipeParams := models.UpdateRecipeParams{
		ID:        id,
		UpdatedAt: time.Now(),
		Content:   content,
	}

	updateRecipeMetadataParams := models.UpdateRecipeMetadataParams{
		RecipeID: updateRecipeParams.ID,
		Title:    meta.Title,
	}

	if _, err := querier.UpdateRecipe(ctx, updateRecipeParams); err != nil {
		return err
	}

	if _, err := querier.UpdateRecipeMetadata(ctx, updateRecipeMetadataParams); err != nil {
		return err
	}

	return nil
}

func (s *RecipeService) Delete(ctx context.Context, id xid.ID) error {
	querier := s.transactions.Querier(ctx)
	_, err := querier.DeleteRecipe(ctx, models.DeleteRecipeParams{ID: id})
	return err
}

func (s *RecipeService) ListImages(ctx context.Context, id xid.ID) ([]models.Asset, error) {
	querier := s.transactions.Querier(ctx)
	return querier.ListRecipeAssets(ctx, models.ListRecipeAssetsParams{RecipeID: id})
}

func (s *RecipeService) ImageWriter(ctx context.Context, id xid.ID, filename, mediaType string) (*assetWriter, error) {
	querier := s.transactions.Querier(ctx)

	w, err := s.assets.Writer(ctx, filename, mediaType)
	if err != nil {
		return nil, err
	}

	createRecipeAssetParams := models.CreateRecipeAssetParams{
		RecipeID: id,
		AssetID:  w.ID,
	}

	if err := querier.CreateRecipeAsset(ctx, createRecipeAssetParams); err != nil {
		return nil, err
	}

	return w, nil
}
