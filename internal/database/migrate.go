package database

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/spf13/viper"
)

//go:embed migrations/*.sql
var migrationsFs embed.FS

func init() {
	viper.SetDefault("database.migration.tablename", "migrations")
}

func migrateUp(db *sql.DB) error {
	config := sqlite3.Config{
		MigrationsTable: viper.GetString("database.migration.tablename"),
	}

	source, err := iofs.New(migrationsFs, "migrations")
	if err != nil {
		return fmt.Errorf("could not load migrations: %w", err)
	}

	database, err := sqlite3.WithInstance(db, &config)
	if err != nil {
		return fmt.Errorf("could not create database driver: %w", err)
	}

	migrator, err := migrate.NewWithInstance("iofs", source, "sqlite3", database)
	if err != nil {
		return fmt.Errorf("could not create migrator: %w", err)
	}

	migrator.Log = migratorLogger{slog.Default()}

	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("could not apply migrations: %w", err)
	}

	return nil
}

type migratorLogger struct {
	logger *slog.Logger
}

func (l migratorLogger) Printf(format string, v ...interface{}) {
	l.logger.Debug(format, v...)
}

func (l migratorLogger) Verbose() bool {
	return true
}
