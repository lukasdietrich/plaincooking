package main

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	setupConfig()

	if err := run(); err != nil {
		slog.Error("Fatal", slog.Any("error", err))
	}
}

func setupConfig() {
	viper.SetTypeByDefaultValue(true)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("PLAINCOOKING")
}

func run() error {
	server, err := InjectServer()
	if err != nil {
		return fmt.Errorf("could not inject server: %w", err)
	}

	slog.Info("starting http server", slog.String("addr", server.Addr))
	return server.ListenAndServe()
}
