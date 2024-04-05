package web

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mattn/go-sqlite3"
	"github.com/rs/xid"

	"github.com/lukasdietrich/plaincooking/internal/oidc"
	"github.com/lukasdietrich/plaincooking/internal/parser"
)

var (
	ErrInternal = ApiError{
		Status: http.StatusInternalServerError,
		Code:   "internal",
	}

	ErrUnauthorized = ApiError{
		Status: http.StatusUnauthorized,
		Code:   "unauthorized",
	}

	ErrResourceNotFound = ApiError{
		Status: http.StatusNotFound,
		Code:   "resource.notfound",
	}

	ErrContraintViolation = ApiError{
		Status: http.StatusConflict,
		Code:   "validation.constraint.violation",
	}

	ErrRecipeInvalid = ApiError{
		Status: http.StatusUnprocessableEntity,
		Code:   "validation.recipe.invalid",
	}

	ErrMultipartUnexpectedPart = ApiError{
		Status: http.StatusUnprocessableEntity,
		Code:   "validation.multipart.unexpectedpart",
	}
)

type ApiError struct {
	Status   int    `json:"status"`
	Code     string `json:"code"`
	Internal error  `json:"-"`
} // @name ApiError

func (e ApiError) Error() string {
	return fmt.Sprintf("api error status=%d, code=%q: %v", e.Status, e.Code, e.Internal)
}

func (e ApiError) WithInternal(err error) ApiError {
	e.Internal = err
	return e
}

func marshalApiError() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			err := next(ctx)

			if err, ok := err.(ApiError); ok {
				return ctx.JSON(err.Status, err)
			}

			return err
		}
	}
}

func logError() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			err := next(ctx)
			if err != nil {
				slog.Warn("request error",
					slog.String("url", ctx.Request().RequestURI),
					slog.String("err", err.Error()))
			}

			return err
		}
	}
}

func handleBusinessError() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if err := next(ctx); err != nil {
				mapped := mapBusinessError(err)

				if apiError, ok := mapped.(ApiError); ok {
					return apiError.WithInternal(err)
				}

				return mapped
			}

			return nil
		}
	}
}

func mapBusinessError(err error) error {
	if errors.Is(err, sql.ErrNoRows) || errors.Is(err, xid.ErrInvalidID) {
		return ErrResourceNotFound
	}

	if errors.Is(err, parser.ErrInvalidRecipe) {
		return ErrRecipeInvalid
	}

	if sqliteErr, ok := errorAs[sqlite3.Error](err); ok {
		return mapSqliteError(sqliteErr)
	}

	if _, ok := errorAs[oidc.AuthorizationError](err); ok {
		return ErrUnauthorized
	}

	return ErrInternal
}

func mapSqliteError(err sqlite3.Error) error {
	if err.Code == sqlite3.ErrConstraint {
		return ErrContraintViolation
	}

	return err
}

func errorAs[E error](err error) (E, bool) {
	var errTyped E
	ok := errors.As(err, &errTyped)

	return errTyped, ok
}
