package oidc

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

const (
	queryOrigin = "origin"
	queryCode   = "code"
	queryState  = "state"

	extraIdToken = "id_token"
)

var (
	ErrMissingIdToken = errors.New("missing id_token")
)

type Handler struct {
	provider *oidc.Provider
	config   *oauth2.Config
	session  *Session
}

func NewHandler(provider *oidc.Provider, config *oauth2.Config, session *Session) *Handler {
	return &Handler{
		provider: provider,
		config:   config,
		session:  session,
	}
}

func (h *Handler) Redirect(ctx echo.Context) error {
	origin := ctx.QueryParam(queryOrigin)
	url := h.config.AuthCodeURL(origin)

	slog.Debug("redirecting to oidc provider", slog.String("url", url))
	return ctx.Redirect(http.StatusFound, url)
}

func (h *Handler) Callback(ctx echo.Context) error {
	token, err := h.exchange(ctx)
	if err != nil {
		return fmt.Errorf("could not exchange access token: %w", err)
	}

	idToken, err := h.verify(ctx, token)
	if err != nil {
		return fmt.Errorf("could not verify idToken: %w", err)
	}

	slog.Debug("exchanged idToken",
		slog.String("subject", idToken.Subject),
		slog.String("issuer", idToken.Issuer))

	userInfo, err := h.userInfo(ctx, token)
	if err != nil {
		return fmt.Errorf("could not fetch userInfo: %w", err)
	}

	slog.Debug("received userInfo",
		slog.String("subject", userInfo.Subject),
		slog.String("email", userInfo.Email))

	if err := h.session.saveClaims(ctx, userInfo); err != nil {
		return fmt.Errorf("could not save claims: %w", err)
	}

	return h.redirectToOrigin(ctx)
}

func (h *Handler) redirectToOrigin(ctx echo.Context) error {
	target := ctx.QueryParam(queryState)
	if target == "" {
		target = "/"
	}

	return ctx.Redirect(http.StatusFound, target)
}

func (h *Handler) exchange(ctx echo.Context) (*oauth2.Token, error) {
	rctx := ctx.Request().Context()
	code := ctx.QueryParam(queryCode)

	return h.config.Exchange(rctx, code)
}

func (h *Handler) verify(ctx echo.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIdToken, ok := token.Extra(extraIdToken).(string)
	if !ok {
		return nil, ErrMissingIdToken
	}

	rctx := ctx.Request().Context()

	verifier := h.provider.VerifierContext(
		rctx,
		&oidc.Config{
			ClientID: h.config.ClientID,
		},
	)

	return verifier.Verify(rctx, rawIdToken)
}

func (h *Handler) userInfo(ctx echo.Context, token *oauth2.Token) (*oidc.UserInfo, error) {
	rctx := ctx.Request().Context()
	tokenSource := oauth2.StaticTokenSource(token)

	return h.provider.UserInfo(rctx, tokenSource)
}
