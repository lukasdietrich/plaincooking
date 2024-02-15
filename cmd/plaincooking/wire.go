//go:build wireinject

package main

import (
	"net/http"

	"github.com/google/wire"

	"github.com/lukasdietrich/plaincooking/internal/web"
)

func InjectServer() (*http.Server, error) {
	wire.Build(
		web.NewServer,
		web.NewRouter,
	)

	return nil, nil
}
