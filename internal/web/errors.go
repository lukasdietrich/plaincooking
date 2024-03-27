package web

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mattn/go-sqlite3"
	"github.com/rs/xid"

	"github.com/lukasdietrich/plaincooking/internal/parser"
)

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
				return mapBusinessError(err)
			}

			return nil
		}
	}
}

func mapBusinessError(err error) error {
	if errors.Is(err, xid.ErrInvalidID) {
		return ApiError{
			Status:   http.StatusBadRequest,
			Code:     "invalid id",
			Internal: err,
		}
	}

	if errors.Is(err, sql.ErrNoRows) {
		return ApiError{
			Status:   http.StatusNotFound,
			Code:     "resource not found",
			Internal: err,
		}
	}

	if errors.Is(err, parser.ErrInvalidRecipe) {
		return ApiError{
			Status:   http.StatusUnprocessableEntity,
			Code:     "invalid recipe",
			Internal: err,
		}
	}

	if sqliteErr, ok := errorAs[sqlite3.Error](err); ok {
		return mapSqliteError(sqliteErr)
	}

	return ApiError{
		Status:   http.StatusInternalServerError,
		Code:     "internal",
		Internal: err,
	}
}

func mapSqliteError(err sqlite3.Error) error {
	if err.Code == sqlite3.ErrConstraint {
		return ApiError{
			Status:   http.StatusConflict,
			Code:     "constraint violation",
			Internal: err,
		}
	}

	return err
}

func errorAs[E error](err error) (E, bool) {
	var errTyped E
	ok := errors.As(err, &errTyped)

	return errTyped, ok
}
