package web

import (
	"path"

	"github.com/rs/xid"

	"github.com/lukasdietrich/plaincooking/internal/database/models"
	"github.com/lukasdietrich/plaincooking/internal/oidc"
)

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

func mapUserInfo(claims *oidc.Claims) UserInfoResponse {
	return UserInfoResponse{
		Subject:    claims.Subject,
		Email:      claims.Email,
		Username:   claims.PreferredUsername,
		Name:       claims.Name,
		PictureUrl: claims.Picture,
	}
}

func mapRecipeMetadataDto(entity models.ListRecipeMetadataRow) RecipeMetadataDto {
	return RecipeMetadataDto{
		ID:        entity.RecipeID,
		Title:     entity.Title,
		ImageHref: resolveAssetHref(entity.AssetID),
	}
}

func mapAssetMetadataDto(entity models.Asset) AssetMetadataDto {
	return AssetMetadataDto{
		ID:   entity.ID,
		Href: resolveAssetHref(entity.ID),
	}
}

func resolveAssetHref(id xid.ID) string {
	if id.IsNil() || id.IsZero() {
		return ""
	}

	return path.Join("/api/assets", id.String())
}
