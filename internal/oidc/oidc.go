package oidc

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

const (
	scopeProfile = "profile"
	scopeEmail   = "email"
)

func init() {
	viper.SetDefault("oidc.issuer", nil)
	viper.SetDefault("oidc.client.id", nil)
	viper.SetDefault("oidc.client.secret", nil)
	viper.SetDefault("oidc.redirect.url", "http://localhost:8080/oauth2/callback")
}

func NewProvider() (*oidc.Provider, error) {
	issuer := viper.GetString("oidc.issuer")

	return oidc.NewProvider(context.Background(), issuer)
}

func NewConfig(provider *oidc.Provider) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     viper.GetString("oidc.client.id"),
		ClientSecret: viper.GetString("oidc.client.secret"),
		RedirectURL:  viper.GetString("oidc.redirect.url"),

		Endpoint: provider.Endpoint(),

		Scopes: []string{
			oidc.ScopeOpenID,
			scopeProfile,
			scopeEmail,
		},
	}
}
