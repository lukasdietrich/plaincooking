package database

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/spf13/viper"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	viper.SetDefault("database.filename", "plaincooking.sqlite")
	viper.SetDefault("database.journalmode", "delete")
}

func Open() (*sql.DB, error) {
	dsn := buildDataSourceName()

	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	if err := migrateUp(conn); err != nil {
		return nil, fmt.Errorf("could not migrate database: %w", err)
	}

	return conn, nil
}

func buildDataSourceName() string {
	opts := make(url.Values)
	opts.Add("_foreign_keys", "true")
	opts.Add("_journal_mode", viper.GetString("database.journalmode"))

	dsn := url.URL{
		Scheme:   "file",
		Opaque:   viper.GetString("database.filename"),
		RawQuery: opts.Encode(),
	}

	return dsn.String()
}
