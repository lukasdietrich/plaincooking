//go:build wireinject

package main

import (
	"database/sql"
	"net/http"

	"github.com/google/wire"

	"github.com/lukasdietrich/plaincooking/internal/database"
	"github.com/lukasdietrich/plaincooking/internal/database/models"
	"github.com/lukasdietrich/plaincooking/internal/web"
)

func InjectServer() (*http.Server, error) {
	wire.Build(
		web.NewServer,
		web.NewRouter,

		database.Open,
		models.New,

		wire.Bind(new(models.DBTX), new(*sql.DB)),
		wire.Bind(new(models.Querier), new(*models.Queries)),
	)

	return nil, nil
}
