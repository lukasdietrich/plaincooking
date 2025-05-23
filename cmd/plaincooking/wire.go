//go:build wireinject

package main

import (
	"net/http"

	"github.com/google/wire"

	"github.com/lukasdietrich/plaincooking/internal/database"
	"github.com/lukasdietrich/plaincooking/internal/oidc"
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
		web.NewSessionController,
		web.NewProbeController,

		// OpenID Connect
		oidc.NewProvider,
		oidc.NewConfig,
		oidc.NewHandler,
		oidc.NewSession,

		// Service
		service.NewTransactionService,
		service.NewAssetService,
		service.NewRecipeService,

		// Parser
		parser.NewParser,

		// Database
		database.Open,
	)

	return nil, nil
}
