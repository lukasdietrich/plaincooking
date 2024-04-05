package main

import (
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("logging.level", "info")
}

func main() {
	setupConfig()
	setupLogger()
	printConfig()

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

func setupLogger() {
	levelString := viper.GetString("logging.level")

	level := slog.LevelDebug
	err := level.UnmarshalText([]byte(levelString))

	options := slog.HandlerOptions{
		AddSource: false,
		Level:     level,
	}

	handler := slog.NewTextHandler(os.Stderr, &options)
	logger := slog.New(handler)

	slog.SetDefault(logger)

	if err != nil {
		slog.Error("invalid logging level", slog.String("level", levelString))
	}

	slog.Debug("set logging level", slog.Any("level", level))
}

func printConfig() {
	keys := viper.AllKeys()
	sort.Strings(keys)

	for _, key := range keys {
		value := viper.Get(key)
		slog.Debug("config", slog.Any(key, value))
	}
}

func run() error {
	server, err := InjectServer()
	if err != nil {
		return fmt.Errorf("could not inject server: %w", err)
	}

	slog.Info("starting http server", slog.String("addr", server.Addr))
	return server.ListenAndServe()
}
