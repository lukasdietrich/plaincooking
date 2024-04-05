package oidc

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
)

const (
	cookieName = "__PLAINCOOKING_TOKEN"
	contextKey = "__PLAINCOOKING_CLAIMS"
)

type AuthorizationError struct {
	message string
	cause   error
}

func (e AuthorizationError) Error() string {
	return fmt.Sprintf("%s: %v", e.message, e.cause)
}

type Claims struct {
	jwt.RegisteredClaims

	Email             string `json:"email"`
	PreferredUsername string `json:"preferred_username"`
	Name              string `json:"name"`
	Picture           string `json:"picture"`
}

type Session struct {
	secret []byte
}

func NewSession() (*Session, error) {
	secret, err := randomBytes(64)
	if err != nil {
		return nil, err
	}

	return &Session{
		secret: secret,
	}, nil
}

func (s *Session) Require(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		claims, err := s.Claims(ctx)
		if err != nil {
			return err
		}

		ctx.Set(contextKey, claims)
		return next(ctx)
	}
}

func (s *Session) Claims(ctx echo.Context) (*Claims, error) {
	if claims, ok := ctx.Get(contextKey).(*Claims); ok {
		return claims, nil
	}

	cookie, err := ctx.Cookie(cookieName)
	if err != nil {
		return nil, AuthorizationError{"could not read session cookie", err}
	}

	claims, err := s.decodeClaims(cookie.Value)
	if err != nil {
		s.setCookie(ctx, "", -1)
		return nil, AuthorizationError{"could not decode session claims", err}
	}

	return claims, nil
}

func (s *Session) saveClaims(ctx echo.Context, userInfo *oidc.UserInfo) error {
	claims, err := s.claimsFromUserInfo(userInfo)
	if err != nil {
		return err
	}

	token, err := s.encodeAndSignClaims(claims)
	if err != nil {
		return err
	}

	s.setCookie(ctx, token, 0)
	return nil
}

func (s *Session) setCookie(ctx echo.Context, token string, maxAge int) {
	cookie := http.Cookie{
		Name:  cookieName,
		Value: token,

		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
	}

	ctx.SetCookie(&cookie)
}

func (s *Session) claimsFromUserInfo(userInfo *oidc.UserInfo) (*Claims, error) {
	var claims Claims

	if err := userInfo.Claims(&claims); err != nil {
		return nil, err
	}

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(time.Hour * 2)

	claims.RegisteredClaims = jwt.RegisteredClaims{
		ID:        xid.New().String(),
		Subject:   userInfo.Subject,
		IssuedAt:  jwt.NewNumericDate(issuedAt),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	return &claims, nil
}

func (s *Session) encodeAndSignClaims(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	return token.SignedString(s.secret)
}

func (s *Session) decodeClaims(token string) (*Claims, error) {
	var claims Claims

	_, err := jwt.ParseWithClaims(token, &claims, s.key)
	return &claims, err
}

func (s *Session) key(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
	}

	return s.secret, nil
}

func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, b)
	return b, err
}
