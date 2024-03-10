package web

import "github.com/lukasdietrich/plaincooking/internal/database/models"

func mapSlice[Src, Dst any](sourceSlice []Src, fn func(Src) Dst) []Dst {
	if sourceSlice == nil {
		return nil
	}

	destinationSlice := make([]Dst, len(sourceSlice))
	for i, s := range sourceSlice {
		destinationSlice[i] = fn(s)
	}

	return destinationSlice
}

func mapRecipeMetadataDto(entity models.RecipeMetadata) RecipeMetadataDto {
	return RecipeMetadataDto{
		ID:    entity.RecipeID,
		Title: entity.Title,
	}
}
