//go:build wireinject

package main

import (
	"database/sql"
	"net/http"

	"github.com/google/wire"

	"github.com/lukasdietrich/plaincooking/internal/database"
	"github.com/lukasdietrich/plaincooking/internal/database/models"
	"github.com/lukasdietrich/plaincooking/internal/parser"
	"github.com/lukasdietrich/plaincooking/internal/service"
	"github.com/lukasdietrich/plaincooking/internal/web"
)

func InjectServer() (*http.Server, error) {
	wire.Build(
		// Web
		web.NewServer,
		web.NewRouter,
		web.NewAssetController,
		web.NewRecipeController,

		// Service
		service.NewAssetService,
		service.NewRecipeService,

		// Parser
		parser.NewParser,

		// Database
		database.Open,
		database.NewQuerier,

		wire.Bind(new(models.DBTX), new(*sql.DB)),
	)

	return nil, nil
}
