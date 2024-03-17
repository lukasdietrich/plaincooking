package service

import (
	"context"
	"time"

	"github.com/rs/xid"

	"github.com/lukasdietrich/plaincooking/internal/database"
	"github.com/lukasdietrich/plaincooking/internal/database/models"
	"github.com/lukasdietrich/plaincooking/internal/parser"
)

type RecipeService struct {
	querier database.Querier
	parser  *parser.RecipeParser
}

func NewRecipeService(querier database.Querier, parser *parser.RecipeParser) *RecipeService {
	return &RecipeService{
		querier: querier,
		parser:  parser,
	}
}

func (s *RecipeService) List(ctx context.Context) ([]models.RecipeMetadata, error) {
	return s.querier.ListRecipeMetadata(ctx)
}

func (s *RecipeService) Read(ctx context.Context, id xid.ID) ([]byte, error) {
	recipe, err := s.querier.ReadRecipe(ctx, models.ReadRecipeParams{ID: id})
	if err != nil {
		return nil, err
	}

	return recipe.Content, nil
}

func (s *RecipeService) Create(ctx context.Context, content []byte) (xid.ID, error) {
	querier, tx, err := s.querier.Begin(ctx)
	if err != nil {
		return xid.NilID(), err
	}

	defer tx.Rollback()

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

	if err := tx.Commit(); err != nil {
		return xid.NilID(), err
	}

	return createRecipeParams.ID, nil
}

func (s *RecipeService) Update(ctx context.Context, id xid.ID, content []byte) error {
	querier, tx, err := s.querier.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback()

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

	return tx.Commit()
}

func (s *RecipeService) Delete(ctx context.Context, id xid.ID) error {
	_, err := s.querier.DeleteRecipe(ctx, models.DeleteRecipeParams{ID: id})
	return err
}