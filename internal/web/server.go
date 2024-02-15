package web

import (
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("http.addr", ":8080")
	viper.SetDefault("http.timeout.read", 60*time.Second)
	viper.SetDefault("http.timeout.header", 60*time.Second)
	viper.SetDefault("http.timeout.write", 60*time.Second)
	viper.SetDefault("http.timeout.idle", 10*time.Second)
}

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              viper.GetString("http.addr"),
		Handler:           handler,
		ReadTimeout:       viper.GetDuration("http.timeout.read"),
		ReadHeaderTimeout: viper.GetDuration("http.timeout.header"),
		WriteTimeout:      viper.GetDuration("http.timeout.write"),
		IdleTimeout:       viper.GetDuration("http.timeout.idle"),
	}
}
